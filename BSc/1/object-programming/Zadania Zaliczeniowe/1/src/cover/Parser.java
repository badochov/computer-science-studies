package cover;

import java.util.ArrayList;
import java.util.Scanner;

public class Parser {
    Scanner sc;
    ArrayList<SolutionSet> sets;
    Integer next;

    Parser() {
        this.sc = new Scanner(System.in);
        this.sets = new ArrayList<>();
        this.next = this.sc.hasNextInt() ? this.sc.nextInt() : null;
    }

    public void parse() {
        while (this.hasNext()) {
            Data d = this.parseData();
            if (d.getQuery() != null) {
                d.getQuery().query(this.sets);
            } else {
                this.sets.add(d.getSet());
            }
        }
    }

    private int nextInt() {
        int c = this.next;
        this.next = this.sc.hasNextInt() ? this.sc.nextInt() : null;
        return c;
    }

    private int peekNextInt() {
        return this.next;
    }

    private boolean hasNext() {
        return this.next != null;
    }

    private Data parseData() {
        Query q = null;
        SolutionSet s = null;
        if (this.peekNextInt() < 0) {
            q = this.parseQuery();
        } else {
            s = this.parseSet();
        }

        return new Data(q, s);
    }

    private SolutionSet parseSet() {
        SolutionSet s = new SolutionSet();
        while (true) {
            SolutionSetElement el = this.parseSetElement();
            if (el == null) {
                return s;
            }
            s.addElement(el);
        }
    }

    private SolutionSetElement parseSetElement() {
        int num = this.nextInt();
        if (num == 0) {
            return null;
        }
        if (this.peekNextInt() < 0) {
            int r = -this.nextInt();
            if (this.peekNextInt() < 0) {
                return new Finite(num, r, -this.nextInt());
            }
            return new Infinite(num, r);
        }
        return new Element(num);
    }

    private Query parseQuery() {
        int n = -this.nextInt();
        int mode = this.nextInt();

        return new Query(n, mode);
    }
}
