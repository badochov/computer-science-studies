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

constexpr int N = 1007;
constexpr int MOD = 1e9;

int memo[N][N][2];

void zero_mem(int n) {
    FOR(0, n)FOR2(0, n)memo[i][j][0] = memo[i][j][1] = 0;
}

void add_to_diag(int n) {
    FOR(0, n)memo[i][i][0] = memo[i][i][1] = 1;
}

void second_diag(int n, const int tab[N]) {
    FOR(0, n - 1) {
        bool res = tab[i] < tab[i + 1];
        memo[i][i + 1][0] = res;
        memo[i][i + 1][1] = res;
    }
}

void rest(int n, const int tab[N]) {
    FOR(2, n) {
        FOR2(0, n - i) {
            int h = i + j;
            memo[j][h][0] = (tab[j] < tab[j + 1]) * memo[j + 1][h][0] + (tab[j] < tab[h]) * memo[j + 1][h][1];
            memo[j][h][1] = (tab[j] < tab[h]) * memo[j][h - 1][0] + (tab[h-1] < tab[h]) * memo[j][h - 1][1];
            memo[j][h][0]%=MOD;
            memo[j][h][1]%=MOD;
        }
    }
}

int main() {
    setupIO();

    const int n = input_unsigned_number();
    int tab[N];
    FOR(0, n) tab[i] = input_unsigned_number();
    if (n == 1) {
        cout << 1;
        return 0;
    }
    zero_mem(n);
    add_to_diag(n);
    second_diag(n, tab);
    rest(n, tab);
//    FOR(0, n) {
//        FOR2(0, n) cout << memo[i][j][0];
//        cout << endl;
//    }
//
//    cout << endl;
//    FOR(0, n) {
//        FOR2(0, n) cout << memo[i][j][1];
//        cout << endl;
//    }
    cout << (memo[0][n - 1][0] + memo[0][n - 1][1]) % MOD;
    return 0;
}