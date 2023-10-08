package badocha.hubert;

public class NapisTab extends NapisAbstr {
    private char[] znaki;
    private int dlugosc;

    NapisTab(char[] znaki) {
        this.znaki = znaki;
        dlugosc = znaki.length;
    }

    @Override
    public int dlugosc() {
        return dlugosc;
    }

    @Override
    public char dajZnak(int pozycja) {
        if (pozycja < dlugosc && pozycja >= 0) {
            return znaki[pozycja];
        }
        return (char) -1;
    }

    @Override
    public boolean ustawZnak(int pozycja, char znak) {
        if (pozycja < 0 || pozycja > dlugosc) {
            return false;
        }
        powiększ(pozycja + 1);
        znaki[pozycja] = znak;
        if (pozycja > dlugosc) {
            dlugosc = pozycja;
        }
        return true;
    }

    @Override
    public void dodajNapis(Napis napis) {
        int rozmiar = dlugosc + napis.dlugosc();
        powiększ(rozmiar);
        for(int i =0;i<napis.dlugosc();i++){
            znaki[dlugosc+i] = napis.dajZnak(i);
        }
    }

    @Override
    public void wypisz() {
        System.out.println(znaki);
    }

    private void powiększ(int rozmiar) {
        if (rozmiar > znaki.length) {
            char[] temp = new char[rozmiar];
            System.arraycopy(znaki, 0, temp, 0, znaki.length);
            znaki = temp;
        }
    }
}
