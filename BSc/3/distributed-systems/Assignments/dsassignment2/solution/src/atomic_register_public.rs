use crate::Broadcast;
use crate::OperationReturn;
use crate::ReadReturn;
use crate::SectorVec;
use crate::SystemCommandHeader;
use std::collections::HashSet;
use std::convert::TryInto;
use std::future::Future;
use std::pin::Pin;
use std::sync::Arc;
use uuid::Uuid;

use crate::{
    ClientRegisterCommand, ClientRegisterCommandContent, OperationComplete, RegisterClient,
    SectorIdx, SectorsManager, StableStorage, StatusCode, SystemRegisterCommand,
    SystemRegisterCommandContent,
};

#[async_trait::async_trait]
pub trait AtomicRegister: Send + Sync {
    /// Send client command to the register. After it is completed, we expect
    /// callback to be called. Note that completion of client command happens after
    /// delivery of multiple system commands to the register, as the algorithm specifies.
    async fn client_command(&mut self, cmd: ClientRegisterCommand, operation_complete: Callback);

    /// Send system command to the register.
    async fn system_command(&mut self, cmd: SystemRegisterCommand);
}

pub type Callback =
    Box<dyn FnOnce(OperationComplete) -> Pin<Box<dyn Future<Output = ()> + Send>> + Send + Sync>;

/// Idents are numbered starting at 1 (up to the number of processes in the system).
/// Storage for atomic register algorithm data is separated into StableStorage.
/// Communication with other processes of the system is to be done by register_client.
/// And sectors must be stored in the sectors_manager instance.
pub async fn build_atomic_register(
    self_ident: u8,
    metadata: Box<dyn StableStorage>,
    register_client: Arc<dyn RegisterClient>,
    sectors_manager: Arc<dyn SectorsManager>,
    processes_count: usize,
) -> Box<dyn AtomicRegister> {
    tokio::task::spawn(async move {
        let res = AtomicRegisterImpl::new(
            self_ident,
            metadata,
            register_client,
            sectors_manager,
            processes_count,
        );

        Box::new(res)
    })
    .await
    .unwrap()
}

struct AtomicRegisterImpl {
    self_ident: u8,
    metadata: Box<dyn StableStorage>,
    register_client: Arc<dyn RegisterClient>,
    sectors_manager: Arc<dyn SectorsManager>,
    processes_count: usize,

    data: AtomicRegisterData,
}

struct AtomicRegisterData {
    rid: u64,
    readlist: HashSet<u8>,
    acklist: HashSet<u8>,
    highest: Option<((u64, u8), SectorVec)>,
    reading: bool,
    writing: bool,
    readval: Option<SectorVec>,
    writeval: Option<SectorVec>,
    write_phase: bool,
    request_identifier: Option<u64>,
    callback: Option<Callback>,
}

impl AtomicRegisterData {
    pub fn new() -> Self {
        Self {
            rid: 0,
            readlist: HashSet::new(),
            acklist: HashSet::new(),
            reading: false,
            writing: false,
            readval: None,
            writeval: None,
            write_phase: false,
            request_identifier: None,
            callback: None,
            highest: None,
        }
    }
}

impl AtomicRegisterImpl {
    pub fn new(
        self_ident: u8,
        metadata: Box<dyn StableStorage>,
        register_client: Arc<dyn RegisterClient>,
        sectors_manager: Arc<dyn SectorsManager>,
        processes_count: usize,
    ) -> Self {
        Self {
            self_ident,
            metadata,
            register_client,
            sectors_manager,
            processes_count,
            data: AtomicRegisterData::new(),
        }
    }

    async fn restore_rid(&mut self, sector_id: SectorIdx) {
        let rid_key = format!("rid_{}", sector_id);

        self.data.rid = match self.metadata.get(&rid_key).await {
            None => 0,
            Some(rid) => u64::from_be_bytes(rid.try_into().unwrap()),
        };

        self.data.rid += 1;

        self.metadata
            .put(&rid_key, &self.data.rid.to_be_bytes())
            .await
            .unwrap();
    }

    async fn restore_ts_wr_val(&mut self, sector_id: SectorIdx) {
        let metadata = self.sectors_manager.read_metadata(sector_id).await;
        let data = self.sectors_manager.read_data(sector_id).await;

        self.data.highest = Some((metadata, data));
    }
}

#[async_trait::async_trait]
impl AtomicRegister for AtomicRegisterImpl {
    /// Send client command to the register. After it is completed, we expect
    /// callback to be called. Note that completion of client command happens after
    /// delivery of multiple system commands to the register, as the algorithm specifies.
    async fn client_command(
        &mut self,
        cmd: ClientRegisterCommand,
        operation_complete: Box<
            dyn FnOnce(OperationComplete) -> Pin<Box<dyn Future<Output = ()> + Send>> + Send + Sync,
        >,
    ) {
        let sector_id = cmd.header.sector_idx;

        self.restore_rid(sector_id).await;
        self.restore_ts_wr_val(sector_id).await;

        self.data.callback = Some(operation_complete);
        self.data.readlist.clear();
        self.data.acklist.clear();

        self.data.request_identifier = Some(cmd.header.request_identifier);

        match cmd.content {
            ClientRegisterCommandContent::Read => {
                self.data.reading = true;
            }
            ClientRegisterCommandContent::Write { data } => {
                self.data.writing = true;
                self.data.writeval = Some(data);
            }
        }

        self.register_client
            .broadcast(Broadcast {
                cmd: Arc::new(SystemRegisterCommand {
                    header: SystemCommandHeader {
                        process_identifier: self.self_ident,
                        msg_ident: Uuid::new_v4(),
                        read_ident: self.data.rid,
                        sector_idx: cmd.header.sector_idx,
                    },
                    content: SystemRegisterCommandContent::ReadProc,
                }),
            })
            .await;
    }

    /// Send system command to the register.
    async fn system_command(&mut self, cmd: SystemRegisterCommand) {
        let sector_idx = cmd.header.sector_idx;
        let process_identifier = cmd.header.process_identifier;
        let header = SystemCommandHeader {
            process_identifier: self.self_ident,
            msg_ident: Uuid::new_v4(),
            read_ident: cmd.header.read_ident,
            sector_idx,
        };

        match cmd.content {
            // Those two commands are performed on a helper registers.
            SystemRegisterCommandContent::ReadProc => {
                let sector_data = self.sectors_manager.read_data(sector_idx).await;
                let (ts, wr) = self.sectors_manager.read_metadata(sector_idx).await;

                self.register_client
                    .send(crate::Send {
                        cmd: Arc::new(SystemRegisterCommand {
                            header,
                            content: SystemRegisterCommandContent::Value {
                                timestamp: ts,
                                write_rank: wr,
                                sector_data,
                            },
                        }),
                        target: process_identifier as usize,
                    })
                    .await
            }
            SystemRegisterCommandContent::WriteProc {
                timestamp,
                write_rank,
                data_to_write,
            } => {
                let ts_wr = self.sectors_manager.read_metadata(sector_idx).await;
                if (timestamp, write_rank) > ts_wr {
                    self.sectors_manager
                        .write(sector_idx, &(data_to_write, timestamp, write_rank))
                        .await;
                }

                self.register_client
                    .send(crate::Send {
                        cmd: Arc::new(SystemRegisterCommand {
                            header,
                            content: SystemRegisterCommandContent::Ack,
                        }),
                        target: process_identifier as usize,
                    })
                    .await
            }

            // Those two commands are performed on the main register.
            SystemRegisterCommandContent::Ack => {
                if self.data.rid != cmd.header.read_ident || !self.data.write_phase {
                    return;
                }

                self.data.acklist.insert(process_identifier);

                if self.data.acklist.len() > self.processes_count / 2
                    && (self.data.reading || self.data.writing)
                {
                    self.data.acklist.clear();
                    self.data.write_phase = false;

                    let op_return: OperationReturn;
                    if self.data.reading {
                        self.data.reading = false;
                        op_return = OperationReturn::Read(ReadReturn {
                            read_data: std::mem::replace(&mut self.data.readval, None),
                        })
                    } else {
                        self.data.writing = false;
                        op_return = OperationReturn::Write;
                    }

                    let callback = std::mem::replace(&mut self.data.callback, None).unwrap();

                    callback(OperationComplete {
                        status_code: StatusCode::Ok,
                        request_identifier: self.data.request_identifier.unwrap(),
                        op_return,
                    })
                    .await;
                }
            }
            SystemRegisterCommandContent::Value {
                timestamp,
                write_rank,
                sector_data,
            } => {
                if self.data.rid != cmd.header.read_ident || self.data.write_phase {
                    return;
                }

                self.data.readlist.insert(process_identifier);

                if let Some((ts_wr, _data)) = &self.data.highest {
                    let new_ts_wr = (timestamp, write_rank);
                    if new_ts_wr > *ts_wr {
                        self.data.highest = Some(((new_ts_wr), sector_data));
                    }
                }

                if self.data.readlist.len() > self.processes_count / 2
                    && (self.data.reading || self.data.writing)
                {
                    let ((max_ts, rr), data_to_write) =
                        std::mem::replace(&mut self.data.highest, None).unwrap();

                    self.data.readval = Some(data_to_write.clone());

                    self.data.acklist.clear();
                    self.data.readlist.clear();
                    self.data.write_phase = true;

                    let content = if self.data.reading {
                        SystemRegisterCommandContent::WriteProc {
                            timestamp: max_ts,
                            write_rank: rr,
                            data_to_write,
                        }
                    } else {
                        SystemRegisterCommandContent::WriteProc {
                            timestamp: max_ts + 1,
                            write_rank: self.self_ident,
                            data_to_write: std::mem::replace(&mut self.data.writeval, None)
                                .unwrap(),
                        }
                    };

                    self.register_client
                        .broadcast(Broadcast {
                            cmd: Arc::new(SystemRegisterCommand { header, content }),
                        })
                        .await;
                }
            }
        }
    }
}
