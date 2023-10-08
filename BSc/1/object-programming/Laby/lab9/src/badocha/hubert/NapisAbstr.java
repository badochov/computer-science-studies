package badocha.hubert;

abstract public class NapisAbstr implements Napis{
    abstract public int dlugosc();
    abstract public char dajZnak(int pozycja);
    abstract public boolean ustawZnak(int pozycja, char znak);
    abstract public void dodajNapis(Napis napis);
    abstract public void wypisz();
}
