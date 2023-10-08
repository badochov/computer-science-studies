#include <bits/stdc++.h>

#define FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name < end; var_name++)
#define FOR_HELPER_1(start, end) FOR_HELPER(start, end, i)
#define FOR_CHOOSER(x, start, end, increment, FOR_MACRO, ...) FOR_MACRO

#define FOR(...) FOR_CHOOSER(,##__VA_ARGS__, FOR_HELPER(__VA_ARGS__), FOR_HELPER_1(__VA_ARGS__))
#define FOR2(start, end) FOR_HELPER(start, end, j)

#define REV_FOR_HELPER(start, end, var_name) for(auto var_name = start; var_name >= end; var_name--)
#define REV_FOR_HELPER_1(start, end) FOR_HELPER(start, end, i)
#define REV_FOR_HELPER_2(start) FOR_HELPER(start, 0, i)

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

int get_min_amount(const vector<int> &prog, const vector<pair<int, int>> &dep, int k) {
    vector<pair<int, int>> g;
    g.reserve(prog.size());
    vector<vector<int>> out(prog.size(), vector<int>());
    for (int p : prog)
        g.emplace_back(0, p);
    for (const auto[from, to] : dep) {
        g[to - 1].first++;
        out[from -1].push_back(to-1);
    }

    int min_prog_needed = 0;
    priority_queue<pair<int,int>, vector<pair<int,int>>, greater<>> pq;
    FOR(0, g.size()){
        const auto [count, prg] = g[i];
        if(count == 0)
            pq.emplace(prg, i);
    }

    FOR(0, k) {
        auto v = pq.top();
        pq.pop();
        for(int to : out[v.second]){
            auto &p = g[to];
            p.first --;
            if(p.first == 0){
                pq.emplace(p.second, to);
            }
        }
        min_prog_needed = max(min_prog_needed, v.first);

    }

    return min_prog_needed;
}

int main() {
    setupIO();
    int n = input_unsigned_number(), m = input_unsigned_number(), k = input_unsigned_number();
    vector<int> prog;
    prog.reserve(n);
    vector<pair<int, int>> dep;
    dep.reserve(m);
    FOR(0, n)prog.push_back(input_unsigned_number());
    FOR(0, m)dep.emplace_back(input_unsigned_number(), input_unsigned_number());

    cout << get_min_amount(prog, dep, k);


    return 0;
}