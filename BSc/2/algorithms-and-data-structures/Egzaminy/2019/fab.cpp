#include <bits/stdc++.h>

using namespace std;

void setupIO() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
}


struct node {
    vector<int> connections;
    int color;

    explicit node(int c) : color(c) {
    }
};

int res = 0;


vector<node> cities;
vector<bool> visited;

vector<int> calc(int id) {
    const node &n = cities[id];
    visited[id] = true;
    vector<int> c(10, 0);

    for (int ch :n.connections) {
        if (!visited[ch]) {
            const node &child = cities[ch];
            auto d = calc(ch);
            for (int i = 0; i < 10; i++) {

                res = max(res, c[i] + d[i]);
                c[i] = max(c[i], d[i]);
            }
        }
    }
    for (int i = 0; i < 10; i++) {
        if (c[i] > 0)
            c[i]++;
    }

    if (c[n.color] == 0) {
        c[n.color]++;
    } else {
        res = max(res, c[n.color] - 1);
    }

    return c;
}

int main() {
    setupIO();
    unsigned long long n;
    cin >> n;
    cities.reserve(n);
    visited.reserve(n);

    for (int i = 0; i < n; i++) {
        visited.push_back(false);
        int c;
        cin >> c;
        cities.emplace_back(c - 1);
    }

    for (int i = 0; i < n - 1; i++) {
        int a, b;
        cin >> a >> b;
        cities[a - 1].connections.push_back(b - 1);
        cities[b - 1].connections.push_back(a - 1);
    }

    calc(0);

    cout << res;
    return 0;
}
