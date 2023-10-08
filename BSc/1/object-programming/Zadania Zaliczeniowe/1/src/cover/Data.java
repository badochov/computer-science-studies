package cover;

public class Data {
    private Query query;
    private SolutionSet set;

    Data(Query query, SolutionSet set) {
        this.query = query;
        this.set = set;
    }

    public Query getQuery() {
        return this.query;
    }

    public SolutionSet getSet() {
        return this.set;
    }
}
