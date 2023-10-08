package badocha.hubert;

public class Szczepionka {
    private Kwas ciag;

    Szczepionka(String szukanyCiag) {
        this.ciag = new Kwas(szukanyCiag);
    }

    public boolean zwalczaWirusa(Wirus w) {
        Kwas kwas = w.getKwas();
        for (int i = 0; i <= kwas.dlugosc() - this.ciag.dlugosc(); i++) {
            boolean zwalcza = true;
            for (int j = 0; j < this.ciag.dlugosc(); j++) {
                char c = this.ciag.getNukleotyd(j);
                if (kwas.getNukleotyd(i+j) != c) {
                    zwalcza = false;
                    break;
                }
            }
            if (zwalcza) {
                return true;
            }
        }
        return false;
    }
}
