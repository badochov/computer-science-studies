import java.util.Scanner;
import java.util.Vector;

class Newton {
    public static void reduce(Vector<Integer> den, Vector<Integer> num) {
        for (int i = num.size() - 1; i >= 0; i--) {
            int number = num.get(i);
            for (int j = 0; j < den.size(); j++) {
                if (den.get(j) % number == 0) {
                    den.set(j, den.get(j) / number);
                    break;
                }
            }
        }
    }

    public static int choose(int n, int k) {
        if (k < 0 || (n > 0 && n < k)) {
            return 0;
        }

        Vector<Integer> den = new Vector<>();
        Vector<Integer> num = new Vector<>();

        for (int i = 0; i < k; i++) {
            den.add(n - i);
            num.add(i + 1);
        }

        reduce(den, num);

        int res = 1;
        for (int number : den) {
            res *= number;
        }
        return res;

    }

    public static void test() {
        assert choose(10, -1) == 0;
        assert choose(10, 1) == 10;
        assert choose(10, 0) == 1;
        assert choose(-2, 2) == 3;
        assert choose(-1, 1) == -1;
        assert choose(1, 1) == 1;
        assert choose(1, 2) == 0;
        assert choose(150, 4) == 20260275;
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();

        System.out.println(choose(a, b));
        sc.close();
    }
}