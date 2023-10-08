package badocha.hubert;

public class Informatyk extends Pracownik {
    private String ulubionyJezyk;

    Informatyk(String imię, String nazwisko, String pracodawca, String ulubionyJezyk) {
        super(imię, nazwisko, pracodawca);

        this.ulubionyJezyk = ulubionyJezyk;
    }

    @Override
    public String toString() {
        return super.toString() + " > " + "Informatyk{" +
                "ulubionyJezyk='" + ulubionyJezyk + '\'' +
                '}';
    }
}
