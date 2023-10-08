use std::collections::{HashMap, HashSet};
use std::iter::FromIterator;
use std::net::SocketAddr;
use std::sync::Arc;
use std::time::Duration;

use bytes::BytesMut;
use log::debug;
use serde::{Deserialize, Serialize};
use tokio::net::UdpSocket;
use uuid::Uuid;

use executor::{Handler, ModuleRef, System, Tick};

/// A message which disables a process. Used for testing
pub struct Disable;

pub struct FailureDetectorModule {
    enabled: bool,
    addresses: HashMap<Uuid, SocketAddr>,
    all_idents: HashSet<Uuid>,
    alive: HashSet<Uuid>,
    suspected: HashSet<Uuid>,
    socket: Arc<UdpSocket>,
    were_alive: HashSet<Uuid>,
    ident: Uuid,
    // TODO add whatever fields necessary.
}

impl FailureDetectorModule {
    pub async fn new(
        system: &mut System,
        delay: Duration,
        addresses: &HashMap<Uuid, SocketAddr>,
        ident: Uuid,
        all_idents: HashSet<Uuid>,
    ) -> ModuleRef<Self> {
        let addr = addresses.get(&ident).unwrap();
        let socket = Arc::new(UdpSocket::bind(addr).await.unwrap());

        let module_ref = system
            .register_module(Self {
                enabled: true,
                ident,
                addresses: addresses.clone(),
                alive: all_idents.clone(),
                were_alive: all_idents.clone(),
                all_idents,
                suspected: HashSet::new(),
                socket: socket.clone(),
            })
            .await;

        tokio::spawn(deserialize_and_forward(socket, module_ref.clone()));

        system.request_tick(&module_ref, delay).await;

        module_ref
    }

    async fn send_detector_operation(&self, op: DetectorOperation, addr: SocketAddr) {
        let msg = bincode::serialize(&op).unwrap();
        self.socket.send_to(msg.as_slice(), addr).await.unwrap();
    }
}

/// New operation arrived on socket.
#[async_trait::async_trait]
impl Handler<DetectorOperationUdp> for FailureDetectorModule {
    async fn handle(&mut self, item: DetectorOperationUdp) {
        if self.enabled {
            // Perform health check and query status.
            let DetectorOperationUdp(op, addr) = item;
            match op {
                DetectorOperation::HeartbeatRequest => {
                    self.send_detector_operation(
                        DetectorOperation::HeartbeatResponse(self.ident),
                        addr,
                    )
                    .await
                }
                DetectorOperation::HeartbeatResponse(uuid) => {
                    self.alive.insert(uuid);
                }
                DetectorOperation::AliveRequest => {
                    self.send_detector_operation(
                        DetectorOperation::AliveInfo(Vec::from_iter(self.were_alive.clone())),
                        addr,
                    )
                    .await
                }
                DetectorOperation::AliveInfo(_) => {}
            }
        }
    }
}

/// Called periodically to check send broadcast and update alive processes.
#[async_trait::async_trait]
impl Handler<Tick> for FailureDetectorModule {
    async fn handle(&mut self, _msg: Tick) {
        self.were_alive = self.alive.clone();
        if self.enabled {
            for p in self.all_idents.iter() {
                let alive = self.alive.contains(p);
                let suspected = self.suspected.contains(p);
                if !alive && !suspected {
                    self.suspected.insert(*p);
                } else if alive && suspected {
                    self.suspected.remove(p);
                }

                self.send_detector_operation(
                    DetectorOperation::HeartbeatRequest,
                    *self.addresses.get(p).unwrap(),
                )
                .await
            }
        }
        self.alive.clear();
    }
}

#[async_trait::async_trait]
impl Handler<Disable> for FailureDetectorModule {
    async fn handle(&mut self, _msg: Disable) {
        self.enabled = false;
    }
}

async fn deserialize_and_forward(
    socket: Arc<UdpSocket>,
    module_ref: ModuleRef<FailureDetectorModule>,
) {
    let mut buffer = BytesMut::new();
    while socket.readable().await.is_ok() {
        if let Ok((_, sender)) = socket.try_recv_buf_from(&mut buffer) {
            match bincode::deserialize(&buffer.split()) {
                Ok(msg) => module_ref.send(DetectorOperationUdp(msg, sender)).await,
                Err(err) => {
                    debug!("Invalid format of detector operation ({})!", err);
                }
            }
        }
    }
}

struct DetectorOperationUdp(DetectorOperation, SocketAddr);

#[derive(Serialize, Deserialize)]
pub enum DetectorOperation {
    /// Request to receive a heartbeat.
    HeartbeatRequest,
    /// Response to heartbeat, contains uuid of the receiver of HeartbeatRequest.
    HeartbeatResponse(Uuid),
    /// Request to receive information about working processes.
    AliveRequest,
    /// Vector of processes which are alive according to AliveRequest receiver.
    AliveInfo(Vec<Uuid>),
}
