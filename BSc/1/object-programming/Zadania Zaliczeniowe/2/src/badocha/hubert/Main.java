package badocha.hubert;

import badocha.hubert.constituencies.Constituency;
import badocha.hubert.constituencies.SingleConstituency;
import badocha.hubert.factories.ConstituencyFactory;
import badocha.hubert.factories.PartyFactory;
import badocha.hubert.party.Party;
import badocha.hubert.party.PartyStrategy;

import java.io.FileNotFoundException;

public class Main {

    public static void main(String[] args) {
        if (args.length != 1) {
            System.err.println(
                    "Program powinien zostać wywołany z jednym argumentem - ścieżką do pliku z danymi wejściowymi");
            System.exit(1);
        }
        try (Parser parser = new Parser(args[0])) {
            int[] generalData = parser.parseGeneralData();

            int n = generalData[0];
            int p = generalData[1];
            int d = generalData[2];
            int c = generalData[3];

            int[][] mergedConstituencies = parser.parseMergedConstituencies();
            String[] partiesNames = parser.parsePartiesNames(p);
            int[] partiesBudgets = parser.parsePartiesBudgets(p);
            PartyStrategy[] partiesStrategies = parser.parsePartiesStrategies(p);
            int[] votersInConstituency = parser.parseNumberOfVoters(n);
            SingleConstituency[] constituencies =
                    parser.parseConstituencies(c, votersInConstituency, p);

            Action[] actions = parser.parseActions(d, c);
            Party[] parties =
                    PartyFactory.convertToParties(partiesNames, partiesBudgets, partiesStrategies,
                            actions, constituencies);
            Constituency[] finalConstituencies =
                    ConstituencyFactory.createConstituencies(constituencies, mergedConstituencies);

            Simulation simulation = new Simulation(parties, finalConstituencies);
            simulation.simulate();

        } catch (FileNotFoundException e) {
            e.printStackTrace();
            System.exit(1);
        }

    }
}
