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

using graph_t = map<int, vector<pair<int, int>>>;



using best_t = pair<unsigned long long, vector<int>>;
using cont = tuple<unsigned long long, int, vector<int>>;

vector<int> get_dist(const graph_t &g, int n) {
    vector<best_t> best(n + 1, {1e18, {}});
    vector<int> in_q(n + 1, 0);
    auto l = [](const cont &a, const cont &b) { return get<0>(a) > get<0>(b); };
    priority_queue<cont, vector<cont>, decltype(l)> pq(l);
    pq.emplace(0, 1, 0);
    in_q[1] = 1;
    while (!pq.empty()) {
        auto[dist, point, path] = pq.top();
        path.emplace_back(point);
        pq.pop();
        best[point] = min(best[point], {dist, path});
        in_q[point]--;
        if (point == n) {
            continue;
        }
        if (in_q[point] == 0) {
            for (auto &[p, d] : g.at(point)) {
                unsigned long long new_dist = dist + d;
                if (best[p].first >= new_dist) {
                    pq.emplace(new_dist, p, best[point].second);
                    in_q[p]++;
                }
            }
        }
    }

    return best[n].second;
}

int main() {
    setupIO();
    int n, m;
    cin >> n >> m;
    graph_t g;

    FOR(0, m) {
        int a, b, d;
        cin >> a >> b >> d;
        g[a].emplace_back(b, d);
        g[b].emplace_back(a, d);
    }

    auto res = get_dist(g, n);
    for (auto el : res) {
        cout << el << " ";
    }

    return 0;
}