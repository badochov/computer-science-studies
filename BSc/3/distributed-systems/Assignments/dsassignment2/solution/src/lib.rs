pub use atomic_register_public::*;
use helpers::get_socket_addr;
pub use register_client_public::*;
pub use sectors_manager_public::*;
pub use stable_storage_public::*;
use std::io::Error;
use std::sync::Arc;
use tokio::io::{AsyncWriteExt, BufReader, BufWriter};
use tokio::net::tcp::{OwnedReadHalf, OwnedWriteHalf};
use tokio::net::{TcpListener, TcpStream};
use tokio::sync::mpsc::{channel, Receiver, Sender};
use tokio::task::JoinHandle;
pub use transfer_public::*;
use transfer_responses::serialize_client_response;

pub use crate::domain::*;

mod atomic_register_public;
mod domain;
mod helpers;
mod register_client_public;
mod sectors_manager_public;
mod stable_storage_public;
mod transfer;
mod transfer_public;
mod transfer_responses;

pub async fn run_register_process(config: Configuration) {
    if !storage_dir_exists(&config) {
        panic!("Provided path doesn't exist.");
    }

    let tcp_listener = open_tcp_port(&config).await;

    let mut dispatcher = Dispatcher::new(config);

    while let Ok((stream, _)) = tcp_listener.accept().await {
        log::debug!("Connection {:?}", &stream);
        dispatcher.handle_connection(stream);
    }

    dispatcher.shutdown().await;
}

fn storage_dir_exists(config: &Configuration) -> bool {
    config.public.storage_dir.exists()
}

const NUM_ATOMIC_REGISTERS: usize = 256;

struct Handler {
    register: Box<dyn AtomicRegister>,
    client_rx: Receiver<(ClientRegisterCommand, Sender<WriteMessage>)>,
    system_rx: Receiver<SystemRegisterCommand>,
    cancellable_register_client: Arc<dyn CancellableRegisterClient>,
    max_sector: SectorIdx,
    op_tx: Sender<()>,
    op_rx: Receiver<()>,
    handled_sector: Option<SectorIdx>,
}

impl Handler {
    pub fn new(
        register: Box<dyn AtomicRegister>,
        client_rx: Receiver<(ClientRegisterCommand, Sender<WriteMessage>)>,
        system_rx: Receiver<SystemRegisterCommand>,
        max_sector: SectorIdx,
        cancellable_register_client: Arc<dyn CancellableRegisterClient>,
    ) -> Self {
        let (op_tx, op_rx) = channel(1);

        Self {
            register,
            client_rx,
            system_rx,
            max_sector,
            cancellable_register_client,
            handled_sector: None,
            op_tx,
            op_rx,
        }
    }

    pub fn handle_register(
        register: Box<dyn AtomicRegister>,
        client_rx: Receiver<(ClientRegisterCommand, Sender<WriteMessage>)>,
        system_rx: Receiver<SystemRegisterCommand>,
        max_sector: SectorIdx,
        cancellable_register_client: Arc<dyn CancellableRegisterClient>,
    ) -> JoinHandle<()> {
        tokio::task::spawn(async move {
            let handler = Self::new(
                register,
                client_rx,
                system_rx,
                max_sector,
                cancellable_register_client,
            );
            handler.handle().await;
        })
    }

    async fn handle(mut self) {
        loop {
            let v = tokio::select! {
                msg = self.client_rx.recv(), if self.handled_sector.is_none() => self.handle_client_message(msg).await,
                msg = self.system_rx.recv() => self.handle_system_message(msg).await,
                _ = self.op_rx.recv(), if self.handled_sector.is_some() => {
                    let handled_sector= std::mem::replace(&mut self.handled_sector, None);

                    self.cancellable_register_client.cancel(handled_sector.unwrap()).await;

                    Ok(())
                }
            };

            if let Err(e) = v {
                log::debug!("Err in Handler.handle: {:?}", e);
            }
        }
    }

    async fn handle_system_message(
        &mut self,
        system_msg: Option<SystemRegisterCommand>,
    ) -> Result<(), ()> {
        match system_msg {
            None => Err(()),
            Some(msg) => {
                self.register.system_command(msg).await;

                Ok(())
            }
        }
    }

    async fn handle_client_message(
        &mut self,
        client_msg: Option<(ClientRegisterCommand, Sender<WriteMessage>)>,
    ) -> Result<(), ()> {
        match client_msg {
            None => Err(()),
            Some((msg, rx)) => {
                if msg.header.sector_idx >= self.max_sector {
                    return rx
                        .send(WriteMessage::InvalidSectorIndex(msg))
                        .await
                        .map_err(|_| ());
                }
                let op_tx = self.op_tx.clone();

                self.handled_sector = Some(msg.header.sector_idx);

                self.register
                    .client_command(
                        msg,
                        Box::new(move |op| {
                            Box::pin(async move {
                                if let Err(e) = rx.send(WriteMessage::Ok(op)).await {
                                    log::debug!("While finalizint client message {:?}", e)
                                }
                                op_tx.send(()).await.unwrap();
                            })
                        }),
                    )
                    .await;

                Ok(())
            }
        }
    }
}

struct Dispatcher {
    tasks: Vec<JoinHandle<()>>,
    server_hmac: [u8; 64],
    client_hmac: [u8; 32],
    main_loop: JoinHandle<()>,

    client_txs: Vec<Sender<(ClientRegisterCommand, Sender<WriteMessage>)>>,
    system_txs: Vec<Sender<SystemRegisterCommand>>,
}

impl Dispatcher {
    pub fn new(config: Configuration) -> Self {
        let server_hmac = config.hmac_system_key;
        let client_hmac = config.hmac_client_key;

        let (client_txs, client_rxs) = Dispatcher::create_channels();
        let (system_txs, system_rxs) = Dispatcher::create_channels();

        let main_loop = Dispatcher::get_main_loop_handle(config, client_rxs, system_rxs);

        Self {
            tasks: Vec::new(),
            server_hmac,
            client_hmac,
            main_loop,
            client_txs,
            system_txs,
        }
    }

    pub fn handle_connection(&mut self, stream: TcpStream) {
        let task = self.spawn_new_task(stream);
        self.tasks.push(task);
    }

    pub async fn shutdown(self) {
        std::mem::drop(self.client_txs);
        std::mem::drop(self.system_txs);

        for task in self.tasks {
            if let Err(e) = task.await {
                log::debug!("While shutting down system {:?}", e);
            }
        }

        if let Err(e) = self.main_loop.await {
            log::debug!("While shutting down system {:?}", e);
        }
    }

    fn create_channels<T>() -> (Vec<Sender<T>>, Vec<Receiver<T>>) {
        let num_registers = NUM_ATOMIC_REGISTERS;

        let mut txs = Vec::new();
        let mut rxs = Vec::new();

        for _ in 0..num_registers {
            let (tx, rx) = channel(16);

            txs.push(tx);
            rxs.push(rx);
        }

        (txs, rxs)
    }

    fn get_main_loop_handle(
        config: Configuration,
        client_rxs: Vec<Receiver<(ClientRegisterCommand, Sender<WriteMessage>)>>,
        system_rxs: Vec<Receiver<SystemRegisterCommand>>,
    ) -> JoinHandle<()> {
        tokio::task::spawn(async move {
            let handlers = Self::spawn_handlers(config, client_rxs, system_rxs).await;

            for handler in handlers {
                handler.await.unwrap();
            }
        })
    }

    async fn spawn_handlers(
        config: Configuration,
        client_rxs: Vec<Receiver<(ClientRegisterCommand, Sender<WriteMessage>)>>,
        system_rxs: Vec<Receiver<SystemRegisterCommand>>,
    ) -> Vec<JoinHandle<()>> {
        let max_sector = config.public.max_sector;

        let cancellable_register_client = get_cancellable_register_client(&config);

        let registers =
            Self::spawn_atomic_registers(config, cancellable_register_client.clone().to_base())
                .await;

        client_rxs
            .into_iter()
            .zip(system_rxs)
            .zip(registers)
            .map(|((client_rx, system_rx), register)| {
                Handler::handle_register(
                    register,
                    client_rx,
                    system_rx,
                    max_sector,
                    cancellable_register_client.clone(),
                )
            })
            .collect()
    }

    async fn spawn_atomic_registers(
        config: Configuration,
        register_client: Arc<dyn RegisterClient>,
    ) -> Vec<Box<dyn AtomicRegister>> {
        tokio::task::spawn(async move {
            let process_count = config.public.tcp_locations.len();
            let num_atomic_registers = NUM_ATOMIC_REGISTERS;

            let sectors_manager_dir = config.public.storage_dir.join("sectors_manager");
            tokio::fs::create_dir_all(&sectors_manager_dir)
                .await
                .unwrap();
            let sectors_manager = build_sectors_manager(sectors_manager_dir);

            let base_storage_dir = config.public.storage_dir.join("stable_storage");
            tokio::fs::create_dir_all(&base_storage_dir).await.unwrap();

            let tasks = (1..num_atomic_registers).map(|i| {
                let storage_dir = base_storage_dir.join(i.to_string());
                let rank = config.public.self_rank;
                let register_client_clone = register_client.clone();
                let sectors_manager_clone = sectors_manager.clone();
                tokio::task::spawn(async move {
                    tokio::fs::create_dir_all(&storage_dir).await.unwrap();
                    let metadata = build_stable_storage(storage_dir).await;

                    build_atomic_register(
                        rank,
                        metadata,
                        register_client_clone,
                        sectors_manager_clone,
                        process_count,
                    )
                    .await
                })
            });

            let mut res = Vec::new();
            res.reserve(tasks.len());
            for task in tasks {
                res.push(task.await.unwrap());
            }

            res
        })
        .await
        .unwrap()
    }

    fn spawn_new_task(&mut self, stream: TcpStream) -> JoinHandle<()> {
        let (read_stream, write_stream) = stream.into_split();
        let (write_tx, write_rx) = channel(128);

        let read_handler = ReadHandler {
            write_tx,
            data: BufReader::new(read_stream),
            client_txs: self.client_txs.clone(),
            system_txs: self.system_txs.clone(),
            hmac_client: self.client_hmac,
            hmac_server: self.server_hmac,
        };

        let write_handler = WriteHandler {
            writer: BufWriter::new(write_stream),
            rx: write_rx,
            hmac_key: self.client_hmac,
        };

        tokio::task::spawn(async move {
            tokio::join!(read_handler.handle(), write_handler.handle());
        })
    }
}

#[derive(Debug)]
enum WriteMessage {
    Ok(OperationComplete),
    InvalidHmac(ClientRegisterCommand),
    InvalidSectorIndex(ClientRegisterCommand),
}

struct WriteHandler {
    rx: Receiver<WriteMessage>,
    writer: BufWriter<OwnedWriteHalf>,
    hmac_key: [u8; 32],
}

impl WriteHandler {
    async fn handle(mut self) {
        while let Some(msg) = self.rx.recv().await {
            let res = match msg {
                WriteMessage::Ok(cmd) => self.send_response(cmd).await,
                WriteMessage::InvalidHmac(cmd) => self.send_invalid_hmac_response(cmd).await,
                WriteMessage::InvalidSectorIndex(cmd) => {
                    self.send_invalid_sector_response(cmd).await
                }
            };

            if let Err(e) = res {
                log::debug!("Err in WriteHandler.handle: {:?}", e);

                return;
            }
        }
    }

    async fn send_response(&mut self, msg: OperationComplete) -> Result<(), Error> {
        serialize_client_response(msg, &mut self.writer, &self.hmac_key).await?;

        self.writer.flush().await
    }

    async fn send_invalid_sector_response(
        &mut self,
        msg: ClientRegisterCommand,
    ) -> Result<(), Error> {
        self.send_error_response(msg, StatusCode::InvalidSectorIndex)
            .await
    }

    async fn send_invalid_hmac_response(
        &mut self,
        msg: ClientRegisterCommand,
    ) -> Result<(), Error> {
        self.send_error_response(msg, StatusCode::AuthFailure).await
    }

    async fn send_error_response(
        &mut self,
        msg: ClientRegisterCommand,
        status_code: StatusCode,
    ) -> Result<(), Error> {
        serialize_client_response(
            OperationComplete {
                status_code,
                request_identifier: msg.header.request_identifier,
                op_return: match msg.content {
                    ClientRegisterCommandContent::Read => {
                        OperationReturn::Read(ReadReturn { read_data: None })
                    }
                    ClientRegisterCommandContent::Write { data: _ } => OperationReturn::Write,
                },
            },
            &mut self.writer,
            &self.hmac_key,
        )
        .await?;

        self.writer.flush().await
    }
}

struct ReadHandler {
    write_tx: Sender<WriteMessage>,
    data: BufReader<OwnedReadHalf>,
    client_txs: Vec<Sender<(ClientRegisterCommand, Sender<WriteMessage>)>>,
    system_txs: Vec<Sender<SystemRegisterCommand>>,
    hmac_server: [u8; 64],
    hmac_client: [u8; 32],
}

impl ReadHandler {
    pub async fn handle(mut self) {
        loop {
            match deserialize_register_command(&mut self.data, &self.hmac_server, &self.hmac_client)
                .await
            {
                Err(e) => {
                    log::debug!("Err in ReadHandler.handle: {:?}", e);

                    return;
                }
                Ok((msg, hmac_ok)) => {
                    if let Err(_) = self.handle_message(msg, hmac_ok).await {
                        return;
                    }
                }
            }
        }
    }

    async fn handle_message(&self, msg: RegisterCommand, hmac_ok: bool) -> Result<(), ()> {
        if !hmac_ok {
            match msg {
                RegisterCommand::Client(msg) => {
                    self.write_tx
                        .send(WriteMessage::InvalidHmac(msg))
                        .await
                        .map_err(|_| ())?;
                }
                _ => {}
            };

            Ok(())
        } else {
            match msg {
                RegisterCommand::System(msg) => self.handle_system_message(msg).await,
                RegisterCommand::Client(msg) => self.handle_client_message(msg).await,
            }
        }
    }

    async fn handle_system_message(&self, msg: SystemRegisterCommand) -> Result<(), ()> {
        self.system_txs[self.get_idx(msg.header.sector_idx)]
            .send(msg)
            .await
            .map_err(|_| ())
    }

    async fn handle_client_message(&self, msg: ClientRegisterCommand) -> Result<(), ()> {
        self.client_txs[self.get_idx(msg.header.sector_idx)]
            .send((msg, self.write_tx.clone()))
            .await
            .map_err(|_| ())
    }
    fn get_idx(&self, sector_idx: SectorIdx) -> usize {
        (sector_idx % self.client_txs.len() as u64) as usize
    }
}

async fn open_tcp_port(config: &Configuration) -> TcpListener {
    let location = &config.public.tcp_locations[(config.public.self_rank - 1) as usize];

    TcpListener::bind(get_socket_addr(location).await)
        .await
        .unwrap()
}
