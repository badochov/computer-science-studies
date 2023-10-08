package l4;

import java.util.Scanner;

class NWD {
    public static int nwd(int a, int b) {
        while (b != 0) {
            int c = a % b;
            a = b;
            b = c;
        }
        return a;
    }

    public static int nww(int a, int b) {
        return a / nwd(a, b) * b;
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int nwd = nwd(a, b);
        int nww = nww(a, b);
        System.out.printf("NWD: %d\n", nwd);
        System.out.printf("NWW: %d\n", nww);
        sc.close();
    }

    public static void test() {
        assert nwd(3, 9) == 3;
        assert nww(3, 9) == 9;
        assert nwd(3, 1) == 1;
        assert nwd(21, 37) == 1;
        assert nww(21, 37) == 21 * 37;
    }
}