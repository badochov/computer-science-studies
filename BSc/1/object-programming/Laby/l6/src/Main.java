public class Main {

    public static void main(String[] args) {
        Polynomial p = new Polynomial(new Integer[]{1,2,});
        Polynomial p2 = new Polynomial(p);
        Polynomial p3 = new Polynomial(0);
        p.print();
        p2.multiply(p);
        p2.print();
        p.add(p3);
        p.print();
        p3.print();

        System.out.println(p.valueIn(2));
        System.out.println(p2.valueIn(2));
    }
}
