#ifndef GRAFY_GRAPH_H
#define GRAFY_GRAPH_H

#include <vector>
#include <algorithm>

template<class T>
class Graph {
  public:
    explicit Graph(unsigned int maxIndex = 0);

    Graph<T> *addEdge(unsigned int s, T e);

    Graph<T> *addEdge(const std::pair<unsigned int, T> &edge);

    Graph<T> *addEdges(const std::vector<std::pair<unsigned int, T>> &edges);

    std::vector<std::pair<unsigned int, T>> getEdges();

    [[nodiscard]] unsigned int size() const;

    std::vector<std::vector<T>> neighbours;

    template<class Dist>
    static Graph<std::pair<unsigned int, Dist>>
    fromAdjMatrix(const std::vector<std::vector<Dist>> &adjMatrix, Dist noPath = -1);
};

template<class T>
Graph<T>::Graph(unsigned int maxIndex) {
    this->neighbours = std::vector<std::vector<T>>(maxIndex, std::vector<T>());
}


template<class T>
Graph<T> *Graph<T>::addEdge(unsigned int s, T e) {
    if (s >= this->neighbours.size()) {
        this->neighbours.resize(s + 1);
    }
    this->neighbours[s].push_back(e);
    return this;
}

template<class T>
Graph<T> *Graph<T>::addEdge(const std::pair<unsigned int, T> &edge) {
    return this->addEdge(edge.first, edge.second);
}

template<class T>
Graph<T> *Graph<T>::addEdges(const std::vector<std::pair<unsigned int, T>> &edges) {
    for (const std::pair<int, T> &edge: edges) {
        this->addEdge(edge);
    }
    return this;
}

template<class T>
std::vector<std::pair<unsigned int, T>> Graph<T>::getEdges() {
    std::vector<std::pair<unsigned int, T>> edges;
    std::pair<unsigned int, T> edge;
    for (typename std::vector<std::pair<unsigned int, T>>::size_type i = 0; i < this->neighbours.size(); i++) {
        edge.first = i;
        for (T neighbour : this->neighbours[i]) {
            edge.second = neighbour;
            edges.push_back(edge);
        }
    }
    return edges;
}

template<class T>
unsigned int Graph<T>::size() const {
    return this->neighbours.size();
}

template<class T>
template<class Dist>
Graph<std::pair<unsigned int, Dist>>
Graph<T>::fromAdjMatrix(const std::vector<std::vector<Dist>> &adjMatrix, Dist noPath) {
    int width = adjMatrix.size();
    if (width != 0) {
        int height = adjMatrix[0].size();
        Graph<std::pair<unsigned int, Dist>> g(std::max(height, width));
        for (int x = 0; x < width; x++) {
            for (int y = 0; y < height; y++) {
                if (adjMatrix[x][y] != noPath) {
                    g.addEdge(x, std::make_pair(y, adjMatrix[x][y]));
                }
            }
        }

        return g;
    } else {
        return Graph<std::pair<unsigned int, Dist>>();
    }
}


#endif //GRAFY_GRAPH_H
