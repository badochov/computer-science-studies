import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

public class Polynomial {
    private Map<Integer, Integer> coefficients;

    public Polynomial(Integer num) {
        this(new Integer[]{num});
    }

    public Polynomial(Integer[] nums) {
        this.coefficients = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != 0) {
                this.coefficients.put(i, nums[i]);
            }
        }
    }

    public Polynomial(Polynomial p) {
        this.coefficients = new HashMap<>(p.coefficients);
    }

    private static int raiseToPower(int x, int pow) {
        if (pow == 0) {
            return 1;
        }
        if (pow % 2 == 0) {
            return Polynomial.raiseToPower(x * x, pow / 2);
        } else {
            return x * Polynomial.raiseToPower(x, pow - 1);
        }
    }

    public int valueIn(int x) {
        int res = 0;
        for (Map.Entry<Integer, Integer> coefficient : this.coefficients.entrySet()) {
            res += Polynomial.raiseToPower(x, coefficient.getKey()) * coefficient.getValue();
        }

        return res;
    }

    private int getCoefficient(int n) {
        return this.coefficients.getOrDefault(n, 0);
    }

    public Polynomial add(Polynomial p) {
        p.coefficients.forEach((deg, coefficient) -> {
            int new_coefficient = this.getCoefficient(deg) + coefficient;
            if (new_coefficient == 0) {
                this.coefficients.remove(deg);
            } else {
                this.coefficients.put(deg, new_coefficient);
            }
        });

        return this;
    }

    public Polynomial multiply(Polynomial p) {
        HashMap<Integer, Integer> new_coefficients = new HashMap<>();
        p.coefficients.forEach((deg, coefficient) -> {
            this.coefficients.forEach((deg2, coefficient2) -> {
                int new_deg = deg + deg2;
                int new_coefficient = new_coefficients.getOrDefault(new_deg, 0) + coefficient * coefficient2;
                new_coefficients.put(new_deg, new_coefficient);
            });
        });

        this.coefficients = new_coefficients;
        return this;
    }

    public void print() {
        String s = "";
        if (this.coefficients.size() == 0) {
            s = "0";
        } else {
            ArrayList<String> coefficients = new ArrayList<>(this.coefficients.size());
            new TreeMap<>(this.coefficients).forEach((deg, value) -> {
                if (deg == 0) {
                    coefficients.add(value.toString());
                } else if (deg == 1) {
                    coefficients.add(value.toString().concat("x"));
                } else {
                    coefficients.add(value.toString().concat("x^").concat(deg.toString()));
                }
            });
            s = String.join(" + ", coefficients);
        }
        System.out.println(s);
    }
}
