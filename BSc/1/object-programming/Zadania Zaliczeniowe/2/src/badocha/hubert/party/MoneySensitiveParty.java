package badocha.hubert.party;

import badocha.hubert.Action;
import badocha.hubert.VoteCountingType;
import badocha.hubert.constituencies.Constituency;

import java.util.Comparator;

/**
 * Klasa abstrakcyjna partii wybierającej akcje względem ceny
 */
public abstract class MoneySensitiveParty extends Party {
    MoneySensitiveParty(String partyName, int partyBudget, Action[] actionsAvailable,
                        Constituency[] constituenciesAvailable) {
        super(partyName, partyBudget, actionsAvailable, constituenciesAvailable);
    }

    protected void performOperation(VoteCountingType voteCountingType, int operationNumber) {
        var availableActions = getAvailableActions();
        availableActions.sort(Comparator.comparingInt(Party::totalPrice));
        int index = (availableActions.size() + operationNumber) % availableActions.size();

        applyAction(availableActions.get(index));
    }
}
