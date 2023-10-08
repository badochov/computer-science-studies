package cover;

public class Infinite implements SolutionSetElement {
    private Integer a;
    private Integer r;

    @Override
    public QuerySet subtractFromQuerySet(QuerySet qs) {
        for (int i = 0; this.nthElement(i) <= qs.maxPossibleEl(); i++) {
            qs.subtract(this.nthElement(i));
        }
        return qs;
    }

    Infinite(int a, int r) {
        this.a = a;
        this.r = r;
    }

    public int nthElement(int n) {
        return this.a + this.r * n;
    }

    @Override
    public String toString() {
        return this.a.toString().concat(" ").concat(this.r.toString());
    }
}
