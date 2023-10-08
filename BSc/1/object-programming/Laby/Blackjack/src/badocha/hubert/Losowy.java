package badocha.hubert;

import java.util.Random;

public class Losowy extends GraczAbstrakcyjny {
    Losowy() {
    }

    @Override
    public boolean czyBierze() {
        return new Random().nextInt(2) == 1;
    }
}
