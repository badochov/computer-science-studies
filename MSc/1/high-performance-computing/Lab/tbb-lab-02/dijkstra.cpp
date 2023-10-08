#include <iostream>
#include <vector>
#include <unordered_map>
#include <limits>
#include <queue>
#include <unordered_set>
#include <chrono>
#include <random>
#include "tbb/tbb.h"

constexpr unsigned long INF = std::numeric_limits<unsigned long>::max();

class Graph {
  typedef std::pair<unsigned long, unsigned long> distance_vertex;

  std::unordered_map<unsigned long, std::unordered_map<unsigned long, unsigned long>> adj;
  unsigned long vertices;
 public:

  explicit Graph(unsigned long v) : vertices(v) {}

  bool addEdge(unsigned long from, unsigned long to, unsigned long weight) {
    auto it = adj.find(from);
    if (it == adj.end()) {
      adj[from] = {{to, weight}};
    } else {
      if (it->second.count(to) == 0) {
        it->second[to] = weight;
      } else {
        return false;
      }
    }
    return true;
  }

  std::vector<unsigned long> dijkstra_seq(unsigned long source) {
    std::priority_queue<distance_vertex, std::vector<distance_vertex>, std::greater<>> pq;
    std::vector<unsigned long> res(vertices, INF);
    res[source] = 0;
    std::unordered_set<unsigned long> visited;

    pq.emplace(0, source);

    while (!pq.empty()) {
      auto [d, v] = pq.top();
      pq.pop();

      if (visited.count(v)) {
        continue;
      }
      visited.emplace(v);

      for (auto const [nv, nd] : adj[v]) {
        if (res[nv] > nd + d) {
          res[nv] = nd + d;
          pq.emplace(nd + d, nv);
        }
      }
    }

    return res;
  }

  std::vector<unsigned long> dijkstra_par(unsigned long source) {
    tbb::concurrent_priority_queue<distance_vertex, std::greater<>> pq;
    tbb::concurrent_vector<unsigned long> res(vertices, INF);
    tbb::spin_mutex mutexes[vertices];

    pq.emplace(0, source);
    bool init;

    tbb::parallel_for_each(&init, &init + 1, [&mutexes, &pq, &res, this]
        (bool _, tbb::feeder<bool> &feeder) {
      distance_vertex vd;
      pq.try_pop(vd);
      auto [d, v] = vd;

      mutexes[v].lock();
      if (res[v] != d) {
        mutexes[v].unlock();
        return;
      }
      mutexes[v].unlock();

      for (auto const [nv, nd] : adj[v]) {
        mutexes[nv].lock();
        if (res[nv] > nd + d) {
          pq.emplace(nd + d, nv);
          feeder.add(true);
        }
        mutexes[nv].unlock();
      }
    });

    std::vector<unsigned long> r(vertices);
    for (unsigned long i = 0; i < vertices; i++) {
      r[i] = res[i];
    }
    return r;
  }

  void generate(unsigned long edges) {
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<unsigned long> distr(0, vertices - 1);

    std::random_device rd2;
    std::mt19937 gen2(rd2());
    std::uniform_int_distribution<unsigned long> distr2(1, 128);

    for (unsigned long i = 0; i < edges; i++) {
      unsigned long from, to, length;
      do {
        from = distr(gen);
        to = distr(gen);
        length = distr2(gen2);
      } while (!addEdge(from, to, length));
    }
  }
};

int main(int argc, char **argv) {
  if (argc != 3) {
    std::cerr << "Usage: " << argv[0] << "<vertices> <edges>";
  }
  unsigned long vertices = strtoul(argv[1], nullptr, 0);
  unsigned long edges = strtoul(argv[2], nullptr, 0);
  Graph g(vertices);

  g.generate(edges);

  auto start = std::chrono::high_resolution_clock::now();
  auto seq = g.dijkstra_seq(0);
  auto stop = std::chrono::high_resolution_clock::now();

  std::cout << "Time sequential: " << std::chrono::duration_cast<std::chrono::microseconds>(stop - start).count()
            << "us" << std::endl;

  start = std::chrono::high_resolution_clock::now();
  auto par = g.dijkstra_par(0);
  stop = std::chrono::high_resolution_clock::now();

  std::cout << "Time parallel: " << std::chrono::duration_cast<std::chrono::microseconds>(stop - start).count()
            << "us" << std::endl;

  bool ok = true;
  for (unsigned long i = 0; i < vertices; i++) {
    if (seq[i] != par[i]) {
      std::cerr << "Mismatch at " << i << " seq: " << seq[i] << " par: " << par[i] << std::endl;
      ok = false;
    }
  }
  if (ok) {
    std::cout << "Outputs match" << std::endl;
    return 0;
  }
  return 1;
}