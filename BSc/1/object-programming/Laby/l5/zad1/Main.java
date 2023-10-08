package l5.zad1;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        BuforOsób b = new BuforOsób();
        Scanner sc = new Scanner(System.in);
        System.out.println(
                "Proszę podać liste osób, w formacie jedna osoba na linię, imię i nazwisko oddzielone spacjami, po poadniu listy prosze wprowadzić pustą linijkę");
        while (true) {
            String line = sc.nextLine();
            if (line.equals("")) {
                break;
            }
            String tokens[] = line.split(" +");
            if (tokens.length != 2) {
                System.out.println("Zła liczba argumentów !");
                continue;
            }
            Osoba o = new Osoba(tokens[0], tokens[1]);
            b.wstawKoniec(o);

        }
        sc.close();

        while (!b.buforPusty()) {
            Osoba o = b.usuńKoniec();
            System.out.printf("%s %s\n", o.getImie(), o.getNazwisko());
        }
    }
}