package badocha.hubert;

public class ZlecenieWplaty extends ZlecenieWplatyLubWyplaty {

    protected ZlecenieWplaty(int kwota, String waluta, String numerKonta) {
        super(kwota, waluta, numerKonta);
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public String nazwaOperacji() {
        return null;
    }
}
