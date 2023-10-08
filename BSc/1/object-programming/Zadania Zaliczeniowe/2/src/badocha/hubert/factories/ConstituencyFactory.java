package badocha.hubert.factories;

import badocha.hubert.constituencies.Constituency;
import badocha.hubert.constituencies.MultiConstituency;
import badocha.hubert.constituencies.SingleConstituency;

import java.util.ArrayList;
import java.util.Arrays;

/**
 * Klasa pomocnicza produkująca obiekty spełniające interfejs okręgu wyborczego.
 */
public abstract class ConstituencyFactory {
    public static Constituency[] createConstituencies(SingleConstituency[] singles,
                                                      int[][] mergedConstituencies) {
        ArrayList<Constituency> constituencies = new ArrayList<>();
        for (int[] merged : mergedConstituencies) {
            constituencies.add(new MultiConstituency(
                    new ArrayList<>(Arrays.asList(singles[merged[0]], singles[merged[1]]))));
        }
        for (int i = 0; i < singles.length; i++) {
            boolean isInMerged = false;
            for (int[] merged : mergedConstituencies) {
                for (int index : merged) {
                    if (i == index) {
                        isInMerged = true;
                        break;
                    }
                }
                if (isInMerged) {
                    break;
                }
            }
            if (!isInMerged) {
                constituencies.add(singles[i]);
            }
        }

        return constituencies.toArray(new Constituency[]{});
    }
}
