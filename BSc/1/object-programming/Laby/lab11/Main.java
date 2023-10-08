package PO_lab11;

class BłądOtwierania extends Exception {
}

class BłądCzytania extends Exception {
}

class Zasób1 implements AutoCloseable {
    private boolean czyMa = false;

    public Zasób1(String nazwa) throws BłądOtwierania {
        if (Math.random() > 0.75)
            throw new BłądOtwierania();
    }

    public boolean maLiczbę() {
        return czyMa = Math.random() > 0.5;
    }

    public int dajLiczbę() throws BłądCzytania {
        if (!czyMa)
            throw new BłądCzytania();
        else
            return 42;
    }

    @Override
    public void close() throws Exception {
        System.out.println("Zasób1 zamknięty!");
    }
}

class Zasób2 implements AutoCloseable {

    public void zróbCoś() {
    }

    @Override
    public void close() throws Exception {
        System.out.println("Zasób2 zamknięty!");
    }

}

class dziwnaKlasa {
    public int mojaMetoda(String[] nazwyPlików, int k) throws BłądCzytania, BłądOtwierania {
        // szuka w podanych plikach pierwszej liczby którą dzieli k
        try{
        Zasób2 z2 = new Zasób2();
        if (k == 0)
            return -1;

        for (int i = 0; i < nazwyPlików.length; i++) {
            Zasób1 z = new Zasób1(nazwyPlików[i]);   // może być wyjątek

            int wyn = 0;
            while (z.maLiczbę()) {
                wyn = z.dajLiczbę();
                if (wyn % k == 0)
                    return wyn;
            }
        }
    }
        return 0;
    }
}

public class Main {
    public static void main(String[] args) throws BłądCzytania, BłądOtwierania {
        dziwnaKlasa m = new dziwnaKlasa();
        m.mojaMetoda(new String[] { "", "", "", "" }, 13);
    }
}
