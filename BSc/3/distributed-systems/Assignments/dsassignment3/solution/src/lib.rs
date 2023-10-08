use std::collections::{HashMap, HashSet};
use std::time::Duration;
use std::time::SystemTime;

use async_channel::Sender;

use executor::{Handler, ModuleRef, System};
use rand::Rng;
use uuid::Uuid;

mod state_machine_handler;
use state_machine_handler::*;

mod timer;
use timer::*;
mod state;
use state::*;

pub use domain::*;

mod domain;

/// Process of Raft.
pub struct Raft {
    state: State,
    config: ServerConfig,
    message_sender: Box<dyn RaftSender>,
    state_machine_mod_ref: ModuleRef<StateMachineHandler>,
    process_type: ProcessType,
    enabled: bool,
    election_timer: Timer,
    minimal_election_timer: Timer,
    heartbeat_timer: Timer,
    catchup_timer: Timer,
    self_ref: Option<ModuleRef<Self>>,
}

impl Raft {
    pub async fn new(
        system: &mut System,
        config: ServerConfig,
        first_log_entry_timestamp: SystemTime,
        state_machine: Box<dyn StateMachine>,
        stable_storage: Box<dyn StableStorage>,
        message_sender: Box<dyn RaftSender>,
    ) -> ModuleRef<Self> {
        let state = State::new(stable_storage, &config, first_log_entry_timestamp).await;

        let election_timeout = Self::sample_election_timeout(&config);

        let serialized = state.get_snapshot_data();

        let state_machine_mod_ref =
            StateMachineHandler::new(system, state_machine, config.session_expiration).await;

        let self_ref = system
            .register_module(Self {
                state,
                heartbeat_timer: Timer::new(config.heartbeat_timeout),
                minimal_election_timer: Timer::new(*config.election_timeout_range.start()),
                catchup_timer: Timer::new(2 * (*config.election_timeout_range.end())),
                config,
                election_timer: Timer::new(election_timeout),
                state_machine_mod_ref: state_machine_mod_ref.clone(),
                message_sender,
                process_type: ProcessType::Follower {
                    is_leader_reliable: false,
                },
                enabled: true,
                self_ref: None,
            })
            .await;

        state_machine_mod_ref
            .send(InitStateMachine {
                raft_ref: self_ref.clone(),
                serialized,
            })
            .await;

        self_ref
            .send(Init {
                raft_ref: self_ref.clone(),
            })
            .await;
        self_ref
    }

    fn sample_election_timeout(config: &ServerConfig) -> Duration {
        rand::thread_rng().gen_range(config.election_timeout_range.clone())
    }

    fn reset_election_timer(&mut self) {
        self.election_timer
            .reset(self.get_self_ref(), ElectionTimeout);
    }

    fn reset_catchup_timer(&mut self) {
        self.catchup_timer
            .reset(self.get_self_ref(), CatchUpTimeout);
    }

    fn stop_catchup_timer(&mut self) {
        self.catchup_timer.stop();
    }

    fn get_self_ref(&self) -> ModuleRef<Self> {
        self.self_ref.clone().unwrap()
    }

    /// Set the process's term to the higher number.
    fn update_term(&mut self, new_term: u64) {
        assert!(self.state.persistent.current_term < new_term);
        self.state.persistent.current_term = new_term;
        self.state.persistent.voted_for = None;
        self.state.persistent.leader_id = None;
    }

    /// Broadcast heartbeat.
    async fn broadcast_heartbeat(&mut self) {
        if let ProcessType::Leader {
            next_index,
            match_index,
            ..
        } = &self.process_type
        {
            for server in self.state.get_servers().iter() {
                self.send_heartbeat(server, next_index[server], match_index[server])
                    .await
            }
        } else {
            panic!("Not leader should never broadcast append entries!")
        }
    }

    async fn send_heartbeat(&self, server: &Uuid, next_index: usize, match_index: usize) {
        let prev_log_index = next_index - 1;

        let entries_option = if prev_log_index == match_index {
            let end_idx = std::cmp::min(
                next_index + self.config.append_entries_batch_size,
                self.state.log_last_index() + 1,
            );

            self.state.log_slice(next_index, end_idx)
        } else {
            Some(Vec::new())
        };

        match entries_option {
            Some(entries) => {
                self.send_append_entries(server, entries, prev_log_index)
                    .await
            }
            None => {
                self.send_install_snapshot(server).await;
            }
        };
    }

    async fn send_append_entries(
        &self,
        server: &Uuid,
        entries: Vec<LogEntry>,
        prev_log_index: usize,
    ) {
        let msg = RaftMessage {
            header: self.get_header(),
            content: RaftMessageContent::AppendEntries(AppendEntriesArgs {
                prev_log_index,
                prev_log_term: self.state.log_get_term(prev_log_index),
                entries,
                leader_commit: self.state.volatile.commit_index,
            }),
        };

        self.send(server, msg).await;
    }

    async fn send_install_snapshot(&self, server: &Uuid) {
        let offset = self.state.get_snapshot_offset(server);

        let msg = RaftMessage {
            header: self.get_header(),
            content: RaftMessageContent::InstallSnapshot(
                self.state.get_install_snapshot_args(offset),
            ),
        };

        self.send(server, msg).await;
    }

    async fn become_candidate(&mut self) {
        if self.state.get_servers().len() == 1 {
            self.become_leader().await;
            return;
        }

        let mut votes_received = HashSet::new();
        votes_received.insert(self.config.self_id);

        self.process_type = ProcessType::Candidate { votes_received };
        self.stop_minimal_election_timer();

        self.update_term(self.state.persistent.current_term + 1);
        self.state.persistent.voted_for = Some(self.config.self_id);
        self.state.save().await;

        self.broadcast(RaftMessage {
            header: self.get_header(),
            content: RaftMessageContent::RequestVote(RequestVoteArgs {
                last_log_index: self.state.log_last_index(),
                last_log_term: self.state.log_last_term(),
            }),
        })
        .await;
    }

    async fn become_leader(&mut self) {
        let log_len = self.state.log_last_index() + 1;

        let next_index = self
            .config
            .servers
            .iter()
            .map(|uuid| (*uuid, log_len))
            .collect();

        let match_index = self
            .state
            .get_servers()
            .iter()
            .map(|uuid| (*uuid, 0))
            .collect();

        let cluster_change =
            if self.state.volatile.last_config_change_idx > self.state.volatile.last_applied {
                ClusterChange::CommitOldConfig {
                    idx: self.state.volatile.last_config_change_idx,
                }
            } else {
                ClusterChange::None
            };

        self.process_type = ProcessType::Leader {
            next_index,
            reply_to: HashMap::new(),
            match_index,
            responded: Self::get_fresh_responded(self.config.self_id),
            cluster_change,
        };

        self.state.persistent.leader_id = Some(self.config.self_id);
        self.state.save().await;

        self.reset_election_timer();
        self.start_heartbeat_timer();
        self.stop_minimal_election_timer();
    }

    fn get_fresh_responded(uuid: Uuid) -> HashSet<Uuid> {
        let mut responded = HashSet::new();
        responded.insert(uuid);

        responded
    }

    fn get_header(&self) -> RaftMessageHeader {
        RaftMessageHeader {
            term: self.state.persistent.current_term,
            source: self.config.self_id,
        }
    }

    async fn handle_append_entries(&mut self, header: RaftMessageHeader, args: AppendEntriesArgs) {
        let newer_term = header.term >= self.state.persistent.current_term;
        let mut success = newer_term;

        if newer_term {
            self.state.persistent.leader_id = Some(header.source);

            match &self.process_type {
                ProcessType::Leader { .. } => {
                    log::debug!("Ignore, heartbeat from self")
                }
                _ => {
                    self.process_type = ProcessType::Follower {
                        is_leader_reliable: true,
                    };
                    self.reset_election_timer();
                    self.stop_heartbeat_timer();
                    self.stop_catchup_timer();
                    self.restart_minimal_election_timer();

                    if self.log_matches(&args) {
                        let last_commit_entry_idx = self.update_log(args).await;
                        self.update_commit_index(last_commit_entry_idx).await;
                    } else {
                        success = false;
                    }
                }
            };

            self.state.save().await;
        }

        self.send(
            &header.source,
            RaftMessage {
                header: self.get_header(),
                content: RaftMessageContent::AppendEntriesResponse(AppendEntriesResponseArgs {
                    success,
                    last_log_index: self.state.log_last_index(),
                }),
            },
        )
        .await;
    }

    async fn handle_append_entries_response(
        &mut self,
        header: RaftMessageHeader,
        args: AppendEntriesResponseArgs,
    ) {
        if let ProcessType::Leader {
            next_index,
            match_index,
            responded,
            cluster_change,
            ..
        } = &mut self.process_type
        {
            if let ClusterChange::CatchUp(CatchUpState {
                reply_to: _,
                new_server,
                round: _,
                idx_to_catchup,
                next_idx,
                match_idx,
            }) = cluster_change
            {
                if *new_server == header.source {
                    if args.success {
                        *next_idx = args.last_log_index + 1;
                        *match_idx = args.last_log_index;

                        if *match_idx >= *idx_to_catchup {
                            self.new_catchup_round().await;
                        }
                    } else {
                        *next_idx = std::cmp::min(*next_idx - 1, args.last_log_index + 1);
                    }

                    self.catch_up().await;
                    return;
                }
            }

            if args.success {
                responded.insert(header.source);
                next_index.insert(header.source, args.last_log_index + 1);
                match_index.insert(header.source, args.last_log_index);

                self.try_apply().await;
            } else {
                next_index.insert(
                    header.source,
                    std::cmp::min(next_index[&header.source] - 1, args.last_log_index + 1),
                );
            }
            self.try_send_heartbeat(&header.source).await;
        } else {
            log::debug!("Ignore, got AppendEntriesResponseArgs as not leader");
        }
    }

    async fn try_send_heartbeat(&self, client_id: &Uuid) {
        if let ProcessType::Leader {
            next_index,
            match_index,
            ..
        } = &self.process_type
        {
            self.send_heartbeat(&client_id, next_index[client_id], match_index[client_id])
                .await;
        }
    }

    async fn try_apply(&mut self) {
        if let ProcessType::Leader { match_index, .. } = &self.process_type {
            let values: Vec<usize> = match_index.values().map(|&val| val).collect();

            while values
                .iter()
                .filter(|&id| *id > self.state.volatile.commit_index)
                .count()
                > self.state.get_servers().len() / 2
            {
                self.state.volatile.commit_index += 1;

                self.apply_entry(self.state.volatile.commit_index).await;
            }
        }
    }

    async fn handle_request_vote(&mut self, header: RaftMessageHeader, args: RequestVoteArgs) {
        let vote_granted = self.should_grant_vote(&header, args);

        self.send(
            &header.source,
            RaftMessage {
                header: self.get_header(),
                content: RaftMessageContent::RequestVoteResponse(RequestVoteResponseArgs {
                    vote_granted,
                }),
            },
        )
        .await;

        if vote_granted {
            self.state.persistent.voted_for = Some(header.source);
        }

        self.reset_election_timer();
    }

    fn should_grant_vote(&self, header: &RaftMessageHeader, args: RequestVoteArgs) -> bool {
        let log_wise = match self.state.log_last() {
            None => true,
            Some(log) => {
                log.term < args.last_log_term
                    || (log.term == args.last_log_term
                        && self.state.log_last_index() <= args.last_log_index)
            }
        };

        let voted_wise = match self.state.persistent.voted_for {
            None => true,
            Some(candidate_id) => candidate_id == header.source,
        };

        voted_wise && self.state.persistent.current_term <= header.term && log_wise
    }

    async fn handle_request_vote_response(
        &mut self,
        header: RaftMessageHeader,
        args: RequestVoteResponseArgs,
    ) {
        if let ProcessType::Candidate { votes_received } = &mut self.process_type {
            if args.vote_granted {
                votes_received.insert(header.source);

                if votes_received.len() > self.state.get_servers().len() / 2 {
                    self.become_leader().await
                }
            }
        } else {
            log::debug!("Got {:?} as {:?}", args, self.process_type);
        }
    }

    async fn handle_install_snapshot(
        &mut self,
        header: RaftMessageHeader,
        args: InstallSnapshotArgs,
    ) {
        let last_included_index = args.last_included_index;

        let done = args.done;
        let args_offset = args.offset;

        let offset = self.state.add_snapshot_chunk(args);

        if done && args_offset <= offset {
            let snapshot_data = self.state.apply_snapshot(last_included_index).await;

            if let Some(serialized) = snapshot_data {
                self.state_machine_mod_ref
                    .send(StateMachineCommand::Initialize { serialized })
                    .await;
            }
        }

        let msg = RaftMessage {
            header: self.get_header(),
            content: RaftMessageContent::InstallSnapshotResponse(InstallSnapshotResponseArgs {
                last_included_index,
                offset,
            }),
        };

        self.send(&header.source, msg).await;
    }

    async fn handle_install_snapshot_response(
        &mut self,
        header: RaftMessageHeader,
        args: InstallSnapshotResponseArgs,
    ) {
        if let Some(last_index) = self.state.get_snapshot_last_included_idx() {
            if args.last_included_index == last_index {
                let done = self.state.set_snapshot_offset(header.source, args.offset);

                if done {
                    if let ProcessType::Leader {
                        match_index,
                        next_index,
                        cluster_change,
                        ..
                    } = &mut self.process_type
                    {
                        if let ClusterChange::CatchUp(CatchUpState {
                            reply_to: _,
                            new_server,
                            round: _,
                            idx_to_catchup,
                            next_idx,
                            match_idx,
                        }) = cluster_change
                        {
                            if *new_server == header.source {
                                *next_idx = args.last_included_index + 1;
                                *match_idx = args.last_included_index;

                                if *match_idx >= *idx_to_catchup {
                                    self.new_catchup_round().await;
                                }

                                self.catch_up().await;

                                return;
                            }
                        }

                        match_index.insert(header.source, args.last_included_index);
                        next_index.insert(header.source, args.last_included_index + 1);
                    }
                    return;
                }
            } else {
                self.state.set_snapshot_offset(header.source, 0);
            }

            self.try_send_heartbeat(&header.source).await;
        } else {
            panic!("GOT HandleInstallSnapshotResponse not having snapshot!")
        }
    }

    async fn send(&self, target: &Uuid, msg: RaftMessage) {
        if *target != self.config.self_id {
            self.message_sender.send(target, msg).await
        }
    }

    async fn broadcast(&self, msg: RaftMessage) {
        let jobs = self
            .config
            .servers
            .iter()
            .map(|server| self.send(server, msg.clone()));

        for job in jobs {
            job.await;
        }
    }

    fn log_matches(&self, args: &AppendEntriesArgs) -> bool {
        self.state.log_last_index() >= args.prev_log_index
            && self.state.log_get_term(args.prev_log_index) == args.prev_log_term
    }

    async fn update_log(&mut self, args: AppendEntriesArgs) -> usize {
        let mut idx = args.prev_log_index + 1;

        for entry in args.entries {
            if idx > self.state.log_last_index() || *self.state.log_get_entry(idx) != entry {
                self.state.log_truncate(idx);
                self.state.log_push(entry);
            }
            idx += 1;
        }

        std::cmp::min(idx - 1, args.leader_commit)
    }

    async fn update_commit_index(&mut self, last_entry_idx: usize) {
        while self.state.volatile.commit_index < last_entry_idx {
            if self.state.volatile.commit_index >= self.state.logs_in_snapshot_len() {
                self.apply_entry(self.state.volatile.commit_index).await
            }

            self.state.volatile.commit_index += 1;
        }
    }

    async fn add_to_log(
        &mut self,
        tx: Sender<ClientRequestResponse>,
        content: LogEntryContent,
        timestamp: SystemTime,
    ) -> bool {
        let res = match &mut self.process_type {
            ProcessType::Leader {
                next_index,
                match_index,
                reply_to,
                ..
            } => {
                reply_to.insert(self.state.log_last_index() + 1, tx);

                let entry = LogEntry {
                    content,
                    term: self.state.persistent.current_term,
                    timestamp,
                };

                self.state.log_push(entry);

                next_index.insert(self.config.self_id, self.state.log_last_index() + 1);
                match_index.insert(self.config.self_id, self.state.log_last_index());

                self.state.save().await;

                self.broadcast_heartbeat().await;

                true
            }
            _ => false,
        };

        res
    }

    async fn send_to_client(sender: Sender<ClientRequestResponse>, msg: ClientRequestResponse) {
        if let Err(e) = sender.send(msg).await {
            log::debug!("Error while sending response to client {}", e)
        }
    }

    async fn handle_client_command(
        &mut self,
        reply_to: Sender<ClientRequestResponse>,
        data: Vec<u8>,
        client_id: Uuid,
        sequence_num: u64,
        lowest_sequence_num_without_response: u64,
    ) {
        if !matches!(self.process_type, ProcessType::Leader { .. }) {
            Self::send_to_client(
                reply_to,
                ClientRequestResponse::CommandResponse(CommandResponseArgs {
                    client_id,
                    sequence_num,
                    content: CommandResponseContent::NotLeader {
                        leader_hint: self.state.persistent.leader_id,
                    },
                }),
            )
            .await;

            return;
        }

        let log_content = LogEntryContent::Command {
            data,
            client_id,
            sequence_num,
            lowest_sequence_num_without_response,
        };

        self.add_to_log(reply_to.clone(), log_content, SystemTime::now())
            .await;
    }

    async fn handle_client_register_client(&mut self, reply_to: Sender<ClientRequestResponse>) {
        let log_content = LogEntryContent::RegisterClient;

        if !self
            .add_to_log(reply_to.clone(), log_content, SystemTime::now())
            .await
        {
            Self::send_to_client(
                reply_to,
                ClientRequestResponse::RegisterClientResponse(RegisterClientResponseArgs {
                    content: RegisterClientResponseContent::NotLeader {
                        leader_hint: self.state.persistent.leader_id,
                    },
                }),
            )
            .await;
        }
    }

    fn should_be_dropped(&self, msg: &RaftMessage) -> bool {
        matches!(msg.content, RaftMessageContent::RequestVote(..))
            && match self.process_type {
                ProcessType::Follower { is_leader_reliable } => is_leader_reliable,
                ProcessType::Candidate { .. } => false,
                ProcessType::Leader { .. } => true,
            }
    }

    fn start_heartbeat_timer(&mut self) {
        self.heartbeat_timer
            .reset(self.get_self_ref(), HeartbeatTimeout)
    }

    fn stop_heartbeat_timer(&mut self) {
        self.heartbeat_timer.stop();
    }

    fn become_follower(&mut self) {
        self.process_type = ProcessType::Follower {
            is_leader_reliable: false,
        };
        self.stop_heartbeat_timer();
        self.stop_catchup_timer();
        self.stop_minimal_election_timer();
    }

    fn restart_minimal_election_timer(&mut self) {
        self.minimal_election_timer
            .reset(self.get_self_ref(), MinimalElectionTimeout)
    }

    fn stop_minimal_election_timer(&mut self) {
        self.minimal_election_timer.stop();
    }

    async fn apply_entry(&mut self, idx: usize) {
        let entry = self.state.log_get_entry(idx).clone();

        self.state_machine_mod_ref
            .send(StateMachineCommand::LogEntry { entry, idx })
            .await;
    }

    async fn handle_snapshot(&mut self, tx: Sender<ClientRequestResponse>) {
        self.state_machine_mod_ref
            .send(StateMachineCommand::Snapshot { tx })
            .await
    }

    async fn handle_log_entry_output(&mut self, content: StateMachineResultContent, idx: usize) {
        self.state.volatile.last_applied = idx;

        if let ProcessType::Leader { reply_to, .. } = &mut self.process_type {
            if let Some(tx) = reply_to.remove(&idx) {
                let response = match content {
                    StateMachineResultContent::RegisterClient(client_id) => {
                        ClientRequestResponse::RegisterClientResponse(RegisterClientResponseArgs {
                            content: RegisterClientResponseContent::ClientRegistered { client_id },
                        })
                    }
                    StateMachineResultContent::Command(cmd) => {
                        self.handle_state_machine_result_command(cmd, idx).await
                    }
                    StateMachineResultContent::Configuration(servers) => {
                        if let Some(msg) = self
                            .handle_state_machine_result_configuration(servers)
                            .await
                        {
                            msg
                        } else {
                            return;
                        }
                    }
                };

                Self::send_to_client(tx, response).await;
            }
        }
    }

    async fn handle_snapshot_output(
        &mut self,
        output: Option<SnapshotOutput>,
        tx: Sender<ClientRequestResponse>,
    ) {
        let content = self.state.update_snapshot(output).await;

        tx.send(ClientRequestResponse::SnapshotResponse(
            SnapshotResponseArgs { content },
        ))
        .await
        .unwrap();
    }

    async fn handle_remove_server(
        &mut self,
        reply_to: Sender<ClientRequestResponse>,
        old_server: Uuid,
    ) {
        if let ProcessType::Leader { cluster_change, .. } = &mut self.process_type {
            if let ClusterChange::None = cluster_change {
                let mut servers = self.state.get_servers().clone();
                if servers.remove(&old_server) {
                    *cluster_change = ClusterChange::RemoveServer { old_server };
                    self.add_config_change_to_log(reply_to, servers).await;
                } else {
                    Self::send_to_client(
                        reply_to,
                        ClientRequestResponse::RemoveServerResponse(RemoveServerResponseArgs {
                            old_server,
                            content: RemoveServerResponseContent::NotPresent,
                        }),
                    )
                    .await;
                }
            } else {
                Self::send_to_client(
                    reply_to,
                    ClientRequestResponse::RemoveServerResponse(RemoveServerResponseArgs {
                        old_server,
                        content: RemoveServerResponseContent::RemoveInProgress,
                    }),
                )
                .await;
            }
        } else {
            Self::send_to_client(
                reply_to,
                ClientRequestResponse::RemoveServerResponse(RemoveServerResponseArgs {
                    old_server,
                    content: RemoveServerResponseContent::NotLeader {
                        leader_hint: self.state.persistent.leader_id,
                    },
                }),
            )
            .await;
        };
    }

    async fn handle_add_server(
        &mut self,
        reply_to: Sender<ClientRequestResponse>,
        new_server: Uuid,
    ) {
        if let ProcessType::Leader { cluster_change, .. } = &mut self.process_type {
            if let ClusterChange::None = cluster_change {
                if self.state.get_servers().contains(&new_server) {
                    Self::send_to_client(
                        reply_to,
                        ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
                            new_server,
                            content: AddServerResponseContent::AlreadyPresent,
                        }),
                    )
                    .await;
                } else {
                    *cluster_change = ClusterChange::CatchUp(CatchUpState {
                        reply_to,
                        new_server,
                        round: 0,
                        idx_to_catchup: 0,
                        next_idx: self.state.log_last_index() + 1,
                        match_idx: 0,
                    });
                    self.new_catchup_round().await;
                }
            } else {
                Self::send_to_client(
                    reply_to,
                    ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
                        new_server,
                        content: AddServerResponseContent::AddInProgress,
                    }),
                )
                .await;
            }
        } else {
            Self::send_to_client(
                reply_to,
                ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
                    new_server,
                    content: AddServerResponseContent::NotLeader {
                        leader_hint: self.state.persistent.leader_id,
                    },
                }),
            )
            .await;
        };
    }

    async fn catch_up(&mut self) {
        if let ProcessType::Leader {
            cluster_change:
                ClusterChange::CatchUp(CatchUpState {
                    new_server,
                    match_idx,
                    next_idx,
                    round,
                    ..
                }),
            ..
        } = self.process_type
        {
            let new_server_clone = new_server.clone();
            let next_index_clone = next_idx.clone();
            let match_index_clone = match_idx.clone();

            self.send_heartbeat(&new_server_clone, next_index_clone, match_index_clone)
                .await;

            if round < self.config.catch_up_rounds {
                self.reset_catchup_timer();
            }
        }
    }

    async fn add_config_change_to_log(
        &mut self,
        reply_to: Sender<ClientRequestResponse>,
        servers: HashSet<Uuid>,
    ) {
        self.add_to_log(
            reply_to,
            LogEntryContent::Configuration { servers },
            SystemTime::now(),
        )
        .await;

        self.state.volatile.last_config_change_idx = self.state.log_last_index();
    }

    async fn handle_state_machine_result_configuration(
        &mut self,
        servers: HashSet<Uuid>,
    ) -> Option<ClientRequestResponse> {
        self.state.config_confirm(servers);
        if let ProcessType::Leader { cluster_change, .. } = &mut self.process_type {
            match cluster_change.clone() {
                ClusterChange::AddServer { new_server } => {
                    *cluster_change = ClusterChange::None;

                    Some(ClientRequestResponse::AddServerResponse(
                        AddServerResponseArgs {
                            new_server,
                            content: AddServerResponseContent::ServerAdded,
                        },
                    ))
                }
                ClusterChange::RemoveServer { old_server } => {
                    *cluster_change = ClusterChange::None;

                    if old_server == self.config.self_id {
                        self.become_follower();
                    }

                    Some(ClientRequestResponse::RemoveServerResponse(
                        RemoveServerResponseArgs {
                            old_server,
                            content: RemoveServerResponseContent::ServerRemoved,
                        },
                    ))
                }
                ClusterChange::CommitOldConfig { idx } => {
                    if self.state.volatile.last_applied >= idx {
                        *cluster_change = ClusterChange::None;
                    }

                    None
                }
                _ => None,
            }
        } else {
            None
        }
    }

    async fn handle_state_machine_result_command(
        &self,
        cmd: StateMachineResultContentCommand,
        idx: usize,
    ) -> ClientRequestResponse {
        let entry = self.state.log_get_entry(idx);
        if let LogEntryContent::Command {
            client_id,
            sequence_num,
            ..
        } = entry.content
        {
            let content = match cmd {
                StateMachineResultContentCommand::SessionExpired => {
                    CommandResponseContent::SessionExpired
                }
                StateMachineResultContentCommand::Applied(output) => {
                    CommandResponseContent::CommandApplied { output }
                }
            };

            ClientRequestResponse::CommandResponse(CommandResponseArgs {
                client_id,
                sequence_num,
                content,
            })
        } else {
            panic!();
        }
    }

    async fn new_catchup_round(&mut self) {
        if let ProcessType::Leader {
            cluster_change,
            next_index,
            match_index,
            ..
        } = &mut self.process_type
        {
            if let ClusterChange::CatchUp(CatchUpState {
                reply_to,
                new_server,
                round,
                idx_to_catchup,
                next_idx,
                match_idx,
            }) = cluster_change
            {
                *idx_to_catchup = self.state.log_last_index();
                *round += 1;

                if *round > self.config.catch_up_rounds {
                    let reply_to_clone = reply_to.clone();

                    let mut servers = self.state.get_servers().clone();
                    servers.insert(new_server.clone());

                    next_index.insert(new_server.clone(), *next_idx);
                    match_index.insert(new_server.clone(), *match_idx);

                    *cluster_change = ClusterChange::AddServer {
                        new_server: new_server.clone(),
                    };
                    self.stop_catchup_timer();
                    self.add_config_change_to_log(reply_to_clone, servers).await;
                } else {
                    self.catch_up().await;
                }
            }
        }
    }
}

#[async_trait::async_trait]
impl Handler<RaftMessage> for Raft {
    async fn handle(&mut self, msg: RaftMessage) {
        if self.enabled {
            log::debug!("{:?} got {:?}", self.config.self_id, msg);
            if self.should_be_dropped(&msg) {
                return;
            }

            if msg.header.term > self.state.persistent.current_term {
                self.update_term(msg.header.term);
                self.become_follower();
            }

            match msg.content {
                RaftMessageContent::AppendEntries(args) => {
                    self.handle_append_entries(msg.header, args).await
                }
                RaftMessageContent::AppendEntriesResponse(args) => {
                    self.handle_append_entries_response(msg.header, args).await
                }
                RaftMessageContent::RequestVote(args) => {
                    self.handle_request_vote(msg.header, args).await
                }
                RaftMessageContent::RequestVoteResponse(args) => {
                    self.handle_request_vote_response(msg.header, args).await
                }
                RaftMessageContent::InstallSnapshot(args) => {
                    self.handle_install_snapshot(msg.header, args).await
                }
                RaftMessageContent::InstallSnapshotResponse(args) => {
                    self.handle_install_snapshot_response(msg.header, args)
                        .await
                }
            }
        }
    }
}

#[async_trait::async_trait]
impl Handler<ClientRequest> for Raft {
    async fn handle(&mut self, msg: ClientRequest) {
        log::debug!("{:?} got CLIENT {:?}", self.config.self_id, msg);
        match msg.content {
            ClientRequestContent::Command {
                command,
                client_id,
                sequence_num,
                lowest_sequence_num_without_response,
            } => {
                self.handle_client_command(
                    msg.reply_to,
                    command,
                    client_id,
                    sequence_num,
                    lowest_sequence_num_without_response,
                )
                .await
            }
            ClientRequestContent::Snapshot => self.handle_snapshot(msg.reply_to).await,
            ClientRequestContent::AddServer { new_server } => {
                self.handle_add_server(msg.reply_to, new_server).await
            }
            ClientRequestContent::RemoveServer { old_server } => {
                self.handle_remove_server(msg.reply_to, old_server).await
            }
            ClientRequestContent::RegisterClient => {
                self.handle_client_register_client(msg.reply_to).await
            }
        }
    }
}

#[derive(Clone)]
struct ElectionTimeout;

#[derive(Clone)]
struct HeartbeatTimeout;

#[derive(Clone)]
struct MinimalElectionTimeout;

#[derive(Clone)]
struct CatchUpTimeout;

struct Init {
    raft_ref: ModuleRef<Raft>,
}

/// Message disabling a process. Used for testing to simulate failures.
pub(crate) struct Disable;

#[async_trait::async_trait]
impl Handler<Init> for Raft {
    async fn handle(&mut self, msg: Init) {
        if self.enabled {
            self.self_ref = Some(msg.raft_ref);

            self.reset_election_timer();
        }
    }
}

/// Handle timer timeout.
#[async_trait::async_trait]
impl Handler<ElectionTimeout> for Raft {
    async fn handle(&mut self, _: ElectionTimeout) {
        if self.enabled {
            match &mut self.process_type {
                ProcessType::Leader { responded, .. } => {
                    if responded.len() > self.state.get_servers().len() / 2 {
                        *responded = Self::get_fresh_responded(self.config.self_id);
                    } else {
                        self.become_follower();
                        self.state.persistent.leader_id = None;
                        self.state.save().await;
                    }
                }
                _ => {
                    self.become_candidate().await;
                }
            }
        }
    }
}

/// Handle timer timeout.
#[async_trait::async_trait]
impl Handler<CatchUpTimeout> for Raft {
    async fn handle(&mut self, _: CatchUpTimeout) {
        if self.enabled {
            if let ProcessType::Leader { cluster_change, .. } = &mut self.process_type {
                if let ClusterChange::CatchUp(CatchUpState {
                    reply_to,
                    new_server,
                    ..
                }) = cluster_change
                {
                    let reply_to_clone = reply_to.clone();
                    let new_server_clone = new_server.clone();

                    *cluster_change = ClusterChange::None;

                    self.stop_catchup_timer();

                    Self::send_to_client(
                        reply_to_clone,
                        ClientRequestResponse::AddServerResponse(AddServerResponseArgs {
                            new_server: new_server_clone,
                            content: AddServerResponseContent::Timeout,
                        }),
                    )
                    .await;
                }
            }
        }
    }
}

/// Handle timer timeout.
#[async_trait::async_trait]
impl Handler<MinimalElectionTimeout> for Raft {
    async fn handle(&mut self, _: MinimalElectionTimeout) {
        if self.enabled {
            match &mut self.process_type {
                ProcessType::Follower { is_leader_reliable } => {
                    *is_leader_reliable = false;
                    self.stop_minimal_election_timer();
                }
                _ => {}
            }
        }
    }
}

/// Handle timer timeout.
#[async_trait::async_trait]
impl Handler<HeartbeatTimeout> for Raft {
    async fn handle(&mut self, _: HeartbeatTimeout) {
        if self.enabled {
            match &self.process_type {
                ProcessType::Leader { .. } => {
                    self.broadcast_heartbeat().await;
                    self.try_apply().await;

                    self.catch_up().await;
                }
                _ => {}
            }
        }
    }
}

#[async_trait::async_trait]
impl Handler<Disable> for Raft {
    async fn handle(&mut self, _: Disable) {
        self.enabled = false;
    }
}

#[async_trait::async_trait]
impl Handler<StateMachineResult> for Raft {
    async fn handle(&mut self, result: StateMachineResult) {
        if self.enabled {
            match result {
                StateMachineResult::LogEntry { content, idx } => {
                    self.handle_log_entry_output(content, idx).await
                }
                StateMachineResult::Snapshot { output, tx } => {
                    self.handle_snapshot_output(output, tx).await
                }
            }
        }
    }
}

/// State of a Raft process with a corresponding (volatile) information.
#[derive(Debug)]
enum ProcessType {
    Follower {
        is_leader_reliable: bool,
    },
    Candidate {
        votes_received: HashSet<Uuid>,
    },
    Leader {
        responded: HashSet<Uuid>,
        reply_to: HashMap<usize, Sender<ClientRequestResponse>>,
        next_index: HashMap<Uuid, usize>,
        match_index: HashMap<Uuid, usize>,
        cluster_change: ClusterChange,
    },
}

#[derive(Debug, Clone)]
enum ClusterChange {
    CatchUp(CatchUpState),
    AddServer { new_server: Uuid },
    RemoveServer { old_server: Uuid },
    CommitOldConfig { idx: usize },
    None,
}

#[derive(Debug, Clone)]
struct CatchUpState {
    reply_to: Sender<ClientRequestResponse>,
    new_server: Uuid,
    round: u64,
    idx_to_catchup: usize,
    next_idx: usize,
    match_idx: usize,
}
