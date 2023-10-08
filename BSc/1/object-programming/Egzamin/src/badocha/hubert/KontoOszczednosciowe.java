package badocha.hubert;

public class KontoOszczednosciowe extends KontoDoPrzechowywaniaSrodkowFinansowych {
    private final float rocznaStopaProcentowa;

    KontoOszczednosciowe(String numer, String klient, String waluta, float rocznaStopaProcentowa) {
        super(numer, klient, waluta, 0);
        this.rocznaStopaProcentowa = rocznaStopaProcentowa;
    }

    @Override public boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return false;
    }

    @Override public boolean koniecMiesiaca() {
        this.saldo = Math.round((float) saldo * (rocznaStopaProcentowa / 12));
        return true;
    }


    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public boolean zamknijKonto() {
        return false;
    }
}
