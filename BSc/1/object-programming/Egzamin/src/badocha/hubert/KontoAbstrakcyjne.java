package badocha.hubert;

public abstract class KontoAbstrakcyjne implements Konto {
    private final String numer;
    private final String klient;
    private final String waluta;
    protected int saldo;

    KontoAbstrakcyjne(String numer, String klient, String waluta, int saldo) {
        this.numer = numer;
        this.klient = klient;
        this.waluta = waluta;
        this.saldo = saldo;
    }

    @Override public final String getKlient() {
        return klient;
    }

    protected final String getWaluta() {
        return waluta;
    }

    @Override public final String getNumer() {
        return numer;
    }

    @Override public boolean wplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return waluta.equals(zlecenieOperacji.getWaluta());
    }

    @Override public boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return waluta.equals(zlecenieOperacji.getWaluta());
    }

    @Override public final boolean wplac(ZlecenieOperacji zlecenieOperacji) {
        if (wplataMozliwa(zlecenieOperacji)) {
            saldo += zlecenieOperacji.getKwota();
            return true;
        }
        return false;
    }

    @Override public final boolean wyplac(ZlecenieOperacji zlecenieOperacji) {
        if (wyplataMozliwa(zlecenieOperacji)) {
            saldo -= zlecenieOperacji.getKwota();
            return true;
        }
        return false;
    }

    @Override public final String getStanKonta() {
        return saldo + waluta;
    }
}
