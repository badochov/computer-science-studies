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

vector<int> create_suffix(const vector<int> &highway) {
    vector<int> suf(highway.size(), 0);
    suf[highway.size() - 1] = highway.back();
    for (int i = (int) highway.size() - 2; i >= 0; i--) {
        suf[i] = suf[i + 1] + highway[i];
    }

    return suf;
}

int get_ans(const vector<int> &highway) {
    auto suf = create_suffix(highway);
    if (suf.front() < 0) {
        return -1;
    }
}

int main() {
    setupIO();
    int n;
    cin >> n;
    vector<int> highway;
    highway.reserve(n);
    FOR(0, n) {
        highway.push_back(n);
    }

    cout << get_ans(highway);

    return 0;
}