package badocha.hubert;

import java.io.File;
import java.io.FileNotFoundException;
import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        System.out.println("Proszę podać plik(puste pole oznacza stdin)");
        Scanner sc = new Scanner(System.in);
        String file = sc.nextLine();
        ArrayList<Człowiek> lista;
        if (file.compareTo("") == 0) {
            System.out.println("Wybrano stdin, proszę podać dane a następnie wcisnąś ctrl+d");
            lista = wczytajLudzi(sc);
        } else {
            System.out.printf("Wybrano: %s%s", file, System.lineSeparator());
            File f = new File(file);
            Scanner s = null;
            try {
                s = new Scanner(f);
            } catch (FileNotFoundException e) {
                e.printStackTrace();
                return;
            }
            lista = wczytajLudzi(s);
            s.close();
        }


        wypiszLudzi(SortCzłowiek.sort(lista));
        sc.close();
    }

    private static void wypiszLudzi(ArrayList<Człowiek> lista) {
        for (Człowiek c : lista) {
            System.out.println(c);
        }
    }

    private static ArrayList<Człowiek> wczytajLudzi(Scanner sc) {
        ArrayList<Człowiek> lista = new ArrayList<>();
        while (sc.hasNext()) {
            String type = sc.next();
            String imię = sc.next();
            String nazwisko = sc.next();

            switch (type) {
                case "C":
                    lista.add(new Człowiek(imię, nazwisko));
                    break;
                case "S":
                    String indeks = sc.next();
                    lista.add(new Student(imię, nazwisko, indeks));
                    break;
                case "P": {
                    String pracodawca = sc.next();
                    lista.add(new Pracownik(imię, nazwisko, pracodawca));
                    break;
                }
                case "I": {
                    String pracodawca = sc.next();
                    String jezyk = sc.next();
                    lista.add(new Informatyk(imię, nazwisko, pracodawca, jezyk));
                    break;
                }
                case "N": {
                    String pracodawca = sc.next();
                    String przedmiot = sc.next();
                    lista.add(new Nauczyciel(imię, nazwisko, pracodawca, przedmiot));
                    break;
                }
            }
        }

        return lista;
    }


}
