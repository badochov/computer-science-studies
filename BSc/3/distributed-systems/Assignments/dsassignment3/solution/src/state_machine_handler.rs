use std::{
    collections::{BTreeMap, HashMap},
    time::{Duration, SystemTime},
};

use executor::ModuleRef;
use uuid::Uuid;

use crate::*;

pub enum StateMachineCommand {
    LogEntry { entry: LogEntry, idx: usize },
    Snapshot { tx: Sender<ClientRequestResponse> },
    Initialize { serialized: Vec<u8> },
}
pub enum StateMachineResult {
    LogEntry {
        content: StateMachineResultContent,
        idx: usize,
    },
    Snapshot {
        output: Option<SnapshotOutput>,
        tx: Sender<ClientRequestResponse>,
    },
}

pub struct SnapshotOutput {
    pub serialized: Vec<u8>,
    pub client_sessions: HashMap<Uuid, ClientSession>,
}

pub enum StateMachineResultContent {
    RegisterClient(Uuid),
    Command(StateMachineResultContentCommand),
    Configuration(HashSet<Uuid>),
}

pub enum StateMachineResultContentCommand {
    SessionExpired,
    Applied(Vec<u8>),
}

pub struct StateMachineHandler {
    state_machine: Box<dyn StateMachine>,
    mod_ref: Option<ModuleRef<Raft>>,
    client_sessions: HashMap<Uuid, SessionData>,
    session_expiration: Duration,
    was_snapshot_last: bool,
}

pub struct InitStateMachine {
    pub raft_ref: ModuleRef<Raft>,
    pub serialized: Option<Vec<u8>>,
}

impl StateMachineHandler {
    pub async fn new(
        system: &mut System,
        state_machine: Box<dyn StateMachine>,
        session_expiration: Duration,
    ) -> ModuleRef<Self> {
        system
            .register_module(Self {
                was_snapshot_last: false,
                state_machine,
                mod_ref: None,
                client_sessions: HashMap::new(),
                session_expiration,
            })
            .await
    }

    fn get_prev_result(&self, client_id: Uuid, sequence_num: u64) -> Option<Vec<u8>> {
        if let Some(session) = self.client_sessions.get(&client_id) {
            return session.results.get(&sequence_num).map(|res| res.clone());
        }
        None
    }

    fn expire_sessions(&mut self, timestamp: SystemTime) {
        let mut expired_sessions = Vec::new();

        for (uuid, session_data) in self.client_sessions.iter() {
            if !session_data.is_active(timestamp, self.session_expiration) {
                expired_sessions.push(uuid.clone());
            }
        }

        for session in expired_sessions {
            self.client_sessions.remove(&session);
        }
    }

    fn update_client_session(
        &mut self,
        timestamp: SystemTime,
        client_id: Uuid,
        lowest_sequence_num_without_response: u64,
    ) {
        if let Some(session) = self.client_sessions.get_mut(&client_id) {
            session.last_activity = timestamp;
            session.lowest_sequence_num_without_response = lowest_sequence_num_without_response;

            self.remove_confirmed_results(client_id, lowest_sequence_num_without_response);
        }
    }

    fn remove_confirmed_results(
        &mut self,
        client_id: Uuid,
        lowest_sequence_num_without_response: u64,
    ) {
        if let Some(session) = self.client_sessions.get_mut(&client_id) {
            session.results = session
                .results
                .split_off(&lowest_sequence_num_without_response);
        }
    }

    async fn handle_entry(&mut self, entry: LogEntry, idx: usize) {
        self.was_snapshot_last = false;

        let content = match entry.content {
            LogEntryContent::Command {
                data,
                client_id,
                sequence_num,
                lowest_sequence_num_without_response,
            } => {
                if let Some(session) = self.client_sessions.get(&client_id) {
                    if session.lowest_sequence_num_without_response > sequence_num {
                        self.client_sessions.remove(&client_id);

                        StateMachineResultContent::Command(
                            StateMachineResultContentCommand::SessionExpired,
                        )
                    } else {
                        self.update_client_session(
                            entry.timestamp,
                            client_id,
                            lowest_sequence_num_without_response,
                        );
                        let output =
                            if let Some(res) = self.get_prev_result(client_id, sequence_num) {
                                res
                            } else {
                                let output = self.state_machine.apply(&data).await;
                                if let Some(session) = self.client_sessions.get_mut(&client_id) {
                                    session.results.insert(sequence_num, output.clone());
                                }

                                output
                            };

                        StateMachineResultContent::Command(
                            StateMachineResultContentCommand::Applied(output),
                        )
                    }
                } else {
                    StateMachineResultContent::Command(
                        StateMachineResultContentCommand::SessionExpired,
                    )
                }
            }
            LogEntryContent::Configuration { servers } => {
                StateMachineResultContent::Configuration(servers)
            }
            LogEntryContent::RegisterClient => {
                let client_id = Uuid::from_u128(idx as u128);

                self.client_sessions
                    .insert(client_id, SessionData::new(entry.timestamp));

                StateMachineResultContent::RegisterClient(client_id)
            }
        };

        self.mod_ref
            .clone()
            .unwrap()
            .send(StateMachineResult::LogEntry { content, idx })
            .await;

        self.expire_sessions(entry.timestamp);
    }

    async fn handle_snapshot(&mut self, tx: Sender<ClientRequestResponse>) {
        if self.was_snapshot_last {
            self.mod_ref
                .clone()
                .unwrap()
                .send(StateMachineResult::Snapshot { output: None, tx })
                .await;
        } else {
            let serialized = self.state_machine.serialize().await;
            let client_sessions = self
                .client_sessions
                .iter()
                .map(|(&uuid, session)| (uuid, session.to_client_session()))
                .collect();

            self.mod_ref
                .clone()
                .unwrap()
                .send(StateMachineResult::Snapshot {
                    output: Some(SnapshotOutput {
                        serialized,
                        client_sessions,
                    }),
                    tx,
                })
                .await;
        }

        self.was_snapshot_last = true;
    }

    async fn handle_initialize(&mut self, serialized: Vec<u8>) {
        self.state_machine.initialize(&serialized).await;
    }
}

#[async_trait::async_trait]
impl Handler<InitStateMachine> for StateMachineHandler {
    async fn handle(&mut self, msg: InitStateMachine) {
        self.mod_ref = Some(msg.raft_ref);
        if let Some(serialized) = &msg.serialized {
            self.state_machine.initialize(serialized).await;
        }
    }
}

#[async_trait::async_trait]
impl Handler<StateMachineCommand> for StateMachineHandler {
    async fn handle(&mut self, cmd: StateMachineCommand) {
        match cmd {
            StateMachineCommand::LogEntry { entry, idx } => self.handle_entry(entry, idx).await,
            StateMachineCommand::Snapshot { tx } => self.handle_snapshot(tx).await,
            StateMachineCommand::Initialize { serialized } => {
                self.handle_initialize(serialized).await
            }
        }
    }
}

#[derive(Clone)]
pub struct SessionData {
    pub lowest_sequence_num_without_response: u64,
    pub last_activity: SystemTime,
    pub results: BTreeMap<u64, Vec<u8>>,
}

impl SessionData {
    pub fn new(last_activity: SystemTime) -> Self {
        Self {
            last_activity,
            lowest_sequence_num_without_response: 0,
            results: BTreeMap::new(),
        }
    }

    pub fn is_active(&self, timestamp: SystemTime, expiry_duration: Duration) -> bool {
        if let Ok(duration) = timestamp.duration_since(self.last_activity) {
            return duration < expiry_duration;
        }

        true
    }

    pub fn to_client_session(&self) -> ClientSession {
        let responses = self.results.clone().into_iter().collect();

        ClientSession {
            last_activity: self.last_activity,
            responses,
            lowest_sequence_num_without_response: self.lowest_sequence_num_without_response,
        }
    }
}
