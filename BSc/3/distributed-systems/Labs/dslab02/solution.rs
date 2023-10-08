use core::sync::atomic::{AtomicBool,Ordering};
use std::sync::{Arc, Condvar, Mutex};
use std::thread::JoinHandle;

type Task = Box<dyn FnOnce() + Send>;
type SynchronizedTaskQueue = Mutex<std::collections::VecDeque<Task>>;

// You can define new types (e.g., structs) if you need.
// However, they shall not be public (i.e., do not use the `pub` keyword).

/// The thread pool.
pub struct Threadpool { 
    threads: Vec<JoinHandle<()>>,
    task_queue: Arc<SynchronizedTaskQueue>,
    cond_var: Arc<Condvar>,
    should_finish: Arc<AtomicBool>

    // Add here any fields you need.
    // We suggest storing handles of the worker threads, submitted tasks,
    // and an information whether the pool is running or it is to be finished.
}

impl Threadpool {
    /// Create new thread pool with `workers_count` workers.
    pub fn new(workers_count: usize) -> Self {
        let mut threads = vec![];
        let cond_var = Arc::new(Condvar::new());
        let task_queue = Arc::new(Mutex::new(std::collections::VecDeque::new()));
        let should_finish = Arc::new(AtomicBool::new(false));

        for _ in 0..workers_count {
            let cond_var_clone = cond_var.clone(); 
            let task_queue_clone = task_queue.clone(); 
            let should_finish_clone = should_finish.clone(); 
            threads.push(std::thread::spawn(move || Threadpool::worker_loop(cond_var_clone, task_queue_clone, should_finish_clone)))
        }

        return Threadpool {
            threads,
            task_queue,
            cond_var,
            should_finish,
        };
    }

    /// Submit a new task.
    pub fn submit(&self, task: Task) {
        let mut tasks = self.task_queue.lock().unwrap();
        tasks.push_back(task);
        self.cond_var.notify_one();
    }

    // We suggest extracting the implementation of the worker to an associated
    // function, like this one (however, it is not a part of the public
    // interface, so you can delete it if you implement it differently):
    fn worker_loop(new_task_cond: Arc<Condvar>, task_queue: Arc<SynchronizedTaskQueue>, should_finish: Arc<AtomicBool>) {
        loop {
            let task = Threadpool::get_task(new_task_cond.clone(), task_queue.clone(), should_finish.clone());
            match task {
                None => return,
                Some(t) => t(),
            }
        }
    }

    fn get_task(new_task_cond: Arc<Condvar>, task_queue: Arc<SynchronizedTaskQueue>, should_finish: Arc<AtomicBool>) -> Option<Task> {
        let mut tasks = task_queue.lock().unwrap();
        loop {
            let task = tasks.pop_front();
            if task.is_some() {
                return task;
            }
            else if should_finish.load(Ordering::Relaxed){
                return Option::None;
            }
            tasks = new_task_cond.wait(tasks).unwrap();
        }
    }
}

fn notify_threads(pool: &mut Threadpool) {
    pool.should_finish.store(true, Ordering::Relaxed);
    let _guard = pool.task_queue.lock().unwrap();
    pool.cond_var.notify_all();
}

impl Drop for Threadpool {
    /// Gracefully end the thread pool.
    ///
    /// It waits until all submitted tasks are executed,
    /// and until all threads are joined.
    fn drop(&mut self) {
        notify_threads(self);
        
        for thread in self.threads.drain(..) {
            thread.join().unwrap();
        }
    }
}
