package badocha.hubert;

abstract public class KontoKredytowe extends KontoAbstrakcyjne {
    protected final KontoBiezace kontoBiezace;

    KontoKredytowe(String numer, String klient, String waluta, int saldo,
                   KontoBiezace kontoBiezace) {
        super(numer, klient, waluta, saldo);
        this.kontoBiezace = kontoBiezace;
    }

    @Override public final boolean wplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return false;
    }

    protected final boolean splac(int kwota) {
        boolean mozliwe = kontoBiezace.wyplac(new ZleceniePrzelewu(kwota, getWaluta(), getNumer(),
                kontoBiezace.getNumer()));
        if (mozliwe) {
            saldo += kwota;
            return true;
        }
        return false;
    }

}
