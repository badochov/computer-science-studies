package badocha.hubert;

public class Limitowy extends GraczAbstrakcyjny {
    private final int limit;

    Limitowy(int limitPunktów) {
        limit = limitPunktów;
    }


    @Override
    public boolean czyBierze() {
        return liczbaPunktów() < limit;
    }
}
