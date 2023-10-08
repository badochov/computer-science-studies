package badocha.hubert;

import java.util.Random;

public class WirusDziwny extends Wirus {

    WirusDziwny(String kwas, int mutacjeNaMiesiac) {
        this.kwas = new Kwas(kwas);
        this.mutacje = mutacjeNaMiesiac;
    }

    @Override
    public void mutuj() {
        if (this.kwas.dlugosc() == 1) {
            return;
        }
        Random r = new Random();
        int a = r.nextInt(this.kwas.dlugosc());
        int b = r.nextInt(this.kwas.dlugosc() - 1);
        if (b >= a) {
            b++;
        }
        char temp = this.kwas.getNukleotyd(a);

        this.kwas.setNukleotyd(a, this.kwas.getNukleotyd(b));
        this.kwas.setNukleotyd(b, temp);
    }
}
