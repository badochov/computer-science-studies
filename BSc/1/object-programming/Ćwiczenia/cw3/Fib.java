import java.util.Scanner;

class Fib {
    public static int fib(int n) {
        int a = 0;
        int b = 1;
        for (int i = 0; i < n; i++) {
            int c = b + a;
            a = b;
            b = c;
        }
        return a;
    }

    public static void test() {
        assert fib(0) == 0;
        assert fib(1) == 1;
        assert fib(2) == 1;
        assert fib(3) == 2;
        assert fib(42) == 267914296;
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        System.out.println(fib(sc.nextInt()));
        sc.close();
    }
}