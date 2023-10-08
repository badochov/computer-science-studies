package badocha.hubert;

public class Człowiek implements Comparable<Człowiek> {
    private String imię;
    private String nazwisko;

    Człowiek(String imię, String nazwisko) {
        this.imię = imię;
        this.nazwisko = nazwisko;
    }

    @Override
    public String toString() {
        return "Czlowiek{" +
                "imie='" + imię + '\'' +
                ", nazwisko='" + nazwisko + '\'' +
                '}';
    }


    @Override
    public int compareTo(Człowiek człowiek) {
        int cmp = this.nazwisko.compareTo(człowiek.nazwisko);
        return cmp == 0 ? this.imię.compareTo(człowiek.imię) : cmp;
    }
}
