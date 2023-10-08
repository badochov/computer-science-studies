package cover;

import java.util.ArrayList;

public class BruteForceSolution implements Solution {
    private ArrayList<SolutionSet> sets;
    private QuerySet qs;

    @Override
    public ArrayList<Integer> solve() {
        for (int i = 1; i <= this.sets.size(); i++) {
            ArrayList<Integer> res = solve(i, 0, new ArrayList<>());
            if (res != null) {
                return res;
            }
        }

        ArrayList<Integer> res = new ArrayList<>();
        res.add(0);

        return res;
    }

    private ArrayList<Integer> solve(int max_size, int num, ArrayList<Integer> ans) {
        if (ans.size() == max_size) {
            QuerySet qsCopy = this.qs.copy();
            for (int i = 0; i < max_size; i++) {
                this.sets.get(ans.get(i) - 1).subtractFromQuerySet(qsCopy);
            }
            if (qsCopy.isEmpty()) {
                return ans;
            }
        } else {
            for (int i = num; i < this.sets.size(); i++) {
                ans.add(i + 1);
                ArrayList<Integer> res = solve(max_size, i + 1, ans);
                if (res != null) {
                    return res;
                }
                ans.remove(ans.size() - 1);
            }
        }

        return null;
    }

    BruteForceSolution(ArrayList<SolutionSet> sets, QuerySet qs) {
        this.sets = sets;
        this.qs = qs;
    }
}
