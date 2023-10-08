package l5.zad2;

public class Ułamek {
    private int licznik;
    private int mianownik;

    public Ułamek(int licznik, int mianownik) throws Exception {
        this.licznik = licznik;
        this.mianownik = mianownik;

        this.sprawdzZero();
        this.skróć();
    }

    private void sprawdzZero() throws Exception {
        if (this.getMianownik() == 0) {
            throw new Exception("Dzielenie przez 0");
        }
    }

    public Ułamek dodaj(Ułamek u) {
        this.licznik *= u.mianownik;
        this.licznik += u.licznik * this.mianownik;

        this.mianownik *= u.mianownik;

        return this.skróć();
    }

    public Ułamek odejmij(Ułamek u) {
        this.licznik *= u.mianownik;
        this.licznik -= u.licznik * this.mianownik;

        this.mianownik *= u.mianownik;

        return this.skróć();
    }

    public Ułamek pomnóż(Ułamek u) {
        this.licznik *= u.licznik;
        this.mianownik *= u.mianownik;

        return this.skróć();
    }

    public Ułamek podziel(Ułamek u) throws Exception {
        this.licznik *= u.mianownik;
        this.mianownik *= u.licznik;

        this.sprawdzZero();

        return this.skróć();
    }

    static public Ułamek dodaj(Ułamek u, Ułamek u2) {
        return u.dodaj(u2);
    }

    static public Ułamek odejmij(Ułamek u, Ułamek u2) {
        return u.odejmij(u2);
    }

    static public Ułamek pomnóż(Ułamek u, Ułamek u2) {
        return u.pomnóż(u2);
    }

    static public Ułamek podziel(Ułamek u, Ułamek u2) throws Exception {
        return u.podziel(u2);
    }

    static public void wypisz(Ułamek u) {
        u.wypisz();
    }

    public void wypisz() {
        int pełna = this.licznik / this.mianownik;
        if (pełna == 0) {
            if (this.licznik == 0) {
                System.out.println(0);
            } else {
                System.out.printf("%d/%d\n", this.licznik, this.mianownik);
            }
        } else {
            int licznik = this.licznik - pełna * this.mianownik;
            if (licznik == 0) {
                System.out.println(pełna);
            } else if (licznik > 0) {
                System.out.printf("%d+%d/%d\n", pełna, licznik, this.mianownik);
            } else {
                System.out.printf("%d-%d/%d\n", pełna, Math.abs(licznik), this.mianownik);
            }
        }
    }

    private Ułamek skróć() {
        int nwd = this.NWD();
        this.mianownik /= nwd;
        this.licznik /= nwd;
        if (this.mianownik < 0) {
            this.mianownik *= -1;
            this.licznik *= -1;
        }
        return this;
    }

    private int NWD() {
        int a = this.mianownik;
        int b = this.licznik;
        while (b != 0) {
            int c = a % b;
            a = b;
            b = c;
        }
        return a;
    }

    public int getMianownik() {
        return this.mianownik;
    }

    public int getLicznik() {
        return this.licznik;
    }
}