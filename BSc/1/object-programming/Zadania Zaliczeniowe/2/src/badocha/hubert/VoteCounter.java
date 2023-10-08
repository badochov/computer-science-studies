package badocha.hubert;

import badocha.hubert.constituencies.Constituency;

import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

/**
 * Klasa zliczająca głosy.
 */
public abstract class VoteCounter {

    public static Map<String, Integer> getNumberOfMandates(VoteCountingType voteCountingType,
                                                           Map<String, Integer> partyVotes,
                                                           Constituency constituency) {
        switch (voteCountingType) {
            case DHondt:
                return getNumberOfMandatesDHondt(partyVotes, constituency);
            case Sainte_Lague:
                return getNumberOfMandatesSainteLague(partyVotes, constituency);
            case Hare_Niemeyer:
                return getNumberOfMandatesHareNiemeyer(partyVotes, constituency);
        }
        return null;
    }

    private static Map<String, Integer> getNumberOfMandatesDHondt(Map<String, Integer> partyVotes,
                                                                  Constituency constituency) {
        return getNumberOfMandatesIterativeMethod(partyVotes, constituency, (i) -> i + 1);
    }

    private static Map<String, Integer> getNumberOfMandatesSainteLague(
            Map<String, Integer> partyVotes,
            Constituency constituency) {
        return getNumberOfMandatesIterativeMethod(partyVotes, constituency, (i) -> 2 * i + 1);
    }

    private static Map<String, Integer> getNumberOfMandatesIterativeMethod(
            Map<String, Integer> partyVotes,
            Constituency constituency,
            Function<Integer, Integer> f) {
        Map<String, Integer> mandates = new HashMap<>();
        Function<String, Integer> getCoefficient =
                (String name) -> partyVotes.getOrDefault(name, 0) /
                        f.apply(mandates.getOrDefault(name, 0));
        for (int i = 1; i <= constituency.getMandatesCount(); i++) {
            String partyNameMax = partyVotes.keySet().toArray(String[]::new)[0];
            for (String partyName : partyVotes.keySet()) {
                if (getCoefficient.apply(partyName) > getCoefficient.apply(partyNameMax)) {
                    partyNameMax = partyName;
                }
            }

            mandates.merge(partyNameMax, 1, Integer::sum);
        }


        return mandates;
    }


    private static double valueAfterComa(double d) {
        return d - (int) d;
    }

    private static int comparePartyVotes(Map.Entry<String, Double> e1,
                                         Map.Entry<String, Double> e2) {
        return Double.compare(valueAfterComa(e1.getValue()), valueAfterComa(e1.getValue()));
    }

    private static Map<String, Integer> getNumberOfMandatesHareNiemeyer(
            Map<String, Integer> partyVotes,
            Constituency constituency) {
        Map<String, Double> Qs = new HashMap<>();
        partyVotes.forEach(
                (party, votes) -> Qs.put(party,
                        ((double) constituency.getMandatesCount() * votes) /
                                constituency.getNumberOfVotes()));
        Map<String, Integer> mandates = new HashMap<>();
        Qs.forEach((party, Q) -> mandates.put(party, Q.intValue()));
        int sum = Qs.values().stream().mapToInt(Double::intValue).reduce(0, Integer::sum);
        if (sum < constituency.getMandatesCount()) {
            Qs.entrySet().stream().sorted(VoteCounter::comparePartyVotes)
                    .limit(constituency.getMandatesCount() - sum)
                    .forEach((e) -> mandates.merge(e.getKey(), 1, Integer::sum));
        }
        return mandates;
    }
}
