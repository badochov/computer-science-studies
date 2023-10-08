package cw3;

import java.util.Vector;

class Util {

    public static Vector<Integer> toVec(int[] a) {
        Vector<Integer> v = new Vector<>(a.length);

        for (int num : a) {
            v.add(num);
        }

        return v;
    }
}