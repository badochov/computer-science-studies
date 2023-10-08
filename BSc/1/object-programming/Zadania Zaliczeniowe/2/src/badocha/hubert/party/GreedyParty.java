package badocha.hubert.party;

import badocha.hubert.Action;
import badocha.hubert.Candidate;
import badocha.hubert.Pair;
import badocha.hubert.VoteCountingType;
import badocha.hubert.constituencies.Constituency;
import badocha.hubert.voter.TraitsWeightedVoter;
import badocha.hubert.voter.Voter;

/**
 * Klasa partii wybierającej akcje zachłannie
 */
public class GreedyParty extends Party {
    public GreedyParty(String partyName, int partyBudget, Action[] actionsAvailable,
                       Constituency[] constituenciesAvailable) {
        super(partyName, partyBudget, actionsAvailable, constituenciesAvailable);
    }

    @Override public void performAction(VoteCountingType voteCountingType) {
        var availableActions = getAvailableActions();

        var bestAction = getAvailableActions().get(0);
        int maxValue = Integer.MIN_VALUE;
        for (var action : availableActions) {
            int value = calculateActionValue(action);
            if (value > maxValue) {
                maxValue = value;
                bestAction = action;
            }
        }

        applyAction(bestAction);
    }

    private int calculateActionValue(Pair<Action, Constituency> pair) {
        int sum = 0;
        for (Voter voter : pair.getSecond().getVoters()) {
            if (voter instanceof TraitsWeightedVoter) {
                int[] weights = ((TraitsWeightedVoter) voter).getWeights();
                for (int i = 0; i < weights.length; i++) {
                    weights[i] = ((TraitsWeightedVoter) voter)
                            .getNewWeightValue(i, pair.getFirst().getChange(i));
                }
                sum += calculateWeightedSum(weights, pair.getSecond());
            }
        }
        return sum;
    }

    private int calculateWeightedSum(int[] weights, Constituency constituency) {
        int sum = 0;
        for (Candidate candidate : constituency.getPartyCandidates(getName())) {
            for (int i = 0; i < weights.length; i++) {
                sum += weights[i] * candidate.getTrait(i);
            }
        }

        return sum;
    }

    @Override public GreedyParty copy(Constituency[] newConstituencies) {
        return new GreedyParty(getName(), budget, actions, newConstituencies);
    }
}
