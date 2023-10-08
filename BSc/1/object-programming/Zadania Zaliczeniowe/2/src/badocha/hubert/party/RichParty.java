package badocha.hubert.party;

import badocha.hubert.Action;
import badocha.hubert.VoteCountingType;
import badocha.hubert.constituencies.Constituency;

/**
 * Klasa partii wybierającej akcje po najwyższej cenie
 */
public class RichParty extends MoneySensitiveParty {
    public RichParty(String partyName, int partyBudget, Action[] actionsAvailable,
                     Constituency[] constituenciesAvailable) {
        super(partyName, partyBudget, actionsAvailable, constituenciesAvailable);
    }

    @Override public void performAction(VoteCountingType voteCountingType) {
        performOperation(voteCountingType, -1);
    }

    @Override public RichParty copy(Constituency[] newConstituencies) {
        return new RichParty(getName(), budget, actions, newConstituencies);
    }
}
