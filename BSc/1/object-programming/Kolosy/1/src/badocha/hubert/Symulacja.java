package badocha.hubert;

public class Symulacja {
    private Szczepionka szczepionka;
    private Wirus wirus;

    Symulacja(Wirus wirus, Szczepionka szczepionka) {
        this.wirus = wirus;
        this.szczepionka = szczepionka;
    }

    public void symuluj(int liczbaMiesiecy) {
        for (int i = 0; i < liczbaMiesiecy; i++) {
            for (int j = 0; j < this.wirus.getLiczbaMutacjiWMiesiacu(); j++) {
                this.wirus.mutuj();
            }
        }
    }

    public boolean zwalcza() {
        return this.szczepionka.zwalczaWirusa(this.wirus);
    }

    public Wirus getWirus() {
        return this.wirus;
    }


}
