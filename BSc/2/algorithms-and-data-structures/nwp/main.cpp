#include <bits/stdc++.h>

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

// Complete the longestCommonSubsequence function below.
int longestCommonSubsequence(const vector<int> &a, const vector<int> &b, int c) {
    vector<vector<int>> dp(b.size(), vector<int>(20, 0));

    int m = 0;
    for (int el : a) {
        vector<int> carry(20, 0);
        int max_c = 0;
        for (int j = 0; j < b.size(); j++) {
            if (el == b[j]) {
                if (max_c == 19) {
                    return 20;
                }
                auto &v = dp[j];
                for (int l = 1; l < max_c + 1; l++) {
                    int best_we_can_do = min(carry[l], max(carry[l-1], el));
                    if (v[l]) {
                        v[l] = min(v[l], best_we_can_do);
                    } else {
                        v[l] = best_we_can_do;
                    }
                }
                int best_we_can_do = max(carry[max_c], el);
                if (v[max_c + 1])
                    v[max_c + 1] = min(v[max_c + 1], best_we_can_do);
                else
                    v[max_c + 1] = best_we_can_do;
                m = max(m, max_c + 1);
            }
            else if (el + c >= b[j]) {
                const auto &v = dp[j];
                for (int l = 1; l < 20; l++) {
                    if (v[l] > el + c || v[l] == 0) {
                        break;
                    } else {
                        if (carry[l] == 0) {
                            carry[l] = v[l];
                        } else {
                            carry[l] = min(carry[l], v[l]);
                        }
                        max_c = max(max_c, l);
                    }
                }
            }
        }
    }

    return m;
}

int main() {
    setupIO();
    int n = input_unsigned_number(), m = input_unsigned_number(), c = input_unsigned_number();
    vector<int> a, b;
    a.reserve(n);
    b.reserve(m);

    for (int i = 0; i < n; i++)
        a.push_back(input_unsigned_number());
    for (int i = 0; i < m; i++)
        b.push_back(input_unsigned_number());

    cout << longestCommonSubsequence(a, b, c);


    return 0;
}

//1 2 1
//1
//1 1