package badocha.hubert.constituencies;

import badocha.hubert.Candidate;
import badocha.hubert.Votes;
import badocha.hubert.voter.Voter;

import java.util.ArrayList;
import java.util.Map;

/**
 * Interfejs okręgu wyborczego.
 */
public interface Constituency {
    Map<String, ArrayList<Candidate>> getAllCandidates();

    ArrayList<Candidate> getPartyCandidates(String partyName);

    Votes getVotes();

    Votes getVotes(Map<String, ArrayList<Candidate>> allCandidates);

    int getSize();

    int getNumberOfVotes();

    ArrayList<Voter> getVoters();

    Constituency copy();

    String getName();

    int getMandatesCount();
}
