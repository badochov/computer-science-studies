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

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();

        System.out.println(nwd(a, b));
        sc.close();
    }

    public static void test() {
        assert nwd(3, 9) == 3;
        assert nwd(3, 1) == 1;
        assert nwd(21, 37) == 1;
    }
}