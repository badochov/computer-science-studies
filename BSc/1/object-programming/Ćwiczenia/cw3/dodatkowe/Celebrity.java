package cw3.dodatkowe;

import java.util.Arrays;
import java.util.Scanner;

class Celebrity {

    private static boolean check(boolean[][] knows, int sP, boolean[] possible) {
        for (int i = sP + 1; i < possible.length; i++) {
            boolean isKnown = knows[sP][i];
            if (isKnown) {
                possible[sP] = false;
                for (int j = sP + 1; j < possible.length; j++) {
                    if (possible[j]) {
                        sP = j;
                        break;
                    }
                }
                if (!possible[sP]) {
                    return false;
                }
                return check(knows, sP, possible);
            } else {
                possible[i] = false;
            }
        }
        for (int i = 0; i < possible.length; i++) {
            if (i == sP) {
                continue;
            }
            if (knows[sP][i] || !knows[i][sP]) {
                return false;
            }

        }
        return true;
    }

    public static boolean hasPersonality(boolean[][] knows) {
        int smallestPossible = 0;
        boolean possible[] = new boolean[knows.length];
        Arrays.fill(possible, Boolean.TRUE);
        return check(knows, smallestPossible, possible);
    }

    private static void test() {
        boolean test0[][] = { { true, false }, { false, true } };
        boolean test05[][] = { { true, false }, { true, false } };
        boolean test1[][] = { { true, false }, { true, true } };
        boolean test2[][] = { { true, false, true }, { true, true, true }, { false, false, true } };
        boolean test3[][] = { { true, false, true }, { true, true, true }, { false, true, true } };
        assert hasPersonality(test0) == false;
        assert hasPersonality(test05);
        assert hasPersonality(test1);
        assert hasPersonality(test2);
        assert hasPersonality(test3) == false;

    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        int size = sc.nextInt();
        boolean arr[][] = new boolean[size][size];
        for (int i = 0; i < size; i++) {
            for (int j = 0; i < size; j++) {
                arr[i][j] = sc.nextBoolean();
            }
        }
        System.out.println(hasPersonality(arr));
        sc.close();
    }
}