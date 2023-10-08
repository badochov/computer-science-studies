package cp1.solution;

import cp1.base.*;

import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;
import java.util.concurrent.Semaphore;

public class Menago implements TransactionManager {
    private final ConcurrentMap<ResourceId, ResourceStatus> resourcesStatus;
    private final Map<Thread, ResourceId> waitsFor;
    private final Map<ResourceId, Integer> waitingFor;
    private final Map<ResourceId, Resource> resourceMap;
    private final LocalTimeProvider timeProvider;
    private final ConcurrentMap<Thread, TransactionData> transactions;
    private final Semaphore mutex;
    private final Semaphore closing;
    private final Map<ResourceId, Semaphore> resourceSemaphore;

    public Menago(
            Collection<Resource> resources,
            LocalTimeProvider timeProvider
    ) {
        resourcesStatus = new ConcurrentHashMap<>();
        resourceMap = new HashMap<>();
        this.timeProvider = timeProvider;
        transactions = new ConcurrentHashMap<>();
        mutex = new Semaphore(1, true);
        closing = new Semaphore(0, true);
        resourceSemaphore = new HashMap<>();
        waitsFor = new HashMap<>();
        waitingFor = new HashMap<>();

        addResources(resources);
    }

    private void addResources(Collection<Resource> resources) {
        for (Resource resource : resources) {
            resourceSemaphore.put(resource.getId(), new Semaphore(1, true));
            resourcesStatus.put(resource.getId(), new ResourceStatus());
            resourceMap.put(resource.getId(), resource);
            waitingFor.put(resource.getId(), 0);
        }
    }

    @Override
    public void startTransaction() throws AnotherTransactionActiveException {
        if (isTransactionActive()) {
            throw new AnotherTransactionActiveException();
        }
        long time = timeProvider.getTime();
        transactions.put(Thread.currentThread(), new TransactionData(time));
    }

    @Override
    public void operateOnResourceInCurrentTransaction(ResourceId rid, ResourceOperation operation) throws NoActiveTransactionException, UnknownResourceIdException, ActiveTransactionAborted, ResourceOperationException, InterruptedException {
        checkIfActionOnResourceBePerformed();

        Resource resource = getResourceById(rid);
        if (resource == null) {
            throw new UnknownResourceIdException(rid);
        }
        try {
            claimResource(resource);
            if (Thread.interrupted()) {
                throw new InterruptedException();
            }

            getThreadTransaction().execute(operation, resource);
        } catch (InterruptedException ex) {
            if (getThreadTransaction().isAborted()) {
                throw new ActiveTransactionAborted();
            }
            throw ex;
        }
    }

    private ResourceStatus getResourceStatus(Resource resource) {
        return resourcesStatus.get(resource.getId());
    }

    private Thread getBlockingThread(Thread thread) {
        ResourceId resourceId = waitsFor.get(thread);
        if (resourceId == null) {
            return null;
        }
        return resourcesStatus.get(resourceId).getOwner();
    }

    private Thread getThreadToInterrupt(Resource resource) {
        Thread latest = Thread.currentThread();
        TransactionData latestData = transactions.get(latest);
        Thread blocking = resourcesStatus.get(resource.getId()).getOwner();
        while (blocking != null && blocking != Thread.currentThread()) {
            TransactionData blockingData = transactions.get(blocking);

            if (isNewer(latest, latestData, blocking, blockingData)) {
                latest = blocking;
                latestData = blockingData;
            }

            blocking = getBlockingThread(blocking);
        }
        if (blocking == Thread.currentThread()) {
            return latest;
        }
        return null;
    }

    private boolean isNewer(Thread latest, TransactionData latestData, Thread blocking, TransactionData blockingData) {
        if (latestData.getStartTime() < blockingData.getStartTime()) {
            return true;
        }
        return latestData.getStartTime() == blockingData.getStartTime() && latest.getId() < blocking.getId();
    }

    private void handleDeadlock(Resource resource) {
        Thread thread = getThreadToInterrupt(resource);
        if (thread != null) {
            transactions.get(thread).abort();
            if (thread != Thread.currentThread()) {
                waitingFor.computeIfPresent(waitsFor.get(thread), (k, v) -> v - 1);
            }
            waitsFor.remove(thread);
            thread.interrupt();
        }
    }

    private void claimResource(Resource resource) throws InterruptedException {
        if (getResourceStatus(resource).isOwner()) {
            return;
        }

        mutex.acquire();
        boolean hasMutex = true;
        if (!getResourceStatus(resource).isAvailable()) {
            handleDeadlock(resource);
            if (Thread.interrupted()) {
                mutex.release();
                throw new InterruptedException();
            }
            waitsFor.put(Thread.currentThread(), resource.getId());
            waitingFor.computeIfPresent(resource.getId(), (k, v) -> v + 1);
            mutex.release();
            hasMutex = false;
        }
        InterruptedException ex = null;
        try {
            resourceSemaphore.get(resource.getId()).acquire();
        } catch (InterruptedException e) {
            if (!hasMutex) {
                mutex.acquireUninterruptibly();
            }

            ex = e;
        }
        if (!hasMutex && !getThreadTransaction().isAborted()) {
            waitsFor.remove(Thread.currentThread());
            waitingFor.computeIfPresent(resource.getId(), (k, v) -> v - 1);
        }

        if (ex != null) {
            mutex.release();
            throw ex;
        }


        TransactionData transactionData = getThreadTransaction();
        transactionData.addResourceId(resource.getId());
        resourcesStatus.get(resource.getId()).setOwner();

        if (hasMutex) {
            mutex.release();
        } else {
            closing.release();
        }


    }

    private void checkIfActionOnResourceBePerformed() throws NoActiveTransactionException, ActiveTransactionAborted {
        if (!isTransactionActive()) {
            throw new NoActiveTransactionException();
        }
        if (isTransactionAborted()) {
            throw new ActiveTransactionAborted();
        }
    }

    private Resource getResourceById(ResourceId rid) {
        return resourceMap.get(rid);
    }

    private TransactionData getThreadTransaction() {
        return transactions.get(Thread.currentThread());
    }

    @Override
    public void commitCurrentTransaction() throws NoActiveTransactionException, ActiveTransactionAborted {
        checkIfActionOnResourceBePerformed();

        closeTransaction();
    }

    @Override
    public void rollbackCurrentTransaction() {
        TransactionData transactionData = getThreadTransaction();
        if (transactionData == null) {
            return;
        }
        transactionData.undo();

        closeTransaction();
    }

    private synchronized void closeTransaction() {
        mutex.acquireUninterruptibly();
        TransactionData transactionData = getThreadTransaction();

        freeResources(transactionData.getResourceIds());

        transactions.remove(Thread.currentThread());
        mutex.release();
    }

    private void freeResources(Collection<ResourceId> resourceIds) {
        for (ResourceId id : resourceIds) {
            resourcesStatus.get(id).free();
            boolean wait = waitingFor.get(id) > 0;
            resourceSemaphore.get(id).release();
            if (wait) {
                closing.acquireUninterruptibly();
            }
        }
        closing.drainPermits();
    }

    @Override
    public boolean isTransactionActive() {
        return transactions.containsKey(Thread.currentThread());
    }

    @Override
    public boolean isTransactionAborted() {
        TransactionData data = transactions.get(Thread.currentThread());
        if (data != null) {
            return data.isAborted();
        }
        return false;
    }
}
