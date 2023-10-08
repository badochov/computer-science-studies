package badocha.hubert;

public abstract class ZlecenieAbstrakcyjne implements ZlecenieOperacji {

    protected final int kwota;
    protected final String waluta;


    protected ZlecenieAbstrakcyjne(int kwota, String waluta) {
        this.kwota = kwota;
        this.waluta = waluta;
    }


    @Override public int getKwota() {
        return kwota;
    }

    @Override public String getWaluta() {
        return waluta;
    }
}
