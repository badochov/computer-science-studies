package badocha.hubert;

public interface Konto {
    String getNumer();

    String getKlient();

    String getStanKonta();

    boolean wplataMozliwa(ZlecenieOperacji zlecenieOperacji);

    boolean wyplataMozliwa(ZlecenieOperacji zlecenieOperacji);

    boolean wplac(ZlecenieOperacji zlecenieOperacji);

    boolean wyplac(ZlecenieOperacji zlecenieOperacji);

    boolean koniecMiesiaca();

    boolean zamknijKonto();
}
