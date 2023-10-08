package badocha.hubert.factories;

import badocha.hubert.Action;
import badocha.hubert.constituencies.Constituency;
import badocha.hubert.party.*;

/**
 * Klasa abstrakcyjna produkująca obiekty typu partii.
 */
public abstract class PartyFactory {

    public static Party[] convertToParties(String[] names, int[] budgets,
                                           PartyStrategy[] strategies, Action[] actions,
                                           Constituency[] constituencies) {
        int partiesCount = names.length;
        Party[] parties = new Party[partiesCount];
        for (int i = 0; i < partiesCount; i++) {
            parties[i] =
                    convertToParty(names[i], budgets[i], strategies[i], actions, constituencies);
        }

        return parties;
    }

    public static Party convertToParty(String name, int budget, PartyStrategy strategy,
                                       Action[] actions, Constituency[] constituencies) {

        switch (strategy) {
            case Cheap:
                return new CheapParty(name, budget, actions, constituencies);
            case OwnStrategy:
                return new OwnStrategyParty(name, budget, actions, constituencies);
            case Rich:
                return new RichParty(name, budget, actions, constituencies);
            case Greedy:
                return new GreedyParty(name, budget, actions, constituencies);
        }
        return null;
    }

}
