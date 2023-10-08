package badocha.hubert;

public class Nauczyciel extends Pracownik {
    private String ulubionyPrzedmiot;

    Nauczyciel(String imię, String nazwisko, String pracodawca, String ulubionyJezyk) {
        super(imię, nazwisko, pracodawca);

        this.ulubionyPrzedmiot = ulubionyJezyk;
    }

    @Override
    public String toString() {
        return super.toString() + " > " + "Nauczyciel{" +
                "ulubionyPrzedmiot='" + ulubionyPrzedmiot + '\'' +
                '}';
    }
}
