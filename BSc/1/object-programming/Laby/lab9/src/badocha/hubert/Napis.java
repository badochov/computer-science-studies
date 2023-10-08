package badocha.hubert;

public interface Napis {
    int dlugosc();
    char dajZnak(int pozycja);
    boolean ustawZnak(int pozycja, char znak);
    void dodajNapis(Napis napis);
    void wypisz();
}
