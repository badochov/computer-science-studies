use std::{
    sync::{
        atomic::{AtomicBool, Ordering},
        Arc,
    },
    time::Duration,
};

use executor::*;

use crate::*;

pub(crate) struct Timer {
    pub(crate) interval: Duration,
    pub(crate) abort: Arc<AtomicBool>,
}

impl Timer {
    pub(crate) fn new(interval: Duration) -> Self {
        Self {
            interval,
            abort: Arc::new(AtomicBool::new(false)),
        }
    }

    pub(crate) fn stop(&mut self) {
        self.abort.store(true, Ordering::Relaxed);
    }

    pub(crate) fn reset<M>(&mut self, mod_ref: ModuleRef<Raft>, msg: M)
    where
        M: Message + Clone + Send + std::marker::Sync,
        Raft: Handler<M>,
    {
        self.stop();
        self.abort = Arc::new(AtomicBool::new(false));
        let interval = self.interval;

        tokio::spawn(Self::run_timer(mod_ref, interval, self.abort.clone(), msg));
    }

    async fn run_timer<M>(
        raft_ref: ModuleRef<Raft>,
        interval: Duration,
        abort: Arc<AtomicBool>,
        msg: M,
    ) where
        M: Message + Send + Clone,
        Raft: Handler<M>,
    {
        let mut interval = tokio::time::interval(interval);
        interval.tick().await;
        interval.tick().await;
        while !abort.load(Ordering::Relaxed) {
            raft_ref.send(msg.clone()).await;
            interval.tick().await;
        }
    }
}
