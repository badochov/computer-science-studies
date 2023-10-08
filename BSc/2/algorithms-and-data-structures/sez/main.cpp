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
#define REV_FOR2(start, end) REV_HELPER(start, end, j)

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

template<typename T>
int get_tree_size(const vector<T> &graph) {
    int s = 1;
    while (s < graph.size()) {
        s *= 2;
    }
    return s * 2;
}

vector<int> operator+(const vector<int> &v1, const vector<int> &v2) {
    if (v2.empty()) {
        return vector<int>(v1);
    } else {
        vector<int> res;
        res.reserve(v1.size());
        FOR(0, v1.size()) {
            res[i] = v1[i] + v2[i];
        }

        return res;
    }

}

int max_o(int a, int b) {
    return max(a, b);
}

vector<int> max_o(const vector<int> &v1, const vector<int> &v2) {
    if (v2.empty()) {
        return vector<int>(v1);
    } else {
        vector<int> res;
        res.reserve(v1.size());
        FOR(0, v1.size()) {
            res.push_back(max(v1[i], v2[i]));
        }

        return res;
    }
}

int min_o(int a, int b) {

    if (a == 0) {
        return b;
    }
    if (b == 0) {
        return a;
    }
    return min(a, b);
}

vector<int> min_o(const vector<int> &v1, const vector<int> &v2) {
    if (v2.empty()) {
        return vector<int>(v1);
    } else {
        vector<int> res;
        res.reserve(v1.size());
        FOR(0, v1.size()) {
            res.push_back(min(v1[i], v2[i]));
        }

        return res;
    }
}

template<typename T>
vector<T> create_tree_max(const vector<T> &v) {
    int size = get_tree_size(v);

    vector<T> tree(size);
    FOR(0, v.size()) {
        tree[size / 2 + i] = v[i];
    }

    REV_FOR(size / 2 - 1, 1) {
        tree[i] = max_o(tree[i * 2], tree[i * 2 + 1]);
    }

    return tree;
}

template<typename T>
vector<T> create_tree_min(const vector<T> &v) {
    int size = get_tree_size(v);

    vector<T> tree(size);

    FOR(0, v.size()) {
        tree[size / 2 + i] = v[i];
    }

    REV_FOR(size / 2 - 1, 1) {
        tree[i] = min_o(tree[i * 2], tree[i * 2 + 1]);
    }

    return tree;
}

vector<vector<int>> create_tree_helper_max(const vector<vector<int>> &graph) {
    vector<vector<int>> horizontal;
    horizontal.reserve(graph.size());
    for (auto &row: graph) {
        horizontal.push_back(create_tree_max(row));
    }

    return create_tree_max(horizontal);
}

vector<vector<int>> create_tree_helper_min(const vector<vector<int>> &graph) {
    vector<vector<int>> horizontal;
    horizontal.reserve(graph.size());
    for (auto &row: graph) {
        horizontal.push_back(create_tree_min(row));
    }

    return create_tree_min(horizontal);
}

long long get_val_max(const vector<int> &tree, int pos, int start, int end, int y1, int y2) {
    if (y2 < start || end < y1) {
        return 0;
    }
    if (y1 <= start && end <= y2) {
        return tree[pos];
    }

    int mid = (start + end) / 2;
    long long l = get_val_max(tree, 2 * pos, start, mid, y1, y2);

    long long p = get_val_max(tree, 2 * pos + 1, mid + 1, end, y1, y2);

    return max(l, p);
}


long long get_val_max(const vector<vector<int>> &tree, int pos, int start, int end, int x1, int x2, int y1, int y2) {
    if (x2 < start || end < x1) {
        return 0;
    }
    if (x1 <= start && end <= x2) {
        return get_val_max(tree[pos], 1, 1, tree[1].size() / 2, y1, y2);
    }

    int mid = (start + end) / 2;
    long long l = get_val_max(tree, 2 * pos, start, mid, x1, x2, y1, y2);

    long long p = get_val_max(tree, 2 * pos + 1, mid + 1, end, x1, x2, y1, y2);

    return max(l, p);
}


long long get_val_min(const vector<int> &tree, int pos, int start, int end, int y1, int y2) {
    if (y2 < start || end < y1) {
        return 0;
    }
    if (y1 <= start && end <= y2) {
        return tree[pos];
    }

    int mid = (start + end) / 2;
    long long l = get_val_min(tree, 2 * pos, start, mid, y1, y2);

    long long p = get_val_min(tree, 2 * pos + 1, mid + 1, end, y1, y2);

    return min_o(l, p);
}


long long get_val_min(const vector<vector<int>> &tree, int pos, int start, int end, int x1, int x2, int y1, int y2) {
    if (x2 < start || end < x1) {
        return 0;
    }
    if (x1 <= start && end <= x2) {
        return get_val_min(tree[pos], 1, 1, tree[1].size() / 2, y1, y2);
    }

    int mid = (start + end) / 2;
    long long l = get_val_min(tree, 2 * pos, start, mid, x1, x2, y1, y2);

    long long p = get_val_min(tree, 2 * pos + 1, mid + 1, end, x1, x2, y1, y2);

    return min_o(l, p);
}

int m;

int main() {
    setupIO();
    int n = input_unsigned_number();
    m = input_unsigned_number();
    int k = input_unsigned_number();
    vector<vector<int>> graph(n, vector<int>(m, 0));

    FOR(0, n) {
        FOR2(0, m) {
            graph[i][j] = input_unsigned_number();
        }
    }

    vector<vector<int>> max_tree = create_tree_helper_max(graph);
    vector<vector<int>> min_tree = create_tree_helper_min(graph);


    FOR(0, k) {
        int x1 = input_unsigned_number() + 1;
        int y1 = input_unsigned_number() + 1;
        int x2 = input_unsigned_number() + 1;
        int y2 = input_unsigned_number() + 1;

        int max_v = get_val_max(max_tree, 1, 1, max_tree.size() / 2, x1, x2, y1, y2);
        int min_v = get_val_min(min_tree, 1, 1, min_tree.size() / 2, x1, x2, y1, y2);
        cout << max_v - min_v << endl;
    }


    return 0;
}