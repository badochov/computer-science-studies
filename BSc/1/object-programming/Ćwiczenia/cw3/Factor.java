package cw3;

import java.util.Arrays;
import java.util.Scanner;
import java.util.Vector;

class Factor {

    public static Vector<Integer> factor(int n, Vector<Integer> factors, Vector<Integer> primes) {
        if (n == 0) {
            return factors;
        }
        for (int prime : primes) {
            if (n % prime == 0) {
                factors.add(prime);
                return factor(n / prime, factors, primes);
            }
        }

        return factors;
    }

    public static Vector<Integer> factor(int n) {
        Vector<Integer> factors = new Vector<>();
        Vector<Integer> primes = Primes.to(n);
        return factor(n, factors, primes);
    }

    public static void test() {
        int factors2[] = { 2 };
        int factors1[] = {};
        int factors6[] = { 2, 3 };
        int factors8[] = { 2, 2, 2 };
        assert factor(2).equals(Util.toVec(factors2));
        assert factor(1).equals(Util.toVec(factors1));
        assert factor(6).equals(Util.toVec(factors6));
        assert factor(8).equals(Util.toVec(factors8));
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        System.out.println(factor(sc.nextInt()));
        sc.close();
    }
}