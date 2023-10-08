package cover;

import java.util.HashSet;
import java.util.Set;

public class QuerySet {
    private Set<Integer> els;
    private int n;

    QuerySet(int n) {
        this.n = n;

        this.els = new HashSet<>();
        for (int i = 1; i <= n; i++) {
            this.els.add(i);
        }
    }

    private QuerySet(int n, Set<Integer> els) {
        this.n = n;
        this.els = els;
    }

    QuerySet subtract(int n) {
        this.els.remove(n);
        return this;
    }

    public boolean isEmpty() {
        return this.els.size() == 0;
    }

    public int maxPossibleEl() {
        return this.n;
    }

    public int size() {
        return this.els.size();
    }

    public boolean contains(int el) {
        return this.els.contains(el);
    }

    public QuerySet copy() {
        return new QuerySet(this.n, new HashSet<>(this.els));
    }

}
