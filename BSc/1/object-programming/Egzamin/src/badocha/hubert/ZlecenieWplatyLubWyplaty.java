package badocha.hubert;

public abstract class ZlecenieWplatyLubWyplaty extends ZlecenieAbstrakcyjne {
    private final String numerKonta;

    protected ZlecenieWplatyLubWyplaty(int kwota, String waluta, String numerKonta) {
        super(kwota, waluta);
        this.numerKonta = numerKonta;
    }

    public String getNumerKonta() {
        return numerKonta;
    }
}
