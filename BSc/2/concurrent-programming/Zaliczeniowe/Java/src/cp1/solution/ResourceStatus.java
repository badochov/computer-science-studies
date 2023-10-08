package cp1.solution;


public class ResourceStatus {
    private Thread owner;

    ResourceStatus() {
        owner = null;
    }

    public Thread getOwner() {
        return owner;
    }

    public void setOwner() {
        owner = Thread.currentThread();
    }

    public void free() {
        owner = null;
    }

    public boolean isAvailable() {
        return owner == null;
    }

    public boolean isOwner() {
        return owner == Thread.currentThread();
    }
}
