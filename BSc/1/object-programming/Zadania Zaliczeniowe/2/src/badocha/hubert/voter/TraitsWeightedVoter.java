package badocha.hubert.voter;

import badocha.hubert.Action;

import java.util.Arrays;

/**
 * Osoba która wartościuje kandydatów względem zestawu wag dla danych cech.
 */
public class TraitsWeightedVoter extends TraitsAwareVoter {
    public static final int MAX_TRAIT_VALUE = 100;
    public static final int MIN_TRAIT_VALUE = -100;

    public TraitsWeightedVoter(String name, String surname, int[] weights, String partyName) {
        super(name, surname, convertWeights(weights), partyName);
    }

    public TraitsWeightedVoter(String name, String surname, int[] weights) {
        super(name, surname, convertWeights(weights));
    }

    private static int[][] convertWeights(int[] weights) {
        int[][] converted = new int[weights.length][2];
        for (int i = 0; i < weights.length; i++) {
            converted[i] = new int[]{i, weights[i]};
        }

        return converted;
    }

    public int getNewWeightValue(int i, int change) {
        return Math.max(MIN_TRAIT_VALUE,
                Math.min(MAX_TRAIT_VALUE, traitWeights[i][1] + change));
    }

    public void performAction(Action action) {
        for (int i = 0; i < traitWeights.length; i++) {
            traitWeights[i][1] = getNewWeightValue(i, action.getChange(i));
        }
    }

    public int[] getWeights() {
        return Arrays.stream(traitWeights).mapToInt((w) -> w[1]).toArray();
    }

    @Override public TraitsWeightedVoter copy() {
        return new TraitsWeightedVoter(name, surname, getWeights(), party);
    }
}
