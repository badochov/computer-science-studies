package cover;

import java.util.ArrayList;

public class SolutionSet implements SubtractableFromQuerySet {
    ArrayList<SubtractableFromQuerySet> elements;

    @Override
    public QuerySet subtractFromQuerySet(QuerySet qs) {
        for (SubtractableFromQuerySet el : this.elements) {
            el.subtractFromQuerySet(qs);
        }

        return qs;
    }

    public int intersectionSizeWithQuerySet(QuerySet qs) {
        QuerySet qsCopy = qs.copy();
        this.subtractFromQuerySet(qsCopy);

        return qs.size() - qsCopy.size();
    }

    SolutionSet() {
        this.elements = new ArrayList<>();
    }

    public SolutionSet addElement(SubtractableFromQuerySet el) {
        this.elements.add(el);

        return this;
    }

    @Override
    public String toString() {
        String s = "Set:\n";
        for (SubtractableFromQuerySet el : this.elements) {
            s = s.concat("\t").concat(el.toString()).concat("\n");
        }
        return s;
    }

}
