package badocha.hubert;

/**
 * Kandydat startujący w wyborach.
 */
public final class Candidate extends Human {

    private final int[] traits;
    private final String party;

    Candidate(String name, String surname, int[] ts, String partyName) {
        super(name, surname);
        traits = ts;
        party = partyName;
    }

    public int getTrait(int t) {
        return traits[t];
    }


    public String getParty() {
        return party;
    }
}
