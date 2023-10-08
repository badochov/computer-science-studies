package badocha.hubert;

public class KontoKartowe extends KontoKredytowe {
    private final int limitKredytowy;

    KontoKartowe(String numer, String klient, String waluta, int limitKredytowy,
                 KontoBiezace kontoBiezace) {
        super(numer, klient, waluta, 0, kontoBiezace);
        this.limitKredytowy = limitKredytowy;
    }

    @Override public boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return super.wyplataMozliwa(zlecenieOperacji) &&
                saldo - zlecenieOperacji.getKwota() >= limitKredytowy;
    }

    @Override public boolean koniecMiesiaca() {
        return splac(-saldo);
    }


    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public boolean zamknijKonto() {
        return false;
    }
}
