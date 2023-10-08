package badocha.hubert.voter;

import badocha.hubert.Candidate;
import badocha.hubert.Human;

import java.util.ArrayList;
import java.util.Map;

/**
 * Osoba głosująca na danego kandydata.
 */
public class CandidateVoter extends Human implements Voter {
    private final String party;
    private final int pos;

    public CandidateVoter(String name, String surname, String partyName, int position) {
        super(name, surname);
        party = partyName;
        pos = position;
    }

    @Override
    public Candidate getVote(Map<String, ArrayList<Candidate>> candidates) {
        return candidates.get(party).get(pos);
    }

    @Override public CandidateVoter copy() {
        return new CandidateVoter(name, surname, party, pos);
    }


}
