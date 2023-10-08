#[cfg(test)]
mod tests {
    use crate::solution::Threadpool;
    use crossbeam_channel::unbounded;
    use ntest::timeout;
    use std::sync::atomic::{AtomicBool, AtomicU64, Ordering};
    use std::sync::{Arc, Barrier};

    #[test]
    #[timeout(200)]
    fn smoke_test() {
        let (tx, rx) = unbounded();
        let pool = Threadpool::new(1);

        pool.submit(Box::new(move || {
            tx.send(14).unwrap();
        }));

        assert_eq!(14, rx.recv().unwrap());
    }

    #[test]
    #[timeout(200)]
    fn threadpool_is_sync() {
        let send_only_when_threadpool_is_sync = Arc::new(Threadpool::new(1));
        let (tx, rx) = unbounded();

        let _handle = std::thread::spawn(move || {
            tx.send(send_only_when_threadpool_is_sync).unwrap();
        });

        rx.recv().unwrap();
    }

    #[test]
    #[timeout(200)]
    fn sum_1000_numbers() {
        let sum: Arc<AtomicU64> = Arc::new(AtomicU64::new(0));

        let pool = Threadpool::new(4);
        for i in 0..1000 {
            let cloned_sum = sum.clone();
            pool.submit(Box::new(move || {
                cloned_sum.fetch_add(i, Ordering::Relaxed);
            }));
        }

        std::mem::drop(pool);

        assert_eq!(sum.load(Ordering::Relaxed), (1000 * 999) / 2);
    }

    /* It's possible to handle panics but not required
    #[test]
    #[timeout(200)]
    fn panics_are_caught() {
        let pool = Threadpool::new(4);

        // Panic on all threads
        for _ in 0..10 {
            pool.submit(Box::new(|| panic!("Im panicking")));
        }

        // Thread pool still works
        let value: Arc<AtomicBool> = Arc::new(AtomicBool::new(false));
        let cloned_value = value.clone();
        pool.submit(Box::new(move || {
            cloned_value.store(true, Ordering::Relaxed)
        }));

        std::mem::drop(pool);
        assert_eq!(value.load(Ordering::Relaxed), true);
    }
    */
    

    #[test]
    #[timeout(200)]
    fn threadpool_is_concurrent() {
        let barrier1 = Arc::new(Barrier::new(4));
        let barrier2 = Arc::new(Barrier::new(4));

        let pool = Threadpool::new(4);
        for _ in 0..4 {
            let cloned_barrier1 = barrier1.clone();
            let cloned_barrier2 = barrier2.clone();
            pool.submit(Box::new(move || {
                cloned_barrier1.wait();
                cloned_barrier2.wait();
            }));
        }
    }

    #[test]
    #[timeout(1000)]
    fn drop_waits() {
        let pool = Threadpool::new(4);

        let value: Arc<AtomicBool> = Arc::new(AtomicBool::new(false));
        let cloned_value = value.clone();
        pool.submit(Box::new(move || {
            std::thread::sleep(std::time::Duration::from_millis(800));
            cloned_value.store(true, Ordering::Relaxed);
        }));

        std::mem::drop(pool);

        assert_eq!(value.load(Ordering::Relaxed), true);
    }
}
