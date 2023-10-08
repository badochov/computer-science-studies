#include <bits/stdc++.h>

using namespace std;

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}


unsigned long long h[11];

int main() {
    setupIO();
    unsigned long long  n, res = 0;
    cin >> n;
    for (int i = 0; i < n; i++) {
        int p;
        cin >> p;
        h[p]++;
        for (int j = 1; j < p/2+1; j++) {
            if (j == p - j) {
                res += h[j] * (h[j] - 1) / 2;
            } else
                res += h[j] * h[p - j];
        }
    }
    cout << res;
    return 0;
}
