package Transport;


public abstract class Lądowe extends Pojazd {
    public Lądowe(){
        super();
        this.jednostka_spalania = "l/hkm";
        this.przelicznik_spalania = 0.01f;
    }
}
