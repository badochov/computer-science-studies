

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

constexpr int MOD = 1e9 + 7;

int main() {
    setupIO();
    int n = input_unsigned_number();
    vector<int> cuksy;
    cuksy.reserve(n);
    FOR(0, n)cuksy.push_back(input_unsigned_number());
    sort(cuksy.begin(), cuksy.end());

    map<int, int, std::greater<>> na_ile;
    na_ile[0] = 1;

    int max_possible = 0;

    for (int c : cuksy) {
        if (c - max_possible > 1) {
            break;
        }
        max_possible = max_possible + c;
        for (const auto[key, count] : na_ile) {
            if (c - key > 1) {
                break;
            }
            int new_key = key + c;
            if(new_key > 5000){
                new_key = 5001;
            }
            na_ile[new_key] += count;
            na_ile[new_key] %= MOD;
        }
    }


    int sum = -1;
    for (const auto[_key, count] : na_ile) {
        sum += count;
        sum %= MOD;
    }

    cout << sum;
    return 0;
}