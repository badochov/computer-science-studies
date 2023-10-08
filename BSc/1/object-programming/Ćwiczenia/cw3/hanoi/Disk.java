package cw3.hanoi;

class Disk<T> {
    private T name;

    Disk(T name) {
        this.name = name;
    }

    public T getName() {
        return this.name;
    }
}