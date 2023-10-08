package badocha.hubert;

public class Kwas {
    private String kwas;

    Kwas(String k) {
        this.kwas = k;
    }

    public char getNukleotyd(int i) {
        return this.kwas.charAt(i);
    }

    public void setNukleotyd(int i, char n) {
        this.kwas = this.kwas.substring(0, i) + n + this.kwas.substring(i + 1);
    }

    @Override
    public String toString() {
        return this.kwas;
    }

    public int dlugosc() {
        return this.kwas.length();
    }
}
