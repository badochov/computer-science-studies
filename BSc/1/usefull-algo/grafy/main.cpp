#include "graph.h"

#include <iostream>

using namespace std;

int main() {
    int n;
    cin >> n;
    vector<vector<int>> matrix(n, vector<int>(n));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cin >> matrix[i][j];
        }
    }
    Graph g = Graph<pair<unsigned int,int>>::fromAdjMatrix(matrix,0);
    cout << endl << "------" << endl;
    auto edges = g.getEdges();
    for (auto edge : edges) {
        cout << edge.first << " " << edge.second.first<<" "<<edge.second.second << endl;
    }
    return 0;
}