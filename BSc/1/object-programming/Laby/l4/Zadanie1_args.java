package l4;

import java.util.Scanner;

class Zadanie1_args {

    public static void main(String[] args) {
        int a = Integer.parseInt(args[0]);
        int b = Integer.parseInt(args[1]);
        int nwd = NWD.nwd(a, b);
        int nww = NWD.nww(a, b);
        System.out.printf("NWD: %d\n", nwd);
        System.out.printf("NWW: %d\n", nww);
    }
}