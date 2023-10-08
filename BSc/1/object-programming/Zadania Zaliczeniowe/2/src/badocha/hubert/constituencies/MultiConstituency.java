package badocha.hubert.constituencies;

import badocha.hubert.Candidate;
import badocha.hubert.Utils;
import badocha.hubert.Votes;
import badocha.hubert.voter.Voter;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

/**
 * Klasa reprezentująca scalony okrąg.
 */
public class MultiConstituency extends AbstractConstituency {
    private final ArrayList<Constituency> constituencies;

    public MultiConstituency(ArrayList<Constituency> mergedConstituencies) {
        constituencies = mergedConstituencies;
    }

    @Override
    public Map<String, ArrayList<Candidate>> getAllCandidates() {
        Map<String, ArrayList<Candidate>> candidates = new HashMap<>();
        for (Constituency constituency : constituencies) {
            for (var partyCandidates : constituency.getAllCandidates().entrySet()) {
                candidates.merge(partyCandidates.getKey(), partyCandidates.getValue(),
                        Utils::mergeLists);
            }
        }
        return candidates;
    }

    @Override
    public ArrayList<Candidate> getPartyCandidates(String partyName) {
        ArrayList<Candidate> candidates = new ArrayList<>();
        for (Constituency constituency : constituencies) {
            candidates.addAll(constituency.getPartyCandidates(partyName));
        }
        return candidates;
    }

    @Override
    public Votes getVotes(
            Map<String, ArrayList<Candidate>> allCandidates) {
        Votes votes = new Votes();
        for (Constituency constituency : constituencies) {
            votes.addVotes(constituency.getVotes(allCandidates));
        }
        return votes;
    }

    @Override
    public int getSize() {
        return constituencies.stream().map(Constituency::getSize).reduce(0, Integer::sum);
    }

    @Override
    public ArrayList<Voter> getVoters() {
        return constituencies.stream().map(Constituency::getVoters)
                .reduce(new ArrayList<>(), Utils::mergeLists);
    }

    @Override public MultiConstituency copy() {
        return new MultiConstituency((ArrayList<Constituency>) constituencies.stream().map(
                Constituency::copy).collect(Collectors.toList()));
    }

    @Override public String getName() {
        final String[] name = {"("};
        constituencies.forEach(constituency -> name[0] += constituency.getName() + " + ");
        return name[0].substring(0, name[0].length() - 3) + ")";
    }

}
