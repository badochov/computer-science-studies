package badocha.hubert;

abstract public class KontoDoPrzechowywaniaSrodkowFinansowych extends KontoAbstrakcyjne {
    KontoDoPrzechowywaniaSrodkowFinansowych(String numer, String klient, String waluta, int saldo) {
        super(numer, klient, waluta, saldo);
    }
}
