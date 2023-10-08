package badocha.hubert;

import java.util.ArrayList;

/**
 * Klasa z pomocniczymi metodami.
 */
public abstract class Utils {
    public static <T> ArrayList<T> mergeLists(ArrayList<T> l1, ArrayList<T> l2) {
        l1.addAll(l2);
        return l1;
    }
}
