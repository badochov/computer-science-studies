package badocha.hubert;

public class ZlecenieWyplaty extends ZlecenieWplatyLubWyplaty {

    protected ZlecenieWyplaty(int kwota, String waluta, String numerKonta) {
        super(kwota, waluta, numerKonta);
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public String nazwaOperacji() {
        return null;
    }
}
