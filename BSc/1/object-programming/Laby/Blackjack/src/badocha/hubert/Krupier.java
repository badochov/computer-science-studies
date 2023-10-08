package badocha.hubert;

public class Krupier extends GraczAbstrakcyjny {
    Krupier() {
    }

    @Override
    public boolean czyBierze() {
        return liczbaPunktów() < 17;
    }

}
