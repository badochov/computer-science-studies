package cw3;

import java.util.Arrays;
import java.util.Scanner;
import java.util.Vector;

class Primes {
    public static Vector<Integer> to(int n) {
        Vector<Integer> primes = new Vector<>();
        for (int i = 2; i <= n; i++) {
            boolean prime = true;
            for (int num : primes) {
                if (i % num == 0) {
                    prime = false;
                    break;
                }
            }
            if (prime) {
                primes.add(i);
            }
        }
        return primes;
    }

    public static void test() {
        int primesTo100[] = { 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83,
                89, 97 };
        assert to(100).equals(Util.toVec(primesTo100));
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        Vector<Integer> primes = to(sc.nextInt());
        sc.close();
        for (int num : primes) {
            System.out.printf("%d ", num);
        }
    }
}