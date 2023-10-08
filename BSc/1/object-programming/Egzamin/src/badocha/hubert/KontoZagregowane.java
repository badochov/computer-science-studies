package badocha.hubert;

import java.util.Collection;

public class KontoZagregowane implements Konto {
    private final Collection<Konto> konta;
    private final String klient;
    private final String numerKonta;

    KontoZagregowane(Collection<Konto> konta, String klient, String numerKonta) {
        this.konta = konta;
        this.klient = klient;
        this.numerKonta = numerKonta;
    }

    @Override public String getNumer() {
        return numerKonta;
    }

    @Override public String getKlient() {
        return klient;
    }

    @Override public String getStanKonta() {
        String res = konta.stream()
                .reduce("", (stan, konto) -> stan + konto.getStanKonta() + System.lineSeparator(),
                        String::concat);
        if (res.length() > 0) {
            return res.substring(0, res.length() - 1);
        }
        return res;
    }

    @Override public boolean wplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return konta.stream().anyMatch((konto) -> konto.wplataMozliwa(zlecenieOperacji));
    }

    @Override public boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji) {
        return konta.stream().anyMatch((konto) -> konto.wyplataMozliwa(zlecenieOperacji));
    }

    @Override public boolean wplac(ZlecenieOperacji zlecenieOperacji) {
        return konta.stream().anyMatch((konto) -> konto.wplac(zlecenieOperacji));
    }

    @Override public boolean wyplac(ZlecenieOperacji zlecenieOperacji) {
        return konta.stream().anyMatch((konto) -> konto.wyplac(zlecenieOperacji));
    }

    @Override public boolean koniecMiesiaca() {
        return true;
    }

    //NIEZAIMPLEMENTOWANA METODA Z PROJEKTU
    @Override public boolean zamknijKonto() {
        return false;
    }
}
