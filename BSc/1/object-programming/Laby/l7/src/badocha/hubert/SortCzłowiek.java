package badocha.hubert;

import java.util.ArrayList;

public class SortCzłowiek {
    public static ArrayList<Człowiek> sort(ArrayList<Człowiek> lista) {
        for (int i = 0; i < lista.size() - 1; i++) {
            int min = i;
            for (int j = i + 1; j < lista.size(); j++) {
                if (lista.get(min).compareTo(lista.get(j)) > 0) {
                    min = j;
                }
            }
            Człowiek temp = lista.get(i);
            lista.set(i, lista.get(min));
            lista.set(min, temp);
        }

        return lista;
    }

    public static Człowiek[] sort(Człowiek[] lista) {
        for (int i = 0; i < lista.length - 1; i++) {
            int min = i;
            for (int j = i + 1; j < lista.length; j++) {
                if (lista[min].compareTo(lista[j]) > 0) {
                    min = j;
                }
            }
            Człowiek temp = lista[i];
            lista[i] = lista[min];
            lista[min] = temp;
        }

        return lista;
    }
}
