package cw3.hanoi;

import java.util.Vector;

class Main {

    public static void main(String[] args) {
        String[] rodNames = { "A", "B", "C" };
        Integer[] diskNames = { 1, 2, 3 };
        Hanoi<String, Integer> hanoi = new Hanoi<>(rodNames, diskNames);
        var moves = hanoi.solve();
        Move.print(moves);
    }
}