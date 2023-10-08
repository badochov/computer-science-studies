#include <bits/stdc++.h>

using namespace std;

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}

int main() {
    setupIO();
    int n, m;
    cin >> n >> m;
    set<pair<int, int>> left, right;
    for (int i = 0; i < m; i++) {
        pair<int, int> p;
        char s;
        cin >> s >> p.first >> p.second;

        if (s == '-') {
            if (p.first == 1) {
                left.erase(p);
            } else {
                right.erase(p);
            }
        } else {
            if (p.first == 1) {
                left.insert(p);
            } else {
                right.insert(p);
            }
        }
        int r = right.empty() ? n + 1 : right.begin()->first;
        int l = left.empty() ? 0 : left.rbegin()->second;
        int res = max(0, r - l - 1);
        cout << res << endl;
    }


    return 0;
}
