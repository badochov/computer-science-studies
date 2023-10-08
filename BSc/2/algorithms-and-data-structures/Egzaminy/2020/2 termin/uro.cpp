#include <bits/stdc++.h>

using namespace std;

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}


constexpr long long MOD = 1e9;

long long normalize(long long n) {
    return (n + MOD) % MOD;
}

// 20
//1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
int main() {
    setupIO();
    long long n;
    cin >> n;
    vector<long long> dp(n + 1, 0);
    long long sum = 0;

    for (long long i = 0; i < n; ++i) {
        long long a;
        cin >> a;
        long long old = dp[a];
        dp[a] = normalize(sum + 1);
        sum = normalize(sum + dp[a] - old);
    }

    cout << sum;


    return 0;
}
