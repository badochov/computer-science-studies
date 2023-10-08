package Transport;

class Przedstaw{
    private static void przedstawPojazd(Pojazd p){
        System.out.printf("Jestem: %s\n", p.getNazwa());
        System.out.printf("Spalam: %s\n", p.getSpalanieFormat());
    }

    public static void main(String[] args) {
        SamochódOsobowy samochod = new SamochódOsobowy();
        StatekParowy lotniskowiec = new StatekParowy();
        SamolotOdrzutowy samolot = new SamolotOdrzutowy();

        przedstawPojazd(samochod);
        przedstawPojazd(lotniskowiec);
        przedstawPojazd(samolot);
    }
}