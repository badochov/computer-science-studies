package badocha.hubert;

/**
 * Typ wyliczeniowy mówiący o typie zliczania głosów w danych wyborach.
 */
public enum VoteCountingType {
    DHondt("Metoda D'Hondta"), Sainte_Lague("Metoda Sainte-Laguë"),
    Hare_Niemeyer("Metoda Hare’a-Niemeyera");

    private final String description;

    VoteCountingType(String desc) {
        description = desc;
    }

    @Override public String toString() {
        return description;
    }
}
