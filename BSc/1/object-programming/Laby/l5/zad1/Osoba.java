package l5.zad1;

public class Osoba {
    private String imie;
    private String nazwisko;

    public Osoba(String imie, String nazwisko) {
        this.imie = imie;
        this.nazwisko = nazwisko;
    }

    public String getImie() {
        return this.imie;
    }

    public String getNazwisko() {
        return this.nazwisko;
    }

    public Osoba setImie(String imie) {
        this.imie = imie;
        return this;
    }

    public Osoba setNazwisko(String nazwisko) {
        this.nazwisko = nazwisko;
        return this;
    }
}