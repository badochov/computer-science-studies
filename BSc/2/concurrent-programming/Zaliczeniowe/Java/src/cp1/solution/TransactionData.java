package cp1.solution;

import cp1.base.Resource;
import cp1.base.ResourceId;
import cp1.base.ResourceOperation;
import cp1.base.ResourceOperationException;

import java.util.Collection;
import java.util.HashSet;
import java.util.Stack;

public class TransactionData {
    private final long startTime;
    private final Collection<ResourceId> resourceIds;
    private final Collection<OperationData> operations;
    private boolean aborted;

    TransactionData(long time) {
        startTime = time;
        resourceIds = new HashSet<>();
        aborted = false;
        operations = new Stack<>();
    }

    public long getStartTime() {
        return startTime;
    }

    public void addResourceId(ResourceId resourceId) {
        resourceIds.add(resourceId);
    }

    public Collection<ResourceId> getResourceIds() {
        return resourceIds;
    }

    public void abort() {
        aborted = true;
    }

    public void execute(ResourceOperation operation, Resource resource) throws ResourceOperationException, InterruptedException {
        operation.execute(resource);
        if (Thread.interrupted()) {
            operation.undo(resource);
            throw new InterruptedException();
        }
        operations.add(new OperationData(operation, resource));
    }

    public void undo() {
        operations.forEach(this::undoOperation);
    }

    private void undoOperation(OperationData operationData) {
        operationData.getOperation().undo(operationData.getResource());
    }

    public boolean isAborted() {
        return aborted;
    }


}
