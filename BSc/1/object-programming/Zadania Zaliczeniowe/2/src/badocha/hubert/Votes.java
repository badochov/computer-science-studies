package badocha.hubert;

import badocha.hubert.voter.Voter;

import java.util.*;

/**
 * Klasa trzymająca wszystkie głosy w danym okręgu.
 */
public class Votes implements Iterable<Vote> {
    private final ArrayList<Vote> votes;

    public Votes() {
        votes = new ArrayList<>();
    }

    public boolean addVote(Vote vote) {
        return votes.add(vote);
    }

    public ArrayList<Vote> getVotes() {
        return votes;
    }

    public boolean addVotes(Votes newVotes) {
        return votes.addAll(newVotes.votes);
    }

    public Map<Candidate, ArrayList<Voter>> getCandidatesVotes() {
        Map<Candidate, ArrayList<Voter>> votesForCandidate = new HashMap<>();
        votes.forEach((vote) -> votesForCandidate.merge(vote.getCandidate(), new ArrayList<Voter>(
                Collections.singletonList(vote.getVoter())), Utils::mergeLists));

        return votesForCandidate;
    }

    public Map<String, Integer> getPartyVotes() {
        Map<String, Integer> partyVotes = new HashMap<>();
        votes.forEach((vote) -> partyVotes.merge(vote.getCandidate().getParty(), 1, Integer::sum));
        return partyVotes;
    }

    @Override public Iterator<Vote> iterator() {
        return votes.iterator();
    }
}
