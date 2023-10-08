package cw3.hanoi;

import java.util.Vector;

class Move<T, S> {
    private T from;
    private T to;
    private Disk<S> disk;

    public Move(T from, T to, Disk<S> d) {
        this.from = from;
        this.to = to;
        this.disk = d;
    }

    public void print() {
        System.out.printf("z: %s, na: %s, nazwa krążka: %s\n", this.from.toString(), this.to.toString(),
                this.disk.getName().toString());
    }

    public static void print(Vector<? extends Move<?, ?>> moves) {
        for (Move move : moves) {
            move.print();
        }
    }
}