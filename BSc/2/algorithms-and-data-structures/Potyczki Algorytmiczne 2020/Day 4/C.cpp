

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

vector<pair<int, int>> get_neighbours(const pair<int, int> &coords, vector<vector<bool>> &path) {
    static pair<int, int> n[] = {{-1, 0},
                                 {1,  0},
                                 {0,  1},
                                 {0,  -1}};
    int max_x = path.size();
    int max_y = path[0].size();

    vector<pair<int, int>> res;
    for (const auto &[dx, dy] : n) {
        int new_x = coords.first + dx;
        int new_y = coords.second + dy;

        if (new_x >= 0 && new_x < max_x && new_y >= 0 && new_y < max_y) {
            if (path[new_x][new_y]) {
                path[new_x][new_y] = false;
                res.emplace_back(new_x, new_y);
            }
        }
    }

    return res;
}

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}

int maxW = 1e6 + 7;

int min_path(vector<vector<bool>> &path) {
    queue<pair<int, pair<int, int>>> q;
    q.emplace(0, make_pair(0, 0));

    pair<int, int> end = {path.size() - 1, path[0].size() - 1};

    while (!q.empty()) {
        auto const &[len, coords] = q.front();
        q.pop();
        if (coords == end) {
            return len;
        }
        for (const auto &neighbour : get_neighbours(coords, path)) {
            q.emplace(len + 1, neighbour);
        }
    }

    return -1;
}

int main() {
    setupIO();
    int n, m, k;
    cin >> n >> m >> k;
    vector<vector<bool>> podejscie(n, vector<bool>(m, false));


    FOR(0, n) {
        FOR2(0, m) {
            char c;
            cin >> c;
            podejscie[i][j] = c == '.';
        }
    }

    vector<pair<int, int>> explorers;
    explorers.reserve(k);
    FOR(0, k) {
        int up, down;
        cin >> up >> down;
        explorers.emplace_back(up, down);
    }

    int res = min_path(podejscie);
    int theoretical_min = n - 1 + m - 1;
    int up = theoretical_min + (res - theoretical_min) / 2;
    int down = res - up;
    multiset<long long> len;

    for (const auto &[up_speed, down_speed] : explorers) {
        len.insert((long long)up * up_speed + (long long)down * down_speed);
    }

    long long shortest = *len.begin();
    cout << shortest << " " << len.count(shortest);

    return 0;
}