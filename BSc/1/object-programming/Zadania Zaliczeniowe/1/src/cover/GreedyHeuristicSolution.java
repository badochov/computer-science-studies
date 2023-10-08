package cover;

import java.util.ArrayList;

public class GreedyHeuristicSolution implements Solution {
    private ArrayList<SolutionSet> sets;
    private QuerySet qs;

    @Override
    public ArrayList<Integer> solve() {
        ArrayList<Integer> sol = new ArrayList<>();
        boolean used[] = new boolean[this.sets.size()];
        while (true) {
            int max_size = 0;
            int max_index = 0;
            for (int i = 0; i < this.sets.size(); i++) {
                if (!used[i]) {
                    int curr_size = this.sets.get(i).intersectionSizeWithQuerySet(this.qs);
                    if (curr_size > max_size) {
                        max_index = i;
                        max_size = curr_size;
                    }
                }
            }

            if (max_size == 0) {
                break;
            }
            this.sets.get(max_index).subtractFromQuerySet(this.qs);
            sol.add(max_index + 1);
            used[max_index] = true;
        }
        if (this.qs.isEmpty()) {
            return sol;
        }
        sol.clear();
        sol.add(0);

        return sol;
    }

    GreedyHeuristicSolution(ArrayList<SolutionSet> sets, QuerySet qs) {
        this.sets = sets;
        this.qs = qs;
    }
}
