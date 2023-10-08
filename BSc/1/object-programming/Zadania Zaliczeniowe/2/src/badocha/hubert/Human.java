package badocha.hubert;

/**
 * Klasa abstrakcyjna człowieka.
 */
public abstract class Human {
    protected final String name;
    protected final String surname;

    public Human(String first, String sur) {
        name = first;
        surname = sur;
    }

    @Override
    public String toString() {
        return name + " " + surname;
    }
}
