package badocha.hubert;

public class KontoRatalne extends KontoKredytowe {
    private final int liczbaRat;
    private final int kwotaPozyczki;
    private final float roczneOprocentowanieRat;

    KontoRatalne(String numer, String klient, String waluta, int kwotaPozyczki,
                 KontoBiezace kontoBiezace, int liczbaRat, float roczneOprocentowanieRat) {
        super(numer, klient, waluta, -kwotaPozyczki, kontoBiezace);
        this.liczbaRat = liczbaRat;
        this.roczneOprocentowanieRat = roczneOprocentowanieRat;
        this.kwotaPozyczki = kwotaPozyczki;
    }

    @Override public boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return false;
    }

    private float getRataKapitalowa() {
        return (float) kwotaPozyczki / liczbaRat;
    }

    private float getRataOdsetkowa() {
        return -(float) saldo * (roczneOprocentowanieRat / 12);
    }


    @Override public boolean koniecMiesiaca() {
        int rata = Math.round(getRataKapitalowa() + getRataOdsetkowa());
        return splac(rata);
    }


    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public boolean zamknijKonto() {
        return false;
    }
}
