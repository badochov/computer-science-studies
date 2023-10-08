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


//constexpr unsigned long long Q = 1125899839733759;
constexpr unsigned long long Q = 1e9 + 7;
constexpr unsigned long long P = 29;


constexpr int POWERS = 300000 + 7;

array<unsigned long long, POWERS> calc_powers() {
    array<unsigned long long, POWERS> res{};
    res[0] = 1;
    FOR(1, POWERS) {
        res[i] = (res[i - 1] * P) % Q;
    }

    return res;
}

const array<unsigned long long, POWERS> P_POWER = calc_powers();


int v(char c) {
    return c - 'a';
}

vector<unsigned long long> calc_prefix_hashes(const string &w) {
    vector<unsigned long long> hashes;
    hashes.reserve(w.size());
    hashes.push_back(v(w[0]));
    FOR(1, w.size()) {
        hashes.push_back((hashes.back() + v(w[i]) * P_POWER[i]) % Q);
    }

    return hashes;
}

vector<unsigned long long> prefix_hashes;
string w;

unsigned long long get_hash(int s, int e) {
    auto prev_hash = s > 1 ? prefix_hashes[s - 2] : 0;

    return (prefix_hashes[e - 1] + Q - prev_hash) % Q;
}

long long mod_mul(long long a, long long b, long long m) {
    long long q = (long double) a * (long double) b / (long double) m;
    long long p = a * b - m * q;
    return (p + 5 * m) % m;
}

//unsigned long long mod_mul(unsigned long long a, unsigned long long b, unsigned long long mod) {
//    unsigned long long res = 0;
//    int counter = 0;
//    while (b) {
//        counter++;
//        if (counter > 65) {
//            exit(1337);
//        }
//        if (b % 2 == 1) {
//            res = (res + a) % mod;
//        }
//        a = (2 * a) % mod;
//        b /= 2;
//    }
//
//    return res;
//}

bool same(int s1, int e1, int s2, int e2) {
    auto hash1 = get_hash(s1, e1);
    auto hash2 = get_hash(s2, e2);

    if (s1 > s2) {
        hash2 = mod_mul(P_POWER[s1 - s2], hash2, Q);
    } else {
        hash1 = mod_mul(P_POWER[s2 - s1], hash1, Q);
    }

    return hash2 == hash1;
}

int get_common_len(int s1, int e1, int s2, int e2) {
    int l1 = e1 - s1;
    int l2 = e2 - s2;


    int s = 0;
    int e = min(l1, l2);
    while (s <= e) {

        int m = (s + e) / 2;
        bool cmp = same(s1, s1 + m, s2, s2 + m);
        if (cmp) {
            s = m + 1;
        } else {
            e = m - 1;
        }
    }
    return e;
}

char ans(int s1, int e1, int s2, int e2) {
    int common_len = get_common_len(s1, e1, s2, e2);
    int l1 = e1 - s1;
    int l2 = e2 - s2;
    if (common_len == l1 && common_len == l2) {
        return '=';
    }
    if (common_len == l1) {
        return '<';
    }
    if (common_len == l2) {
        return '>';
    }
    int cmp = w[s1 + common_len] - w[s2 + common_len];
    if (cmp > 0) {
        return '>';
    }
    return '<';

}

int main() {
    setupIO();

    int n, m;
    cin >> n >> m;
    w.reserve(n);
    cin >> w;
    prefix_hashes = calc_prefix_hashes(w);

    FOR(0, m) {
        int s1, e1, s2, e2;
        cin >> s1 >> e1 >> s2 >> e2;
        cout << ans(s1, e1, s2, e2) << endl;
    }

    return 0;
}