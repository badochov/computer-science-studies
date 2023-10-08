package badocha.hubert.voter;

import badocha.hubert.Candidate;
import badocha.hubert.Human;

import java.util.ArrayList;
import java.util.Map;
import java.util.Random;

/**
 * Osoba głosująca na losowego kandydata z danej partii.
 */
public class PartyVoter extends Human implements Voter {
    private final String party;

    public PartyVoter(String name, String surname, String partyName) {
        super(name, surname);
        party = partyName;
    }

    @Override
    public Candidate getVote(Map<String, ArrayList<Candidate>> candidates) {
        ArrayList<Candidate> partyCandidates = candidates.get(party);
        return partyCandidates.get(new Random().nextInt(partyCandidates.size()));
    }

    @Override public PartyVoter copy() {
        return new PartyVoter(name, surname, party);
    }
}
