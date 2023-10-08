package badocha.hubert;

import badocha.hubert.constituencies.Constituency;
import badocha.hubert.party.Party;
import badocha.hubert.voter.Voter;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Klasa symulująca przebieg wyborów.
 */
public class Simulation {
    private final Party[] parties;
    private final Constituency[] constituencies;

    Simulation(Party[] partiesInSimulation, Constituency[] constituenciesInSimulation) {
        parties = partiesInSimulation;
        constituencies = constituenciesInSimulation;
    }


    public void simulate() {
        for (VoteCountingType voteCountingType : VoteCountingType.values()) {
            simulate(voteCountingType);
            System.out.println("====================================");
            System.out.println();
        }
    }

    private Party[] deepCloneParties(Constituency[] constituenciesClone) {
        return Arrays.stream(parties).map(p -> p.copy(constituenciesClone)).toArray(Party[]::new);
    }

    private Constituency[] deepCloneConstituencies() {
        return Arrays.stream(constituencies).map(Constituency::copy).toArray(Constituency[]::new);
    }

    public void simulate(VoteCountingType voteCountingType) {
        Constituency[] constituenciesClone = deepCloneConstituencies();
        Party[] partiesCopy = deepCloneParties(constituenciesClone);

        boolean canAnyPartyPerformOperation = true;
        while (canAnyPartyPerformOperation) {
            canAnyPartyPerformOperation = false;
            for (Party party : partiesCopy) {
                if (party.canPerformAction()) {
                    canAnyPartyPerformOperation = true;
                    party.performAction(voteCountingType);
                }
            }
        }

        displayResults(voteCountingType, constituencies);
    }


    private void displayResults(VoteCountingType voteCountingType, Constituency[] constituencies) {
        System.out.println(voteCountingType);
        System.out.println();


        Map<String, Integer> results = new HashMap<>();
        for (Constituency constituency : constituencies) {
            var votes = constituency.getVotes();
            var constituencyResults =
                    VoteCounter.getNumberOfMandates(voteCountingType, votes.getPartyVotes(),
                            constituency);
            assert constituencyResults != null;
            constituencyResults
                    .forEach((key, value) -> results.merge(key, value, Integer::sum));

            System.out.println(constituency.getName());
            displayVotersAndTheirVotes(votes);
            displayCandidatesAndNumberOfVotes(votes.getCandidatesVotes());
            displayResults(constituencyResults);
            System.out.println("------------------------------------");
        }
        displayResults(results);


    }

    private void displayVotersAndTheirVotes(Votes votes) {
        System.out.println();
        System.out.println("Wyborcy i ich głosy:");
        System.out.println();

        votes.forEach((vote) -> System.out.println(vote.getVoter() + " -> " + vote.getCandidate()));
    }

    private void displayCandidatesAndNumberOfVotes(Map<Candidate, ArrayList<Voter>> votes) {
        System.out.println();
        System.out.println("Kandydaci i liczba głosów na nich:");
        System.out.println();

        votes.forEach(
                ((candidate, voters) -> System.out.println(candidate + " -> " + voters.size())));
    }


    private void displayResults(Map<String, Integer> results) {
        System.out.println();
        System.out.println("Partie i liczba mandatów z okręgu:");
        System.out.println();

        results.forEach((partyName, score) -> System.out.println(partyName + " -> " + score));
    }
}
