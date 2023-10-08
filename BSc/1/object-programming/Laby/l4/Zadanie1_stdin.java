package l4;

import java.util.Scanner;

class Zadanie1_stdin {

    public static void main(String[] args) {
        System.out.println("Proszę podać dwie liczby");
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int nwd = NWD.nwd(a, b);
        int nww = NWD.nww(a, b);
        System.out.printf("NWD: %d\n", nwd);
        System.out.printf("NWW: %d\n", nww);
        sc.close();
    }
}