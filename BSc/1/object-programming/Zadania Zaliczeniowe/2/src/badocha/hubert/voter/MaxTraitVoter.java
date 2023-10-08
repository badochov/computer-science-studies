package badocha.hubert.voter;

/**
 * Osoba głosująca względem maksymalnej wartości danej cechy.
 */
public class MaxTraitVoter extends TraitsAwareVoter {
    public MaxTraitVoter(String name, String surname, int traitNumber, String partyName) {
        super(name, surname, new int[][]{{traitNumber, 1}}, partyName);
    }

    public MaxTraitVoter(String name, String surname, int traitNumber) {
        super(name, surname, new int[][]{{traitNumber, 1}});
    }

    @Override public MaxTraitVoter copy() {
        return new MaxTraitVoter(name, surname, traitWeights[0][0], party);
    }
}
