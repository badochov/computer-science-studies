package cover;

import java.util.ArrayList;
import java.util.Collections;

public class Query {
    private int n;
    private int mode;


    public void query(ArrayList<SolutionSet> sets) {
        QuerySet qs = new QuerySet(this.n);

        ArrayList<Integer> solution = new ArrayList<>();
        switch (this.mode) {
            case 1:
                solution = new BruteForceSolution(sets, qs).solve();
                break;
            case 2:
                solution = new GreedyHeuristicSolution(sets, qs).solve();
                break;
            case 3:
                solution = new NaiveHeuristicSolution(sets, qs).solve();
                break;
        }
//        System.out.println(sets);

        printSolution(solution);

    }

    private void printSolution(ArrayList<Integer> solution) {
        String res = "";
        Collections.sort(solution);
        for (Integer i : solution) {
            res = res.concat(i.toString().concat(" "));
        }
        res = res.substring(0, res.length() - 1);

        System.out.println(res);
    }

    Query(int n, int mode) {
        this.n = n;
        this.mode = mode;
    }
}
