package l4;

import java.util.Scanner;

class Voting {

    private static int findCand(int votes[]) {
        int cand = 0;
        for (int i = 1; i < votes.length; i++) {
            if (votes[i] == votes[cand]) {
                continue;
            }
            cand++;
        }
        return votes[cand];
    }

    // Todo make it work
    public static int winnner(int votes[]) {
        int cand = findCand(votes);
        int count = 0;
        for (int vote : votes) {
            if (vote == cand) {
                count++;
            }
        }
        if (count > votes.length / 2) {
            return cand;
        }
        return 0;
    }

    private static void test() {
        int test0[] = { 1, 1 };
        int test1[] = { 1, 2 };
        int test2[] = { 1, 1, 1, 2 };
        int test3[] = { 1, 1, 2, 2 };
        int test4[] = { 999999999, 999999999 };
        int test5[] = { 1, 2, 2, 2, 2, 3 };
        assert winnner(test0) == 1;
        assert winnner(test1) == 0;
        assert winnner(test2) == 1;
        assert winnner(test3) == 0;
        assert winnner(test4) == 999999999;
        assert winnner(test5) == 2;
    }

    public static void main(String[] args) {
        test();
        Scanner sc = new Scanner(System.in);
        System.out.println("Proszę podać liczbe głosów");
        int s = sc.nextInt();
        System.out.println("Prosze podać głosy");
        int votes[] = new int[s];
        for (int i = 0; i < s; i++) {
            votes[i] = sc.nextInt();
        }
        System.out.println("Zwycięzca to (0 jeżeli nie ma):");
        System.out.println(winnner(votes));
        sc.close();
    }

}