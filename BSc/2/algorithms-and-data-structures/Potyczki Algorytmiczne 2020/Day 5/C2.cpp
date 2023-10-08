

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

int maxW = 1e6 + 7;

int main() {
    setupIO();
    unsigned long long n = input_unsigned_number();
    cout << 91 << endl;
    FOR(1, 30) {
        int second = (n & (1ull << (i - 1))) ? i + 30 : -1;

        cout << i + 1 << " " << second << endl;
    }
    int second = (n & (1ull << (30 - 1))) ? 30 + 30 : -1;
    // 30
    cout << -1 << " " << second << endl;

    // 31
    cout << -1 << " " << 61 << endl;
    FOR(30 + 2, 30 + 31) {
        cout << i - 1 << " " << i + 30 << endl;
    }

    // 31
    cout << -1 << " " << 91 << endl;
    FOR(60 + 2, 60 + 31) {
        cout << -1 << " " << i - 30 - 1 << endl;
    }

    // 91
    cout << -1 << " " << -1 << endl;

    return 0;
}