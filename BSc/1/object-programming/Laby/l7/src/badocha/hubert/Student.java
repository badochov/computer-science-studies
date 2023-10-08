package badocha.hubert;

public class Student extends Człowiek {
    private String nrIndeksu;

    Student(String imię, String nazwisko, String indeks) {
        super(imię, nazwisko);
        this.nrIndeksu = indeks;
    }

    @Override
    public String toString() {
        return super.toString() + " > " + "Student{" +
                "nrIndeksu='" + nrIndeksu + '\'' +
                '}';
    }
}
