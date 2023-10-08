package cw3.dodatkowe;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Collections;

class Median {

    private static int slowMedian(List<Integer> l) {
        Collections.sort(l);
        return l.get(l.size() / 2);
    }

    private static int pickPivot(List<Integer> l) {
        if (l.size() < 5) {
            return slowMedian(l);
        }
        List<Integer> medians = new ArrayList<>();
        ArrayList<Integer> chunk = new ArrayList<>();
        int i = 0;
        for (Integer num : l) {
            if (i >= l.size() / 5 * 5) {
                break;
            }
            chunk.add(num);
            if (i % 5 == 4) {
                medians.add(slowMedian(chunk));
                chunk = new ArrayList<>();
            }
            i++;
        }
        int res[] = new int[medians.size()];
        i = 0;
        for (Integer med : medians) {
            res[i] = med.intValue();
            i++;
        }
        return median(res);
    }

    // may be logarithmic not linear -> magic 5 to improve\private static int
    // nth(int arr[], int n) {
    private static int nth(List<Integer> a, int n) {
        if (a.size() == 1) {
            if (n == 0) {
                return a.get(0);
            }
            return -1;
        }
        int pivot = pickPivot(a);
        List<Integer> smaller = new ArrayList<>();
        List<Integer> bigger = new ArrayList<>();
        List<Integer> equal = new ArrayList<>();
        for (Integer num : a) {
            if (num < pivot) {
                smaller.add(num);
            } else if (num == pivot) {
                equal.add(num);
            } else {
                bigger.add(num);
            }
        }

        if (n < smaller.size()) {
            return nth(smaller, n);
        }
        if (n < smaller.size() + equal.size()) {
            return pivot;
        }
        return nth(bigger, n - smaller.size() - equal.size());
    }

    private static int nth(int arr[], int n) {
        Integer a[] = Arrays.stream(arr).boxed().toArray(Integer[]::new);
        return nth(Arrays.asList(a), n);
    }

    // return -1 if median is not decimal
    private static int median(int arr[]) {
        if (arr.length % 2 == 1) {
            return nth(arr, arr.length / 2);
        } else {
            int a = nth(arr, arr.length / 2 - 1);
            int b = nth(arr, arr.length / 2);
            if (a == b) {
                return a;
            }
            return -1;
        }
    }
}