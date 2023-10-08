package badocha.hubert;

public class Main {

    public static void main(String[] args) {
        Gracz[] gracze = {
                new Krupier(),
                new Limitowy(21),
                new Limitowy(15),
                new Limitowy(16),
                new Limitowy(18),
                new Limitowy(19),
                new Limitowy(13),
                new Losowy(),
                new Ostrożny(),
                new Śpiący(),
        };

        int liczbaTur = 1000000;

        Symulacja sym = new Symulacja(gracze, liczbaTur);
        sym.symuluj();
        sym.wypiszWyniki();
    }
}
