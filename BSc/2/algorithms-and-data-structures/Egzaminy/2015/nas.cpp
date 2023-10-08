#include <bits/stdc++.h>

using namespace std;


void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}

int main() {
    setupIO();
    int n, l;
    int maks = 0;
    cin >> n >> l;
    vector<vector<int>> in(n, vector<int>(l));
    unordered_map<int, int> curr;
    unordered_map<int, int> prev;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < l; j++) {
            cin >> in[i][j];
        }
    }

    for (int j = 0; j < l; j++) {
        for (int i = 0; i < n; i++) {
            int k = in[i][j];
            int p = 0;
            if (prev.count(k) > 0) {
                p = prev[k];
            }
            curr[k] = p + 1;

            maks = max(maks, curr[k]);
        }
        prev.swap(curr);
        curr.clear();
    }

    cout << maks;

    return 0;
}
