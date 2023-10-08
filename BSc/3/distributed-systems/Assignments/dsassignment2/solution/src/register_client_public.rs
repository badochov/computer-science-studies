use crate::helpers::get_socket_addr;
use crate::Configuration;
use crate::RegisterCommand;
use crate::{serialize_register_command, SectorIdx};
use std::collections::HashMap;
use std::time::Duration;
use tokio::sync::mpsc::{unbounded_channel, UnboundedReceiver, UnboundedSender};
use tokio::sync::RwLock;

use std::sync::Arc;
use tokio::net::TcpStream;
use tokio::task::JoinHandle;

use crate::SystemRegisterCommand;

#[async_trait::async_trait]
/// We do not need any public implementation of this trait. It is there for use
/// in AtomicRegister. In our opinion it is a safe bet to say some structure of
/// this kind must appear in your solution.
pub trait RegisterClient: core::marker::Send + core::marker::Sync {
    /// Sends a system message to a single process.
    async fn send(&self, msg: Send);

    /// Broadcasts a system message to all processes in the system, including self.
    async fn broadcast(&self, msg: Broadcast);
}

#[async_trait::async_trait]
pub trait CancellableRegisterClient: RegisterClient {
    async fn cancel(&self, sector_idx: SectorIdx);

    fn to_base<'a>(self: Arc<Self>) -> Arc<dyn RegisterClient + 'a>;
}

pub struct Broadcast {
    pub cmd: Arc<SystemRegisterCommand>,
}

pub struct Send {
    pub cmd: Arc<SystemRegisterCommand>,
    /// Identifier of the target process. Those start at 1.
    pub target: usize,
}

struct TargetData {
    sender: UnboundedSender<Arc<SystemRegisterCommand>>,
}

impl TargetData {
    pub fn new(location: (String, u16), configuration: &Configuration) -> Self {
        let (command_tx, command_rx) = unbounded_channel();
        let hmac = configuration.hmac_system_key.clone();

        tokio::task::spawn(async move {
            WorkerThread::new(get_socket_addr(&location).await, command_rx, hmac)
                .start()
                .await
        });

        Self { sender: command_tx }
    }
}

struct WorkerThread {
    addr: std::net::SocketAddr,
    rx: UnboundedReceiver<Arc<SystemRegisterCommand>>,
    hmac: [u8; 64],
}

impl WorkerThread {
    pub fn new(
        addr: std::net::SocketAddr,
        rx: UnboundedReceiver<Arc<SystemRegisterCommand>>,
        hmac: [u8; 64],
    ) -> Self {
        WorkerThread { addr, rx, hmac }
    }

    pub async fn start(mut self) {
        loop {
            if let Ok(mut socket) = TcpStream::connect(self.addr).await {
                while let Some(cmd) = self.rx.recv().await {
                    let cmd_helper = (*cmd).clone();

                    let reg_cmd = RegisterCommand::System(cmd_helper.clone());

                    match serialize_register_command(&reg_cmd, &mut socket, &self.hmac).await {
                        Ok(()) => {}
                        Err(e) => {
                            log::debug!("While serializing system command {:?}", e);
                            break;
                        }
                    }
                }
            }
        }
    }
}

struct LoggedRegisterClient {
    targets: Arc<Vec<TargetData>>,
    to_rebroadcast: Arc<RwLock<HashMap<SectorIdx, Vec<Arc<SystemRegisterCommand>>>>>,
}

impl LoggedRegisterClient {
    pub fn new_ref(configuration: &Configuration) -> Arc<Self> {
        let targets = configuration
            .public
            .tcp_locations
            .iter()
            .map(|location| TargetData::new(location.clone(), configuration))
            .collect();

        let to_rebroadcast = Arc::new(RwLock::new(HashMap::new()));

        let res = Arc::new(Self {
            targets: Arc::new(targets),
            to_rebroadcast,
        });

        LoggedRegisterClient::start_rebroadcast_task(res.clone());

        res
    }

    async fn resend(&self) {
        for (_sector_idx, cmds) in self.to_rebroadcast.read().await.iter() {
            for cmd in cmds {
                self.do_broadcast(&Broadcast { cmd: cmd.clone() }).await;
            }
        }
    }

    fn start_rebroadcast_task(self_ref: Arc<Self>) -> JoinHandle<()> {
        tokio::task::spawn(async move {
            let mut interval = tokio::time::interval(Duration::from_secs(2));

            loop {
                interval.tick().await;
                self_ref.resend().await;
            }
        })
    }

    async fn do_broadcast(&self, msg: &Broadcast) {
        let tasks = (1..self.targets.len() + 1).map(|target| {
            let clone = self.clone();
            let cmd = msg.cmd.clone();
            tokio::task::spawn(async move {
                clone.send(Send { cmd, target }).await;
            })
        });

        for task in tasks {
            task.await.unwrap();
        }
    }
}

#[async_trait::async_trait]
impl RegisterClient for LoggedRegisterClient {
    /// Sends a system message to a single process.
    async fn send(&self, msg: Send) {
        let target = &self.targets[msg.target - 1];

        target.sender.send(msg.cmd).unwrap()
    }

    /// Broadcasts a system message to all processes in the system, including self.
    async fn broadcast(&self, msg: Broadcast) {
        self.do_broadcast(&msg).await;

        self.to_rebroadcast
            .write()
            .await
            .entry(msg.cmd.header.sector_idx)
            .or_insert(Vec::new())
            .push(msg.cmd);
    }
}

#[async_trait::async_trait]
impl CancellableRegisterClient for LoggedRegisterClient {
    async fn cancel(&self, sector_idx: SectorIdx) {
        self.to_rebroadcast.write().await.remove(&sector_idx);
    }

    fn to_base<'a>(self: Arc<Self>) -> Arc<dyn RegisterClient + 'a> {
        self
    }
}

pub fn get_cancellable_register_client(
    configuration: &Configuration,
) -> Arc<dyn CancellableRegisterClient> {
    LoggedRegisterClient::new_ref(configuration)
}

impl Clone for LoggedRegisterClient {
    fn clone(&self) -> LoggedRegisterClient {
        LoggedRegisterClient {
            targets: self.targets.clone(),
            to_rebroadcast: self.to_rebroadcast.clone(),
        }
    }
}
