package Transport;

public abstract class Morskie extends Pojazd {

    public Morskie(){
        super();
        this.jednostka_spalania = "gal/mil";
        this.przelicznik_spalania=5.23f;
    }
}
