package badocha.hubert;

public interface Gracz {
    boolean czyBierze();

    void otrzymajKartę(Karta k);

    int liczbaPunktów();

    void usuńRękę();
}
