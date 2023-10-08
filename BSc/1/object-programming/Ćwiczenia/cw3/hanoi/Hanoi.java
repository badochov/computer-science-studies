package cw3.hanoi;

import java.util.Vector;

class Hanoi<T, D> {
    private T[] rodNames;
    private Vector<Vector<Disk<D>>> disks;
    Vector<Move<T, D>> moves;

    public Hanoi(T[] rods, D[] diskNames) {

        this.rodNames = rods;
        this.disks = new Vector<>(rods.length);

        Vector<Disk<D>> dV = new Vector<>();
        for (int i = diskNames.length - 1; i >= 0; i--) {
            dV.add(new Disk<D>(diskNames[i]));
        }

        for (int i = 0; i < rods.length; i++) {
            this.disks.add(new Vector<>());
        }

        this.disks.set(0, dV);
        this.moves = new Vector<>();
    }

    public Vector<Move<T, D>> solve() {
        move(0, 2, 1, this.disks.get(0).size());
        return this.moves;
    }

    public void move(int from, int to) {
        Vector<Disk<D>> fV = this.disks.get(from);
        Vector<Disk<D>> tV = this.disks.get(to);
        Disk<D> d = fV.remove(fV.size() - 1);
        tV.add(d);
        this.moves.add(new Move<>(this.rodNames[from], this.rodNames[to], d));
    }

    public void move(int from, int to, int helper, int aoDisks) {
        if (aoDisks == 0) {
            return;
        }
        move(from, helper, to, aoDisks - 1);
        move(from, to);
        move(helper, to, from, aoDisks - 1);
    }

    public static void test() {
        int hanoi2[] = {};

    }
}