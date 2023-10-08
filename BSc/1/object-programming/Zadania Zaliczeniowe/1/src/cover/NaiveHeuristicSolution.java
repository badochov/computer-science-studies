package cover;

import java.util.ArrayList;

public class NaiveHeuristicSolution implements Solution {
    private ArrayList<SolutionSet> sets;
    private QuerySet qs;

    @Override
    public ArrayList<Integer> solve() {
        ArrayList<Integer> sol = new ArrayList<>();
        for (int i = 0; i < this.sets.size(); i++) {
            int prevElsCount = this.qs.size();
            this.sets.get(i).subtractFromQuerySet(this.qs);
            if (this.qs.size() < prevElsCount) {
                sol.add(i + 1);
            }
        }
        if (this.qs.isEmpty()) {
            return sol;
        }
        sol.clear();
        sol.add(0);

        return sol;
    }

    NaiveHeuristicSolution(ArrayList<SolutionSet> sets, QuerySet qs) {
        this.sets = sets;
        this.qs = qs;
    }
}
