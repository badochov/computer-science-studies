package badocha.hubert;

import java.util.Random;

public class WirusSpecyficzny extends Wirus {
    private int[] pozycje;
    private Kwas nukleotydy;

    WirusSpecyficzny(String kwas, int mutacjeNaMiesiac, int[] dozwolonePozycje, String nukleotydy) {
        this.kwas = new Kwas(kwas);
        this.mutacje = mutacjeNaMiesiac;
        this.pozycje = dozwolonePozycje;
        this.nukleotydy = new Kwas(nukleotydy);
    }

    @Override
    public void mutuj() {
        if (this.pozycje.length == 0) {
            return;
        }
        Random r = new Random();
        int a = this.pozycje[r.nextInt(this.pozycje.length)];
        while (true) {
            char c = this.nukleotydy.getNukleotyd(r.nextInt(this.nukleotydy.dlugosc()));
            if (c != this.kwas.getNukleotyd(a)) {
                this.kwas.setNukleotyd(a, c);
                break;
            }
        }
    }
}
