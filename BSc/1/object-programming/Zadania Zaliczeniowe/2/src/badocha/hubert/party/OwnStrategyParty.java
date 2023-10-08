package badocha.hubert.party;

import badocha.hubert.Action;
import badocha.hubert.VoteCountingType;
import badocha.hubert.constituencies.Constituency;

import java.util.Random;

/**
 * Klasa partii implementująca własną strategię.
 */
public class OwnStrategyParty extends Party {
    public OwnStrategyParty(String partyName, int partyBudget, Action[] actionsAvailable,
                            Constituency[] constituenciesAvailable) {
        super(partyName, partyBudget, actionsAvailable, constituenciesAvailable);
    }

    @Override public void performAction(VoteCountingType voteCountingType) {
        var availableActions = getAvailableActions();
        int index = new Random().nextInt(availableActions.size());
        applyAction(availableActions.get(index));
    }

    @Override public OwnStrategyParty copy(Constituency[] newConstituencies) {
        return new OwnStrategyParty(getName(), budget, actions, newConstituencies);
    }
}
