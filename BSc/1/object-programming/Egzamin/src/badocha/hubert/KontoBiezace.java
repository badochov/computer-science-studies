package badocha.hubert;

public class KontoBiezace extends KontoDoPrzechowywaniaSrodkowFinansowych {
    KontoBiezace(String numer, String klient, String waluta) {
        super(numer, klient, waluta, 0);
    }

    @Override public boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return super.wyplataMozliwa(zlecenieOperacji) && zlecenieOperacji.getKwota() <= saldo;
    }

    @Override public boolean koniecMiesiaca() {
        return true;
    }


    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public boolean zamknijKonto() {
        return false;
    }
}
