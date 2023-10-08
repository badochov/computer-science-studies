package l5.zad1;

import java.util.Arrays;

public class BuforOsób {
    private Osoba[] osoby;
    private int liczbaOsób;
    private int liczbaUsuniętych;

    public BuforOsób() {
        this.osoby = new Osoba[10];
        this.liczbaOsób = 0;
        this.liczbaUsuniętych = 0;
    }

    public int getliczbaOsób() {
        return this.liczbaOsób;
    }

    public boolean buforPusty() {
        return this.liczbaOsób == 0;
    }

    public BuforOsób wstawPoczatek(Osoba o) {

        if (liczbaUsuniętych == 0) {
            if (this.liczbaOsób >= this.osoby.length) {
                this.osoby = Arrays.copyOf(this.osoby, this.osoby.length * 2);
            }
            System.arraycopy(this.osoby, 0, this.osoby, 1, this.liczbaOsób);
        } else {
            this.liczbaUsuniętych--;
        }

        this.osoby[this.liczbaUsuniętych] = o;
        this.liczbaOsób++;
        return this;
    }

    public BuforOsób wstawKoniec(Osoba o) {
        int index = this.liczbaOsób + this.liczbaUsuniętych;

        if (index >= this.osoby.length) {
            this.osoby = Arrays.copyOf(this.osoby, this.osoby.length * 2);
        }

        this.osoby[index] = o;
        this.liczbaOsób++;

        return this;
    }

    public Osoba usuńPoczątek() {
        this.liczbaOsób--;
        Osoba o = this.osoby[this.liczbaOsób];
        this.osoby[this.liczbaOsób] = null;
        this.liczbaUsuniętych++;
        return o;
    }

    public Osoba usuńKoniec() {
        this.liczbaOsób--;
        int index = this.liczbaOsób + this.liczbaUsuniętych;
        Osoba o = this.osoby[index];
        this.osoby[index] = null;
        return o;
    }
}