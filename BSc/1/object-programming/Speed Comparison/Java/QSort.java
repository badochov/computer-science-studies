import java.io.BufferedReader;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

public class QSort {
    private static void qs(int[] tab, int start, int size) {
        if (size <= 1) {
            return;
        }
        int smallerCount = 0;
        int biggerCount = 0;
        int[] smaller = new int[size];
        int[] bigger = new int[size];
        int d = (int) (Math.random() * (double) size) + start;
        for (int i = start; i < start + size; i++) {
            if (i != d) {
                if (tab[i] < tab[d]) {
                    smaller[smallerCount] = tab[i];
                    smallerCount++;
                } else {
                    bigger[biggerCount] = tab[i];
                    biggerCount++;
                }
            }
        }
        tab[start + smallerCount] = tab[d];
        for (int i = 0; i < smallerCount; i++) {
            tab[start + i] = smaller[i];
        }
        for (int i = 0; i < biggerCount; i++) {
            tab[start + smallerCount + 1 + i] = bigger[i];
        }
        qs(tab, start, smallerCount);
        qs(tab, start + smallerCount + 1, biggerCount);
    }

    public static void qsort(int[] tab) {
        qs(tab, 0, tab.length);
    }

    public static void main(String[] args) {
        File file = new File("test.txt");
        BufferedReader reader = null;
        try {
            reader = new BufferedReader(new FileReader(file));
            int n = Integer.parseInt(reader.readLine());
            int[] tab = new int[n];
            for (int i = 0; i < n; i++) {
                tab[i] = Integer.parseInt(reader.readLine());
            }
            qsort(tab);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            try {
                if (reader != null) {
                    reader.close();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

    }
}