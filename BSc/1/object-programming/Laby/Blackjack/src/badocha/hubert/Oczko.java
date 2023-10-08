package badocha.hubert;

import java.util.ArrayList;
import java.util.Random;

public class Oczko {
    public static int kopieKarty = 4;
    public static int oczko = 21;
    private final Gracz[] gracze;
    private ArrayList<Karta> karty;

    Oczko(Gracz[] graczeWRozgrywce) {
        gracze = graczeWRozgrywce;

        resetuj();
    }

    public ArrayList<Integer> rozegraj() {
        Random r = new Random();

        ArrayList<Integer> zwycięzcy = new ArrayList<>();
        boolean ktośDobrał = true;
        while (ktośDobrał && karty.size() > 0) {
            ktośDobrał = false;
            for (Gracz gracz : gracze) {
                if (gracz.czyBierze()) {
                    ktośDobrał = true;

                    int numerKarty = r.nextInt(karty.size());
                    Karta k = karty.remove(numerKarty);
                    gracz.otrzymajKartę(k);
                }
            }
        }
        int maksPunktów = 0;

        for(int i=0;i<gracze.length;i++){
            int punkty =gracze[i].liczbaPunktów();

            if(punkty > oczko){
                continue;
            }
            if(punkty > maksPunktów){
                maksPunktów = punkty;
                zwycięzcy.clear();
                zwycięzcy.add(i);
            }
            else if(punkty == maksPunktów){
                zwycięzcy.add(i);
            }
        }

        return zwycięzcy;
    }

    public void resetuj() {
        for (Gracz g : gracze) {
            g.usuńRękę();
        }

        karty = new ArrayList<>();
        for (Karta k : Karta.values()) {
            for (int i = 0; i < kopieKarty; i++) {
                karty.add(k);
            }
        }
    }

}
