package badocha.hubert;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Wirus w;

        Scanner sc = new Scanner(System.in);

        System.out.println("Proszę podać szczepionke i nastepnie wcisnąć enter");
        String s = sc.next();

        System.out.println("Proszę podać kwas nukleinowy wirusa i wcisnąć enter");
        String kwas = sc.next();

        System.out.println("Proszę podać typ wirusa(0 - dziwny, 1 - specyficzny)");
        int typ = sc.nextInt();

        System.out.println("Proszę podać liczbę mutacji w miesiącu");
        int mutacje = sc.nextInt();

        if (typ == 1) {
            System.out.println("Proszę podać liczbę dozwolonych indeksów");
            int n = sc.nextInt();
            int[] indeksy = new int[n];
            System.out.println("Proszę podać dozwolone indeksy(numerowane od 0) oddzielone spacjami");
            for (int i = 0; i < n; i++) {
                indeksy[i] = sc.nextInt();
            }
            System.out.println("Proszę podać dozwolone nukletydy(połączone ze sobą)");
            String nuk = sc.next();
            w = new WirusSpecyficzny(kwas, mutacje, indeksy, nuk);
        } else {
            w = new WirusDziwny(kwas, mutacje);
        }
        System.out.println("Proszę podać liczbę miesięcy");
        int miesiace = sc.nextInt();

        sc.close();

        Szczepionka sz = new Szczepionka(s);

        Symulacja sym = new Symulacja(w, sz);
        sym.symuluj(miesiace);

        System.out.println();
        System.out.println();
        System.out.println("Wirus po podanym okresie:");
        System.out.println(sym.getWirus().getKwas());
        if(sym.zwalcza()){
            System.out.println("Szczepionka ciągle zwalcza wirusa! :)");
        }
        else{
            System.out.println("Szczepionka już nie zwalcza wirusa! :(");
        }

    }
}
