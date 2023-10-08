package badocha.hubert.voter;

import badocha.hubert.Candidate;
import badocha.hubert.Human;
import badocha.hubert.Utils;

import java.util.ArrayList;
import java.util.Map;
import java.util.Random;

/**
 * Klasa abstrakcyjna reprezentująca osobę która wartościuje kandydatów względem zestawu wag dla danych cech.
 */
public abstract class TraitsAwareVoter extends Human implements Voter {
    protected final int[][] traitWeights;
    protected final String party;

    TraitsAwareVoter(String name, String surname, int[][] weights, String partyName) {
        super(name, surname);
        traitWeights = weights;
        party = partyName;
    }

    TraitsAwareVoter(String name, String surname, int[][] traitNumber) {
        this(name, surname, traitNumber, "");
    }

    private int getTraitsSum(Candidate c) {
        int sum = 0;
        for (int[] weight : traitWeights) {
            sum += c.getTrait(weight[0]) * weight[1];
        }
        return sum;
    }

    protected ArrayList<Candidate> getCandidatesWithMaxScore(
            Map<String, ArrayList<Candidate>> candidates) {
        ArrayList<Candidate> availableCandidates;
        if (party.equals("")) {
            availableCandidates =
                    candidates.values().stream().reduce(new ArrayList<>(), Utils::mergeLists);
        } else {
            availableCandidates = candidates.get(party);
        }
        ArrayList<Candidate> candidatesWithMaxScore = new ArrayList<>();

        int maxSum = Integer.MIN_VALUE;
        for (Candidate candidate : availableCandidates) {
            int sum = getTraitsSum(candidate);
            if (sum == maxSum) {
                candidatesWithMaxScore.add(candidate);
            } else if (sum > maxSum) {
                maxSum = sum;
                candidatesWithMaxScore.clear();
                candidatesWithMaxScore.add(candidate);
            }
        }

        return candidatesWithMaxScore;
    }


    @Override
    public Candidate getVote(Map<String, ArrayList<Candidate>> candidates) {
        ArrayList<Candidate> candidatesWithMaxScore = getCandidatesWithMaxScore(candidates);
        return candidatesWithMaxScore.get(new Random().nextInt(candidatesWithMaxScore.size()));
    }
}
