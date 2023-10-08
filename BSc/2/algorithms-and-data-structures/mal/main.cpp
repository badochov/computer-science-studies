#include <bits/stdc++.h>

#define FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name < end; var_name++)
#define FOR_HELPER_1(start, end) FOR_HELPER(start, end, i)
#define FOR_CHOOSER(x, start, end, increment, FOR_MACRO, ...) FOR_MACRO

#define FOR(...) FOR_CHOOSER(,##__VA_ARGS__, FOR_HELPER(__VA_ARGS__), FOR_HELPER_1(__VA_ARGS__))
#define FOR2(start, end) FOR_HELPER(start, end, j)

#define REV_FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name >= end; var_name--)
#define REV_FOR_HELPER_1(start, end) REV_FOR_HELPER(start, end, i)
#define REV_FOR_HELPER_2(start) REV_FOR_HELPER(start, 0, i)

#define REV_FOR_CHOOSER(x, start, end, increment, FOR_MACRO, ...) FOR_MACRO

#define REV_FOR(...) REV_FOR_CHOOSER(,##__VA_ARGS__, REV_FOR_HELPER(__VA_ARGS__), REV_FOR_HELPER_1(__VA_ARGS__), REV_FOR_HELPER_2(__VA_ARGS__))
#define REV_FOR2(start, end) REV_FOR_HELPER(start, end, j)

#define END_LINE '\n'

using namespace std;

template<class t = int>
[[maybe_unused]] t input_unsigned_number() {
    t num = 0;
    int c = getchar_unlocked();
    while (c < '0') c = getchar_unlocked();
    while (c >= '0') {
        num = num * 10 + c - '0';
        c = getchar_unlocked();
    }
    return num;
}

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}

int tree_size(int n) {
    int size = 1;
    while (size < n) {
        size *= 2;
    }
    return size * 2;
}

vector<int> init_interval_tree(int n) {
    return vector<int>(tree_size(n));
}

int set_value(vector<int> &tree, int s, int e, int a, int b, int v, int idx, int overwritten_with) {
    int size = (b - a + 1);
    if (overwritten_with != -1) {
        tree[idx] = size * overwritten_with;
    }else{
        if(tree[idx] == size){
            overwritten_with = 1;
        }
        else if(tree[idx] == 0){
            overwritten_with = 0;
        }
    }

    if(tree[idx] == 0 && v == 0){
        return 0;
    }

    if (e < a || b < s) {
        return tree[idx];
    }
    if (b <= e && a >= s) {
        tree[idx] = size*v;
        return tree[idx];
    }
    int mid = (a + b) / 2;
    tree[idx] = set_value(tree, s, e, a, mid, v, idx * 2, overwritten_with) +
           set_value(tree, s, e, mid + 1, b, v, idx * 2 + 1, overwritten_with);


    return tree[idx];
}

void color(vector<int> &tree, int start, int end, char c) {
    int v = c == 'B';
    int b = tree.size()/2;

    set_value(tree, start, end, 1, b, v, 1, -1);
}

int main() {
    setupIO();
    const int n = input_unsigned_number(), m = input_unsigned_number();
    vector<int> p;
    p.reserve(n);
    vector<int> tree = init_interval_tree(n);
    FOR(0, m) {
        int s = input_unsigned_number(), e = input_unsigned_number();
        char c = getchar_unlocked();
        color(tree, s, e, c);
        cout << tree[1] <<endl;
    }
    return 0;
}