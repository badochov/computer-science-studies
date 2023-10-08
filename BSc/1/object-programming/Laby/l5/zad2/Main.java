
package l5.zad2;

public class Main {
    public static void main(String[] args) {
        try {
            Ułamek u1 = new Ułamek(4, 2);
            Ułamek u2 = new Ułamek(6, 9);
            Ułamek u3 = new Ułamek(13, 37);
            Ułamek u4 = new Ułamek(21, 37);

            u1.wypisz();
            u2.wypisz();
            u3.wypisz();
            u4.wypisz();

            u1.dodaj(u2);
            u1.wypisz();

            Ułamek.dodaj(u1, u2);
            u1.wypisz();

            u2.odejmij(u3);
            u2.wypisz();

            Ułamek.odejmij(u2, u3);
            u2.wypisz();

            u3.pomnóż(u4);
            u3.wypisz();

            Ułamek.pomnóż(u3, u4);
            u3.wypisz();

            u4.podziel(u1);
            u4.wypisz();

            Ułamek.podziel(u4, u1);
            u4.wypisz();

            new Ułamek(1, 0);

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}