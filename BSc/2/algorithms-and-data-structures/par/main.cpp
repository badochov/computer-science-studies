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

int LOG = 20;

struct node {
    int right = -1;
    int left = -1;
    int depth = 0;
    int parent[20] = {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1};
};

void preorder(vector<int> &pre, vector<node> &g, int v) {
    if (g[v].left != -1) {
        preorder(pre, g, g[v].left);
    }
    if (g[v].right != -1) {
        preorder(pre, g, g[v].right);
    }


    pre.push_back(v);
}

void add_parents(vector<node> &g, int parent) {
    vector<int> pre;
    preorder(pre, g, parent);

    FOR(1, LOG) {
        for (int v :pre) {
            if (g[v].parent[i - 1] != -1) {
                g[v].parent[i] = g[g[v].parent[i - 1]].parent[i - 1];
            }
        }
    }
}

int find_deepest(int v, vector<node> &g, int lvl) {
    g[v].depth = lvl;
    int l = g[v].left;
    int r = g[v].right;

    int l_deepest = l == -1 ? 0 : find_deepest(l, g, lvl + 1);
    int r_deepest = r == -1 ? 0 : find_deepest(r, g, lvl + 1);
    if (l_deepest == 0) {
        if (r_deepest == 0) {
            return v;
        }

        return r_deepest;
    }
    if (r_deepest == 0) {
        return l_deepest;
    }
    if (g[l_deepest].depth > g[r_deepest].depth) {
        return l_deepest;
    }
    return r_deepest;
}

pair<int, int> find_longest(vector<node> &g, int pos, vector<bool> &v) {
    if (pos == -1) {
        return {0, 0};
    }
    if (v[pos]) {
        return {0, 0};
    }

    v[pos] = true;
    auto res = max(find_longest(g, g[pos].parent[0], v),
                   max(find_longest(g, g[pos].left, v), find_longest(g, g[pos].right, v)));
    res.first++;
    if (res.second == 0) {
        res.second = pos;
    }
    return res;
}

pair<int, int> find_longest_path(vector<node> &g, int parent) {
    int deepest = find_deepest(parent, g, 0);

    vector<bool> v(g.size(), false);
    auto[_, longest] = find_longest(g, deepest, v);

    return {deepest, longest};
}

int lca(const vector<node> &g, int a, int b) {
    if (g[a].depth < g[b].depth) {
        return lca(g, b, a);
    }

    int d = g[a].depth - g[b].depth;

    FOR(0, LOG) {
        if (d & (1 << i)) {
            a = g[a].parent[i];
        }
    }

    if (a == b) {
        return a;
    }

    for (int i = LOG - 1; i >= 0; --i) {

        if (g[a].parent[i] != g[b].parent[i] && g[b].parent[i] != -1) {
            a = g[a].parent[i];
            b = g[b].parent[i];
        }
    }

    return g[b].parent[0];
}

int up(const vector<node> &g, int v, int s) {
    if (s == 0)
        return v;

    int pow = 1, idx = 0;

    while (2 * pow <= s) {
        pow *= 2;
        idx++;
    }

    return up(g, g[v].parent[idx], s - pow);
}

int main() {
    setupIO();
    int n;
    cin >> n;
    vector<node> graph(n + 1);
    set<int> parents;
    FOR(1, n + 1) {
        parents.insert(i);
    }
    FOR(1, n + 1) {
        int left, right;
        cin >> left >> right;
        if (left != -1) {
            graph[i].left = left;
            graph[left].parent[0] = i;
            parents.erase(left);
        }
        if (right != -1) {
            graph[i].right = right;
            graph[right].parent[0] = i;

            parents.erase(right);

        }
    }
    int parent = *parents.begin();

    auto[_l, r] = find_longest_path(graph, parent);

    add_parents(graph, parent);

    int m;
    cin >> m;
    FOR(0, m) {
        int v, d;
        cin >> v >> d;
        int l = lca(graph, v, _l);

        int ans = _l;

        if (graph[v].depth + graph[ans].depth - 2 * graph[l].depth < d) {
            l = lca(graph, v, r);
            ans = r;
            if (graph[v].depth + graph[ans].depth - 2 * graph[l].depth < d) {
                cout << -1 << endl;
                continue;
            }
        }

        if (graph[v].depth - graph[l].depth >= d) {
            cout << up(graph, v, d) << endl;
        } else {
            cout << up(graph, ans, graph[v].depth + graph[ans].depth - 2 * graph[l].depth - d) << endl;
        }
    }


    return 0;
}