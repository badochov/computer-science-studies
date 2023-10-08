package Transport;

public abstract class Powietrzne extends Pojazd {

    public Powietrzne() {
        this.jednostka_spalania = "t/mil";
        this.przelicznik_spalania = 10;
    }
}
