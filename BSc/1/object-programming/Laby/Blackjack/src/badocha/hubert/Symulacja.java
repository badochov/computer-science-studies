package badocha.hubert;

import java.util.ArrayList;

public class Symulacja {
    private final int[] punkty;
    private final Gracz[] gracze;
    private final int tury;

    Symulacja(Gracz[] graczeWRozgrywce, int liczbaTur) {
        punkty = new int[graczeWRozgrywce.length];
        gracze = graczeWRozgrywce;
        tury = liczbaTur;
    }

    public void symuluj() {
        Oczko gra = new Oczko(gracze);
        for (int i = 0; i < tury; i++) {
            ArrayList<Integer> zwycięzcy = gra.rozegraj();
            gra.resetuj();

            for (int numerGracza : zwycięzcy) {
                punkty[numerGracza]++;
            }
        }
    }

    public void wypiszWyniki() {
        System.out.println("Wyniki:");
        for (int i = 0; i < gracze.length; i++) {
            System.out.printf("Gracz %d: %d" + System.lineSeparator(), i + 1, punkty[i]);
        }
    }
}
