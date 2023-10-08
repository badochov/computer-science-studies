package badocha.hubert;

public class Ostrożny extends GraczAbstrakcyjny {
    public final int maks_karta;
    public final int maks_punktów;

    //może nie działać poprawnie dla wartości większych niż 44
    Ostrożny(int maks_punktów_w_grze) {
        maks_punktów = maks_punktów_w_grze;
        int maks = 0;
        for (Karta k : Karta.values()) {
            if (k.punkty > maks) {
                maks = k.punkty;
            }
        }
        maks_karta = maks;
    }

    Ostrożny() {
        this(Oczko.oczko);
    }

    @Override
    public boolean czyBierze() {

        return liczbaPunktów() <= maks_punktów - maks_karta;
    }
}
