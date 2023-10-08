package badocha.hubert;

import java.util.Arrays;
import java.util.Random;


abstract public class Wirus {
    protected Kwas kwas;
    protected int mutacje;

    abstract public void mutuj();

    public int getLiczbaMutacjiWMiesiacu() {
        return this.mutacje;
    }

    public Kwas getKwas() {
        return this.kwas;
    }
}
