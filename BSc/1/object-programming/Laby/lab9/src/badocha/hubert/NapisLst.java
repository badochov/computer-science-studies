package badocha.hubert;

import java.util.List;

public class NapisLst extends NapisAbstr {
    private List<Character> znaki;

    NapisLst(List<Character> znaki) {
        this.znaki = znaki;
    }

    @Override
    public int dlugosc() {
        return znaki.size();
    }

    @Override
    public char dajZnak(int pozycja) {
        return znaki.get(pozycja);
    }

    @Override
    public boolean ustawZnak(int pozycja, char znak) {
        if (pozycja < 0 || pozycja > znaki.size()) {
            return false;
        }
        znaki.set(pozycja, znak);
        return true;
    }

    @Override
    public void dodajNapis(Napis napis) {
        for (int i = 0; i < napis.dlugosc(); i++) {
            znaki.add(znaki.size(), napis.dajZnak(i));
        }
    }

    @Override
    public void wypisz() {
        System.out.println(znaki);
    }
}
