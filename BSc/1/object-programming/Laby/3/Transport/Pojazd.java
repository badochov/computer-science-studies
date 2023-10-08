package Transport;

public abstract class Pojazd {
    // Attributes
    protected int spalanie;
    protected String jednostka_spalania;
    protected float przelicznik_spalania;
    protected String nazwa;

    public String getNazwa(){
        return this.nazwa;
    }
    public float getSpalanie(){
        return this.spalanie;
    }
    public String getJednostka_spalania(){
        return this.jednostka_spalania;
    }
    public float getPrzelicznik_spalania(){
        return this.przelicznik_spalania;
    }

    public String getSpalanieFormat(){
        String spal = String.valueOf(this.getSpalanie());
        return spal.concat(" ").
                concat(this.getJednostka_spalania());
    }
    // Operations
//    public abstract int ile_litrow_paliwa (Miejsce start, Miejsce koniec);
}
