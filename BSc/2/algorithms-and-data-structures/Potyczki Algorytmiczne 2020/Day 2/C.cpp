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

int main() {
    setupIO();
    int n;
    cin >>n;
    vector<vector<char>> bitus(2, vector<char>());
    vector<vector<char>> bajtus(2, vector<char>());
    bitus[0].reserve(ceil(n / 2));
    bajtus[0].reserve(ceil(n / 2));
    bitus[1].reserve(floor(n / 2));
    bajtus[1].reserve(floor(n / 2));
    string a;
    string b;
    a.reserve(n);
    b.reserve(n);
    cin >> a;
    cin >> b;


    FOR(0, n) {
        bitus[i % 2].push_back(a[i]);
    }
    FOR(0, n) {
        bajtus[i % 2].push_back(b[i]);
    }
    FOR(0, 2) {
        sort(bajtus[i].begin(), bajtus[i].end());
        sort(bitus[i].begin(), bitus[i].end());
    }

    cout << (bajtus == bitus ? "TAK" : "NIE");

    return 0;
}