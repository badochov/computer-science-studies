package badocha.hubert;

import badocha.hubert.voter.Voter;

/**
 * Głos oddany.
 */
public class Vote {
    private final Candidate candidate;
    private final Voter voter;

    public Vote(Candidate candidateChosen, Voter voterVoting) {
        candidate = candidateChosen;
        voter = voterVoting;
    }

    public Voter getVoter() {
        return voter;
    }

    public Candidate getCandidate() {
        return candidate;
    }
}
