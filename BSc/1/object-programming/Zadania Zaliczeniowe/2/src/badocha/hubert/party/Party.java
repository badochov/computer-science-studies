package badocha.hubert.party;

import badocha.hubert.Action;
import badocha.hubert.Pair;
import badocha.hubert.VoteCountingType;
import badocha.hubert.constituencies.Constituency;
import badocha.hubert.voter.TraitsWeightedVoter;
import badocha.hubert.voter.Voter;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;

/**
 * Klasa abstrakcyjna z której dziedziczą wszystkie partię.
 */
public abstract class Party {
    protected final Action[] actions;
    protected final Constituency[] constituencies;
    private final String name;
    protected int budget;

    Party(String partyName, int partyBudget, Action[] actionsAvailable,
          Constituency[] constituenciesAvailable) {
        name = partyName;
        budget = partyBudget;
        actions = getActionsSortedByPrice(actionsAvailable);
        constituencies = getConstituenciesSortedByPrice(constituenciesAvailable);
    }

    static protected int totalPrice(Pair<Action, Constituency> p) {
        return totalPrice(p.getFirst(), p.getSecond());
    }

    static protected int totalPrice(Action action, Constituency constituency) {
        return action.getPrice() * constituency.getMandatesCount();
    }

    protected ArrayList<Pair<Action, Constituency>> getAvailableActions() {
        ArrayList<Pair<Action, Constituency>> pairs = new ArrayList<>();
        for (Action action : actions) {
            for (Constituency constituency : constituencies) {
                if (totalPrice(action, constituency) <= budget) {
                    pairs.add(new Pair<>(action, constituency));
                }
            }
        }
        return pairs;
    }

    public String getName() {
        return name;
    }

    abstract public void performAction(VoteCountingType voteCountingType);

    public final boolean canPerformAction() {
        return budget >= totalPrice(actions[0], constituencies[0]);
    }

    private Action[] getActionsSortedByPrice(Action[] actions) {
        return Arrays.stream(actions).sorted(Comparator.comparingInt(Action::getPrice))
                .toArray(Action[]::new);
    }

    private Constituency[] getConstituenciesSortedByPrice(Constituency[] actions) {
        return Arrays.stream(actions).sorted(Comparator.comparingInt(Constituency::getSize))
                .toArray(Constituency[]::new);
    }

    protected void applyAction(Pair<Action, Constituency> pair) {
        for (Voter voter : pair.getSecond().getVoters()) {
            if (voter instanceof TraitsWeightedVoter) {
                ((TraitsWeightedVoter) voter).performAction(pair.getFirst());
            }
        }

        budget -= pair.getFirst().getPrice();
    }

    abstract public Party copy(Constituency[] newConstituencies);
}
