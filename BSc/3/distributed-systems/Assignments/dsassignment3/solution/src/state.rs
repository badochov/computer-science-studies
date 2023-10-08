use crate::*;
use serde::{Deserialize, Serialize};

#[derive(Default, Clone, Serialize, Deserialize)]
pub struct PersistentState {
    pub current_term: u64,
    pub voted_for: Option<Uuid>,
    log: Vec<LogEntry>,
    pub leader_id: Option<Uuid>,
}

#[derive(Default)]
pub struct VolatileState {
    pub commit_index: usize,
    pub last_applied: usize,
    pub last_config_change_idx: usize,
    last_applied_config: HashSet<Uuid>,
    snapshot_offset: HashMap<Uuid, usize>,
    partial_snapshots: HashMap<usize, SnapshotData>,
}

/// State of a Raft process.
/// It shall be kept in stable storage, and updated before replying to messages.
pub struct State {
    pub persistent: PersistentState,
    pub volatile: VolatileState,
    snapshot: Option<SnapshotData>,
    servers: HashSet<Uuid>,
    stable_storage: Box<dyn StableStorage>,
    snapshot_chunk_size: usize,
}

#[derive(Default, Clone, Serialize, Deserialize)]
pub struct SnapshotData {
    last_index: usize,
    last_term: u64,
    serialized: Vec<u8>,
    servers: HashSet<Uuid>,
    client_sessions: HashMap<Uuid, ClientSession>,
}

impl SnapshotData {
    fn from_install_snapshot_args(args: InstallSnapshotArgs) -> Self {
        Self {
            last_index: args.last_included_index,
            last_term: args.last_included_term,
            serialized: args.data,
            servers: args.last_config.unwrap(),
            client_sessions: args.client_sessions.unwrap(),
        }
    }
}

impl PersistentState {
    pub async fn restore(
        stable_storage: &Box<dyn StableStorage>,
        first_log_entry_timestamp: SystemTime,
        servers: HashSet<Uuid>,
    ) -> Self {
        match stable_storage.get(&Self::storage_key()).await {
            None => {
                let mut state = Self::default();

                state.log.push(LogEntry {
                    term: 0,
                    timestamp: first_log_entry_timestamp,
                    content: LogEntryContent::Configuration { servers },
                });

                state
            }
            Some(data) => bincode::deserialize(&data).unwrap(),
        }
    }

    fn storage_key() -> String {
        "persistent_state".to_owned()
    }

    async fn save(&self, stable_storage: &mut Box<dyn StableStorage>) {
        stable_storage
            .put(&Self::storage_key(), &bincode::serialize(&self).unwrap())
            .await
            .unwrap()
    }
}

impl State {
    pub async fn new(
        stable_storage: Box<dyn StableStorage>,
        config: &ServerConfig,
        first_log_entry_timestamp: SystemTime,
    ) -> Self {
        let persistent = PersistentState::restore(
            &stable_storage,
            first_log_entry_timestamp,
            config.servers.clone(),
        )
        .await;

        let mut volatile = VolatileState::default();

        let snapshot = SnapshotData::restore(&stable_storage).await;

        if let Some(s) = &snapshot {
            volatile.last_applied_config = s.servers.clone();
        } else {
            volatile.last_applied_config = config.servers.clone();
        }

        let mut res = Self {
            persistent,
            volatile,
            snapshot,
            servers: HashSet::new(),
            snapshot_chunk_size: config.snapshot_chunk_size,
            stable_storage,
        };

        res.set_newest_config();

        res
    }

    pub fn get_snapshot_last_included_idx(&self) -> Option<usize> {
        self.snapshot.as_ref().map(|s| s.last_index)
    }

    pub fn get_snapshot_offset(&self, client_id: &Uuid) -> usize {
        *self.volatile.snapshot_offset.get(client_id).unwrap_or(&0)
    }

    pub fn set_snapshot_offset(&mut self, client_id: Uuid, offset: usize) -> bool {
        self.volatile.snapshot_offset.insert(client_id, offset);

        if let Some(snapshot) = &self.snapshot {
            snapshot.serialized.len() == offset
        } else {
            false
        }
    }

    pub fn get_install_snapshot_args(&self, offset: usize) -> InstallSnapshotArgs {
        if let Some(snapshot) = &self.snapshot {
            let end_idx = offset + self.snapshot_chunk_size;

            let (data, done) = if end_idx >= snapshot.serialized.len() {
                (snapshot.serialized[offset..].to_vec(), true)
            } else {
                (snapshot.serialized[offset..end_idx].to_vec(), false)
            };

            let (client_sessions, last_config) = if offset == 0 {
                (
                    Some(snapshot.client_sessions.clone()),
                    Some(snapshot.servers.clone()),
                )
            } else {
                (None, None)
            };

            return InstallSnapshotArgs {
                last_included_index: snapshot.last_index,
                last_included_term: snapshot.last_term,
                last_config,
                client_sessions,
                offset,
                data,
                done,
            };
        }

        panic!("Snapshot not initialized");
    }

    /// Reliably save the state.
    pub async fn save(&mut self) {
        self.persistent.save(&mut self.stable_storage).await;
    }

    pub fn log_last_index(&self) -> usize {
        self.logs_in_snapshot_len() + self.persistent.log.len() - 1
    }

    pub fn log_last_term(&self) -> u64 {
        self.log_last()
            .map_or(self.snapshot.as_ref().map_or(0, |s| s.last_term), |e| {
                e.term
            })
    }

    pub fn log_last(&self) -> Option<&LogEntry> {
        self.persistent.log.last()
    }

    pub fn log_push(&mut self, entry: LogEntry) {
        if let LogEntryContent::Configuration { servers } = &entry.content {
            self.servers = servers.clone();
        }

        self.persistent.log.push(entry);
    }

    pub fn log_get_term(&self, idx: usize) -> u64 {
        if let Some(snapshot) = &self.snapshot {
            if idx <= snapshot.last_index {
                return snapshot.last_term;
            }
        }

        self.log_get_entry(idx).term
    }

    pub fn log_get_entry(&self, idx: usize) -> &LogEntry {
        &self.persistent.log[self.map_index(idx)]
    }

    pub fn log_truncate(&mut self, len: usize) {
        let mut removed = self.persistent.log.drain(self.map_index(len)..);

        if removed.any(|entry| matches!(entry.content, LogEntryContent::Configuration { .. })) {
            std::mem::drop(removed);
            self.set_newest_config();
        }
    }

    pub fn log_slice(&self, start: usize, end: usize) -> Option<Vec<LogEntry>> {
        if let Some(snapshot) = &self.snapshot {
            if snapshot.last_index >= start {
                return None;
            }
        }
        Some(self.persistent.log[self.map_index(start)..self.map_index(end)].to_vec())
    }

    fn map_index(&self, idx: usize) -> usize {
        idx - self.logs_in_snapshot_len()
    }

    pub fn logs_in_snapshot_len(&self) -> usize {
        self.snapshot.as_ref().map_or(0, |s| s.last_index + 1)
    }

    pub fn get_servers(&self) -> &HashSet<Uuid> {
        &self.servers
    }

    pub async fn update_snapshot(
        &mut self,
        output: Option<SnapshotOutput>,
    ) -> SnapshotResponseContent {
        let last_included_index = self.volatile.last_applied;

        match output {
            None => SnapshotResponseContent::NothingToSnapshot {
                last_included_index,
            },
            Some(snapshot_output) => {
                if let None = &self.snapshot {
                    self.snapshot = Some(SnapshotData::default());
                }
                if let Some(snapshot) = &mut self.snapshot {
                    let to_remove = last_included_index + 1 - snapshot.last_index;

                    snapshot.serialized = snapshot_output.serialized;
                    snapshot.client_sessions = snapshot_output.client_sessions;
                    snapshot.servers = self.volatile.last_applied_config.clone();
                    snapshot.last_index = last_included_index;
                    snapshot.last_term = self.persistent.log[last_included_index].term;

                    self.persistent.log.drain(0..to_remove);

                    self.volatile.snapshot_offset.clear();

                    snapshot.save(&mut self.stable_storage).await;
                }

                SnapshotResponseContent::SnapshotCreated {
                    last_included_index,
                }
            }
        }
    }

    pub fn get_snapshot_data(&self) -> Option<Vec<u8>> {
        self.snapshot
            .as_ref()
            .map(|snapshot| snapshot.serialized.clone())
    }

    pub fn add_snapshot_chunk(&mut self, args: InstallSnapshotArgs) -> usize {
        if let Some(snapshot) = self
            .volatile
            .partial_snapshots
            .get_mut(&args.last_included_index)
        {
            if snapshot.serialized.len() == args.offset {
                snapshot.serialized.extend(args.data);
            }
            return snapshot.serialized.len();
        }

        if args.offset == 0 {
            let offset = args.data.len();

            self.volatile.partial_snapshots.insert(
                args.last_included_index,
                SnapshotData::from_install_snapshot_args(args),
            );

            offset
        } else {
            0
        }
    }

    pub async fn apply_snapshot(&mut self, last_included_index: usize) -> Option<Vec<u8>> {
        if let Some(snapshot) = self.volatile.partial_snapshots.remove(&last_included_index) {
            let serialized = snapshot.serialized.clone();

            let to_remove = snapshot
                .last_index
                .checked_sub(self.snapshot.as_ref().map_or(0, |s| s.last_index));

            if let Some(count) = to_remove {
                self.persistent
                    .log
                    .drain(0..std::cmp::min(count, self.persistent.log.len()));
            } else {
                return None;
            }

            snapshot.save(&mut self.stable_storage).await;
            self.snapshot = Some(snapshot);

            self.volatile.snapshot_offset.clear();
            self.volatile.partial_snapshots.clear();

            Some(serialized)
        } else {
            panic!("Tried to apply unexisting snapshot");
        }
    }

    pub fn config_confirm(&mut self, servers: HashSet<Uuid>) {
        self.volatile.last_applied_config = servers;
    }

    fn set_newest_config(&mut self) {
        let mut found = false;

        for (idx, entry) in self.persistent.log.iter().enumerate().rev() {
            if let LogEntryContent::Configuration { servers } = &entry.content {
                self.volatile.last_config_change_idx = idx + self.logs_in_snapshot_len();
                self.servers = servers.clone();
                found = true;
                break;
            }
        }

        if !found {
            if let Some(s) = &self.snapshot {
                self.servers = s.servers.clone();
            }
        }
    }
}

impl SnapshotData {
    async fn restore(stable_storage: &Box<dyn StableStorage>) -> Option<SnapshotData> {
        stable_storage
            .get(&Self::storage_key())
            .await
            .map(|data| bincode::deserialize(&data).unwrap())
    }

    fn storage_key() -> String {
        "snapshot_data".to_owned()
    }

    async fn save(&self, stable_storage: &mut Box<dyn StableStorage>) {
        stable_storage
            .put(&Self::storage_key(), &bincode::serialize(&self).unwrap())
            .await
            .unwrap()
    }
}
