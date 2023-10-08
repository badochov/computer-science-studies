package badocha.hubert.constituencies;

import badocha.hubert.Candidate;
import badocha.hubert.Vote;
import badocha.hubert.Votes;
import badocha.hubert.voter.CandidateVoter;
import badocha.hubert.voter.Voter;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

/**
 * Klasa reprezentująca pojedynczy okrąg.
 */
public class SingleConstituency extends AbstractConstituency {
    final private Map<String, ArrayList<Candidate>> candidates;
    final private ArrayList<Voter> voters;
    final private int size;
    final private int id;

    public SingleConstituency(int constituencySize, int constituencyId) {
        candidates = new HashMap<>();
        voters = new ArrayList<>();
        size = constituencySize;
        id = constituencyId;
    }

    public void addCandidate(String p, Candidate c) {
        ArrayList<Candidate> partyCandidates = candidates.getOrDefault(p, new ArrayList<>());
        partyCandidates.add(c);
        candidates.put(p, partyCandidates);
    }

    @Override
    public ArrayList<Candidate> getPartyCandidates(String partyName) {
        return candidates.get(partyName);
    }

    public void addVoter(Voter v) {
        voters.add(v);
    }

    @Override
    public int getSize() {
        return size;
    }

    @Override
    public ArrayList<Voter> getVoters() {
        return voters;
    }

    @Override public Constituency copy() {
        SingleConstituency constituency = new SingleConstituency(size, id);
        for (var partyCandidates : candidates.entrySet()) {
            for (Candidate candidate : partyCandidates.getValue()) {
                constituency.addCandidate(partyCandidates.getKey(), candidate);
            }
        }
        for (Voter voter : voters) {
            constituency.addVoter(voter.copy());
        }

        return constituency;
    }

    @Override
    public Map<String, ArrayList<Candidate>> getAllCandidates() {
        return candidates;
    }


    @Override
    public Votes getVotes(
            Map<String, ArrayList<Candidate>> allCandidates) {
        Votes votes = new Votes();
        for (Voter voter : voters) {
            Candidate candidate;
            if (voter instanceof CandidateVoter) {
                candidate = voter.getVote(candidates);
            } else {
                candidate = voter.getVote(allCandidates);
            }
            votes.addVote(new Vote(candidate, voter));
        }
        return votes;
    }

    @Override public String getName() {
        return String.valueOf(id);
    }
}
