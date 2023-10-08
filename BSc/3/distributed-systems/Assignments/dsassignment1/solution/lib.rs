use std::time::Duration;
use tokio::sync::Mutex;
use tokio::task::JoinHandle;

pub trait Message: Send + 'static {}
impl<T: Send + 'static> Message for T {}

/// A trait for modules capable of handling messages of type `M`.
#[async_trait::async_trait]
pub trait Handler<M: Message>
where
    M: Message,
{
    /// Handles the message.
    async fn handle(&mut self, msg: M);
}

/// The message sent as a result of calling `System::request_tick()`.
#[derive(Debug, Clone)]
pub struct Tick {}

struct Ticker {
    tx: async_channel::Sender<()>,
    rx: async_channel::Receiver<()>,

    ticker_threads: Vec<JoinHandle<()>>,
}

impl Ticker {
    fn new() -> Self {
        let (tx, rx) = async_channel::unbounded::<()>();

        return Ticker {
            tx,
            rx,
            ticker_threads: Vec::new(),
        };
    }

    fn request_tick<T: Handler<Tick> + Send>(&mut self, requester: ModuleRef<T>, delay: Duration) {
        let rx = self.rx.clone();
        self.ticker_threads.push(tokio::spawn(async move {
            let mut interval = tokio::time::interval(delay);
            while let Err(async_channel::TryRecvError::Empty) = rx.try_recv() {
                interval.tick().await;
                requester.send(Tick {}).await;
            }
        }));
    }

    async fn shutdown(&mut self) {
        self.tx.close();

        for thread in self.ticker_threads.drain(..) {
            thread.await.unwrap();
        }
    }
}

// You can add fields to this struct.
pub struct System {
    mutex: Mutex<()>,

    module_threads: Vec<JoinHandle<()>>,
    module_shutdown_handles: Vec<Box<dyn FnOnce()>>,

    ticker: Ticker,
}

impl System {
    /// Schedules a `Tick` message to be sent to the given module periodically
    /// with the given interval. The first tick is sent immediately.
    pub async fn request_tick<T: Handler<Tick> + Send>(
        &mut self,
        requester: &ModuleRef<T>,
        delay: Duration,
    ) {
        self.mutex.lock().await;

        self.ticker.request_tick(requester.clone(), delay);
    }

    /// Registers the module in the system.
    /// Returns a `ModuleRef`, which can be used then to send messages to the module.
    pub async fn register_module<T: Send + 'static>(&mut self, mut module: T) -> ModuleRef<T> {
        self.mutex.lock().await;

        let (tx, rx) = async_channel::unbounded();
        let (shutdown_tx, shutdown_rx) = async_channel::unbounded::<()>();
        let mod_ref = ModuleRef { tx: tx.clone() };
        self.module_shutdown_handles.push(Box::new(move || {
            shutdown_tx.close();
            tx.close();
        }));

        self.module_threads.push(tokio::spawn(async move {
            while let Err(async_channel::TryRecvError::Empty) = shutdown_rx.try_recv() {
                match rx.recv().await {
                    Ok(msg) => msg.get_handled(&mut module).await,
                    _ => break,
                }
            }
        }));

        return mod_ref;
    }

    /// Creates and starts a new instance of the system.
    pub async fn new() -> Self {
        return System {
            mutex: Mutex::new(()),

            module_threads: Vec::new(),
            module_shutdown_handles: Vec::new(),
            ticker: Ticker::new(),
        };
    }

    /// Gracefully shuts the system down.
    pub async fn shutdown(&mut self) {
        self.mutex.lock().await;

        self.ticker.shutdown().await;

        for shutdown_handle in self.module_shutdown_handles.drain(..) {
            shutdown_handle();
        }

        for thread in self.module_threads.drain(..) {
            thread.await.unwrap();
        }
    }
}

pub struct ModuleRef<T: Send + 'static> {
    pub(crate) tx: async_channel::Sender<Box<dyn Handlee<T>>>,
}

impl<T: Send> ModuleRef<T> {
    /// Sends the message to the module.
    pub async fn send<M: Message>(&self, msg: M)
    where
        T: Handler<M>,
    {
        self.tx.clone().send(Box::new(msg)).await.unwrap();
    }
}

impl<T: Send> Clone for ModuleRef<T> {
    /// Creates a new reference to the same module.
    fn clone(&self) -> Self {
        return ModuleRef {
            tx: self.tx.clone(),
        };
    }
}

#[async_trait::async_trait]
trait Handlee<T>: Send + 'static
where
    T: Send,
{
    async fn get_handled(self: Box<Self>, module: &mut T);
}

#[async_trait::async_trait]
impl<M, T> Handlee<T> for M
where
    T: Handler<M> + Send,
    M: Message,
{
    async fn get_handled(self: Box<Self>, module: &mut T) {
        module.handle(*self).await
    }
}
