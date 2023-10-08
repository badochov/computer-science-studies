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
    long long w, h, p;
    cin >> w >> h >> p;
    vector<long long> prev;
    vector<long long> curr(h + 1, 0);
    prev.reserve(h + 1);

    vector<long long> prev_ile;
    vector<long long> ile(h + 1, 0);
    prev_ile.reserve(h + 1);
    prev_ile.push_back(0);
    FOR(0, h + 1) {
        prev.push_back(i);
    }
    FOR(0, h) {
        if (i == 0) {
            prev_ile.push_back(h);
        } else {
            prev_ile.push_back((h - i + prev_ile[i] - i) % p);
        }
    }


    for (int i = 1; i < w; i++) {
        for (long long j = 1; j < h + 1; j++) {
            long long from_prev = j ? curr[j - 1] : 0;
            long long with_only_higher = prev[h - j + 1];
            long long single = prev_ile[j];
            long long higher_helper = ((j - 1) * with_only_higher) % p;
            long long v = from_prev + higher_helper + single;
            v %= p;
            curr[j] = v;
        }


        for (int j = h; j > 0; j--) {
            long long pr = j == h ? 0 : ile[j + 1];
            long long self = curr[j];
            long long pr_shorter = curr[h - j];
            ile[j] = (p + self + pr - pr_shorter) % p;
        }
        std::swap(prev, curr);
        std::swap(prev_ile, ile);
    }

    long long sum = 0;
    for (long long a :prev) {
        sum = (a + sum) % p;
    }
    cout << sum;
    return 0;
}