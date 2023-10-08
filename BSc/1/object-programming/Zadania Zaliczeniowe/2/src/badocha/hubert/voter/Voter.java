package badocha.hubert.voter;

import badocha.hubert.Candidate;

import java.util.ArrayList;
import java.util.Map;

/**
 * Interfejs osoby głosującej.
 */
public interface Voter {
    Candidate getVote(Map<String, ArrayList<Candidate>> candidates);

    Voter copy();
}
