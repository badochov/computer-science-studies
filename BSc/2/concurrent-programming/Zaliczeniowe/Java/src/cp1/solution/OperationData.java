package cp1.solution;

import cp1.base.Resource;
import cp1.base.ResourceOperation;

public class OperationData {
    private final ResourceOperation operation;
    private final Resource resource;

    OperationData(ResourceOperation op, Resource resource){
        operation = op;
        this.resource = resource;
    }

    public ResourceOperation getOperation() {
        return operation;
    }

    public Resource getResource() {
        return resource;
    }
}
