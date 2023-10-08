package badocha.hubert;

public class Pracownik extends Człowiek {
    private String pracodawca;

    Pracownik(String imię, String nazwisko, String pracodawca) {
        super(imię, nazwisko);

        this.pracodawca = pracodawca;
    }

    @Override
    public String toString() {
        return super.toString() + " > " + "Pracownik{" +
                "pracodawca='" + pracodawca + '\'' +
                '}';
    }
}
