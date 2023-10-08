#include <bits/stdc++.h>

using namespace std;

#define XXX 0
#define XXG 1
#define XGX 2
#define GXX 4
#define XGG 3
#define GGX 6
#define GXG 5
#define GGG 7
#define INF numeric_limits<int>::max()

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}


void forward_dp(const vector<int> &old, vector<int> &new_dp, bool g) {
    if (g) {
        if (old[XXX] != INF) {
            new_dp[XXG] = min(new_dp[XXG], old[XXX]);
        }
        if (old[XXG] != INF) {
            new_dp[XGG] = min(new_dp[XGG], old[XXG] + 1);
        }
        if (old[XGX] != INF) {
            new_dp[GXG] = min(new_dp[GXG], old[XGX] + 1);
        }
        if (old[GXX] != INF) {
            new_dp[XXG] = min(new_dp[XXG], old[GXX] + 1);
        }
        if (old[XGG] != INF) {
            new_dp[GGG] = min(new_dp[GGG], old[XGG] + 2);
        }
        if (old[GGX] != INF) {
            new_dp[GXG] = min(new_dp[GXG], old[GGX] + 2);
        }
        if (old[GXG] != INF) {
            new_dp[XGG] = min(new_dp[XGG], old[GXG] + 2);
        }
        if (old[GGG] != INF) {
            new_dp[GGG] = min(new_dp[GGG], old[GGG] + 3);
        }
    } else {
        new_dp[XXX] = min(new_dp[XXX], min(old[XXX], old[GXX]));
        new_dp[XGX] = min(new_dp[XGX], min(old[GXG], old[XXG]));
        new_dp[GXX] = min(new_dp[GXX], min(old[XGX], old[GGX]));
        new_dp[GGX] = min(new_dp[GGX], min(old[GGG], old[XGG]));
    }
}

void self_forward_dp(vector<int> &new_dp, bool g) {
    vector<int> old = vector<int>(new_dp);
    new_dp = vector<int>(8, INF);
    forward_dp(old, new_dp, g);
}

int get_min(const vector<int> &dp) {
    int res = numeric_limits<int>::max();
    for (int n : dp) {
        res = min(res, n);
    }
    return res;
}

int main() {
    setupIO();
    int n, m;
    cin >> n >> m;
    string s;
    s.reserve(n);
    cin >> s;
    unordered_map<string, string> subs;
    vector<vector<int>> dp(n, vector<int>(8, INF));
    vector<vector<int>> dp_1(n, vector<int>(8, INF));
    vector<vector<int>> dp_2(n, vector<int>(8, INF));
    vector<vector<int>> dp_3(n, vector<int>(8, INF));
    for (int i = 0; i < m; i++) {
        pair<string, string> sub;
        cin >> sub.first >> sub.second;
        subs[sub.first] = sub.second;
    }

    if (s[0] == 'G') {
        dp[0][XXG] = 0;
    } else {
        dp[0][XXX] = 0;
    }

    for (int i = 1; i < n; i++) {
        forward_dp(dp[i - 1], dp[i], s[i] == 'G');
        forward_dp(dp_1[i - 1], dp[i], s[i] == 'G');
        forward_dp(dp_2[i - 1], dp_1[i], s[i] == 'G');
        forward_dp(dp_3[i - 1], dp_2[i], s[i] == 'G');
        if (i >= 2) {
            string pattern = s.substr(i - 2, 3);
            if (subs.count(pattern)) {
                string replacement = subs[pattern];
                if (i == 2) {
                    if (replacement[0] == 'G') {
                        dp_3[i][XXG] = 0;
                    } else {
                        dp_3[i][XXX] = 0;
                    }
                } else {
                    forward_dp(dp_3[i - 3], dp_3[i], replacement[0] == 'G');
                    forward_dp(dp_2[i - 3], dp_3[i], replacement[0] == 'G');
                    forward_dp(dp_1[i - 3], dp_3[i], replacement[0] == 'G');
                    forward_dp(dp[i - 3], dp_3[i], replacement[0] == 'G');
                }
                self_forward_dp(dp_3[i], replacement[1] == 'G');
                self_forward_dp(dp_3[i], replacement[2] == 'G');
            }
        }
    }

    cout << min(get_min(dp.back()), min(get_min(dp_1.back()), min(get_min(dp_2.back()), get_min(dp_3.back()))));


    return 0;
}

