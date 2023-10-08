package badocha.hubert.voter;

/**
 * Osoba głosująca względem minimalnej wartości danej cechy.
 */
public class MinTraitVoter extends TraitsAwareVoter {
    public MinTraitVoter(String name, String surname, int traitNumber, String partyName) {
        super(name, surname, new int[][]{{traitNumber, -1}}, partyName);
    }

    public MinTraitVoter(String name, String surname, int traitNumber) {
        super(name, surname, new int[][]{{traitNumber, -1}});
    }

    @Override public MinTraitVoter copy() {
        return new MinTraitVoter(name, surname, traitWeights[0][0], party);
    }
}
