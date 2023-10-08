package cw3.dodatkowe;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Scanner;

class Intersection {
    public static <T> T[] intersect(T[] a1, T[] a2) {
        HashSet<T> h1 = new HashSet<>(Arrays.asList(a1));
        HashSet<T> h2 = new HashSet<>(Arrays.asList(a2));
        h1.retainAll(h2);
        T[] temp = Arrays.copyOf(a1, 0);
        return h1.toArray(temp);
    }

    private static void test() {
        Integer arr1[] = { 1, 2 };
        Integer arr2[] = { 3, 2 };
        Integer[] ans = { 2 };
        assert Arrays.equals(intersect(arr1, arr2), ans);
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        int s1 = sc.nextInt();
        int s2 = sc.nextInt();
        Integer a1[] = new Integer[s1];
        Integer a2[] = new Integer[s2];
        for (int i = 0; i < s1; i++) {
            a1[i] = sc.nextInt();
        }
        for (int i = 0; i < s2; i++) {
            a2[i] = sc.nextInt();
        }
        Integer[] i = intersect(a1, a2);
        for (Integer a : i) {
            System.out.println(a);
        }

        sc.close();
    }
}