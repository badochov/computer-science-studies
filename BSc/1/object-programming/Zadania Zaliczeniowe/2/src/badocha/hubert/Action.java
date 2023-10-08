package badocha.hubert;

/**
 * Akcja którą partia może wykonać w kampanii.
 */
public class Action {
    private final int[] changes;

    Action(int[] traitChanges) {
        changes = traitChanges;
    }

    public int getPrice() {
        int sum = 0;
        for (int change : changes) {
            sum += Math.abs(change);
        }
        return sum;
    }

    public int getChange(int i) {
        return changes[i];
    }

}
