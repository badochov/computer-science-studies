#include "graph.h"

#include <vector>
#include <queue>
#include <iostream>
#include <tuple>

using namespace std;

template<typename Numeric>
vector<Numeric> _dijkstra(const Graph<pair<unsigned int, Numeric>> &g, unsigned int start) {
    vector<int> dist(g.size(), -1);
    vector<int> visited(g.size(), false);
    auto cmp = [](const auto &a, const auto &b) { return a.second > b.second; };
    priority_queue<pair<unsigned int, Numeric>, vector<pair<unsigned int, Numeric>>,
            decltype(cmp)> q(cmp);
    q.push(make_pair(start, 0));

    int v = 0, nV = 0;
    Numeric d, nD;
    while (!q.empty()) {
        tie(v, d) = q.top();
        q.pop();
        if (!visited[v]) {
            visited[v] = true;
            dist[v] = d;
            for (const auto &neighbour: g.neighbours[v]) {
                tie(nV, nD) = neighbour;
                if (!visited[nV]) {
                    q.push(make_pair(nV, nD + d));
                }
            }
        }
    }
    return dist;
}

template<typename T>
vector<T> dijkstra(const Graph<pair<unsigned int, T>> &g, unsigned int start) {
    if (is_arithmetic<T>{}) {
        return _dijkstra<T>(g, start);
    } else {
        return vector<T>();
    }
}


int main() {
    int n, s, x, y, d;
    cin >> n >> s;
    Graph<pair<unsigned int, int>> g(n);
    for (int i = 0; i < n; i++) {
        cin >> x >> y >> d;
        g.addEdge(x, make_pair(y, d));
        g.addEdge(y, make_pair(x, d));
    }
//    int n,s=0;
//    cin >> n;
//    vector<vector<int>> matrix(n, vector<int>(n));
//    for (int i = 0; i < n; i++) {
//        for (int j = 0; j < n; j++) {
//            cin >> matrix[i][j];
//        }
//    }
//    Graph g = Graph<pair<unsigned int, int>>::fromAdjMatrix(matrix, 0);
//    cout << endl << "------" << endl;
//    auto edges = g.getEdges();
//    for (auto edge : edges) {
//        cout << edge.first << " " << edge.second.first << " " << edge.second.second << endl;
//    }
    cout << "a";
    auto D = dijkstra<int>(g, s);
    cout << "b";
    cout << endl << "-----" << endl;
    for (int i = 0; i < D.size(); i++) {
        cout << i << " " << D[i] << endl;
    }
    return 0;
}