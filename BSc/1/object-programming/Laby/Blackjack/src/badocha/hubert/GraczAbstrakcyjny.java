package badocha.hubert;

import java.util.ArrayList;

abstract public class GraczAbstrakcyjny implements Gracz {
    private final ArrayList<Karta> karty;

    GraczAbstrakcyjny() {
        karty = new ArrayList<>();
    }

    @Override
    public void otrzymajKartę(Karta k) {
        this.karty.add(k);
    }

    @Override
    public int liczbaPunktów() {
        return karty.stream().mapToInt(karta -> karta.punkty).sum();
    }

    @Override
    public void usuńRękę() {
        karty.clear();
    }
}
