package cw3.dodatkowe;

import java.lang.reflect.Array;
import java.util.Arrays;

class Polinomial {
    private Number[] coefs;

    public Polinomial(Number[] coefs) {
        this.coefs = coefs;
        this.remZeroes();
    }

    public void remZeroes() {
        int maxNonZeroCoef = this.coefs.length;
        for (int i = this.coefs.length - 1; i >= 0; i--) {
            if (this.coefs[i].floatValue() != 0.0f) {
                break;
            }
            maxNonZeroCoef--;
        }
        this.coefs = Arrays.copyOf(this.coefs, maxNonZeroCoef);
    }

    public void substract(Polinomial p1) {
        int l = Math.max(this.coefs.length, p1.coefs.length);
        this.coefs = Arrays.copyOf(this.coefs, l);
        for (int i = 0; i < p1.coefs.length; i++) {
            this.coefs[i] = this.coefs[i].floatValue() - p1.coefs[i].floatValue();
        }

        this.remZeroes();
    }

    public void printCoefs() {
        for (Number n : this.coefs) {
            System.out.print(n);
            System.out.print(" ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        float nums1[] = { 1.0f, 1.0f, 2.0f, 0.0f };
        float nums2[] = { 1.0f, 2.0f, 2.0f, 0.0f };
        Polynomial p = new Polynomial(nums1);
        Polynomial p1 = new Polynomial(nums2);
        p.printCoefs();
        p.substract(p1);
        p.printCoefs();
    }
}