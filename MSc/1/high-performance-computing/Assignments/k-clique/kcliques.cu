#include <iostream>
#include <fstream>
#include <unordered_map>
#include <vector>
#include <algorithm>

#include <vector>
#include <stdexcept>
#include <iostream>

struct CSR {
  std::vector<size_t> row_ptr;
  std::vector<size_t> col_ind;
};

namespace gpu {
constexpr size_t WARP_SIZE = 32;
constexpr size_t GROUPS_PER_BLOCK = 4;
constexpr size_t THREADS = WARP_SIZE * GROUPS_PER_BLOCK;
constexpr size_t BLOCKS = 256;
constexpr size_t MAX_K = 12;
constexpr size_t STACK_SIZE = MAX_K - 2; // Last frame and first frame are redundant.
constexpr size_t MAX_NEIGHBOURS = 1024;
constexpr unsigned long long MOD = 1e9;

__device__ const size_t FULL_MASK = 0xffffffff;

struct BitwiseAdjList {
  static constexpr size_t LIST_SIZE = MAX_NEIGHBOURS / sizeof(uint32_t) / 8;
  uint32_t list[LIST_SIZE];

  __device__ bool is_neighbour(unsigned int y) {
    return ((list[y / 32]) & (1u << (y % 32))) != 0;
  }

  __device__  int intersect(const BitwiseAdjList &other, BitwiseAdjList *result) {
    int localNeighbours = 0;
    for (unsigned int i = threadIdx.x; i < LIST_SIZE; i += blockDim.x) {
      result->list[i] = list[i] & other.list[i];
      localNeighbours += __popc(result->list[i]);
    }
    return localNeighbours;
  }

  __device__  int intersect(const BitwiseAdjList &other) {
    int localNeighbours = 0;
    for (unsigned int i = threadIdx.x; i < LIST_SIZE; i += blockDim.x) {
      auto el = list[i] & other.list[i];
      localNeighbours += __popc(el);
    }
    return localNeighbours;
  }

  __device__ int copy(const BitwiseAdjList &other) {
    int local_neighbours = 0;
    for (unsigned int i = threadIdx.x; i < LIST_SIZE; i += blockDim.x) {
      list[i] = other.list[i];
      local_neighbours += __popc(list[i]);
    }
    return local_neighbours;
  }
};

__device__ unsigned threadTotalIdx() {
  return threadIdx.x + threadIdx.y * blockDim.x;
}

// AdjMatrix stores adjacency matrix using binary encoding.
struct AdjMatrix {
  BitwiseAdjList matrix[MAX_NEIGHBOURS];
  size_t dirty_list_idx[MAX_NEIGHBOURS];

  __device__ void fill(const size_t *subgraph_vertices,
                       size_t subgraph_vertices_count,
                       const size_t *row_ptr,
                       const size_t *col_ind) {
    // Each threads processes different row.
    for (unsigned int remapV = threadTotalIdx(); remapV < subgraph_vertices_count; remapV += blockDim.x * blockDim.y) {
      size_t remapIdx = 0;
      size_t orgV = subgraph_vertices[remapV];
      uint32_t neighbours = 0;
      for (size_t i = row_ptr[orgV]; i < row_ptr[orgV + 1]; i++) {
        // Neighbours are always in ascending order.
        size_t neighbour = col_ind[i];
        while ((remapIdx < subgraph_vertices_count) && (neighbour > subgraph_vertices[remapIdx])) {
          remapIdx++;
          if (remapIdx % 32 == 0) {
            matrix[remapV].list[(remapIdx / 32) - 1] = neighbours;
            neighbours = 0;
          }
        }
        // End of possible neighbours;
        if (remapIdx >= subgraph_vertices_count) {
          break;
        }
        // Add neighbour.
        if (neighbour == subgraph_vertices[remapIdx]) {
          neighbours |= 1u << (remapIdx % 32);
        }
      }
      size_t last_idx = remapIdx / 32;
      matrix[remapV].list[last_idx] = neighbours;
      // Clean trailing zeroes.
      for (size_t idx = last_idx + 1; idx <= dirty_list_idx[remapV]; idx++) {
        matrix[remapV].list[idx] = 0;
      }
      dirty_list_idx[remapV] = last_idx;
    }
  }
};

__device__ void reduceBlockRes(unsigned long long blockRes[THREADS / 2][MAX_K - 2],
                               const unsigned long long localRes[MAX_K],
                               unsigned long k) {
// Process first row separately to save space.
  for (int i = 1; i < k - 1; i++) { // First cell is not processed.
    atomicAdd(&blockRes[threadTotalIdx() / 2][i - 1], localRes[i]);
  }
  __syncthreads();

  for (int diff = 2; diff < THREADS; diff *= 2) {
    if ((threadTotalIdx() % diff == 0) && (threadTotalIdx() % (diff * 2) != 0)) {
      for (int i = 0; i < k - 2; i++) { // First cell is not processed.
        atomicAdd(&blockRes[(threadTotalIdx() - diff) / 2][i], blockRes[threadTotalIdx() / 2][i]);
      }
    }
    __syncthreads();
  }
}

__device__ void reduceRes(unsigned long long int blockRes[THREADS / 2][MAX_K - 2],
                          const unsigned long long localRes[MAX_K],
                          unsigned long long int *out,
                          unsigned long k) {
  atomicAdd(&out[0], localRes[0]);
  for (size_t i = 1; i < k - 1; i++) {
    atomicAdd(&out[i], blockRes[0][i - 1] % MOD);
  }
}

__global__ void clique_kernel(size_t *row_ptr,
                              size_t row_ptr_size,
                              size_t *col_ind,
                              unsigned long k,
                              unsigned long long *out,
                              unsigned long long *counter,
                              AdjMatrix matrices[BLOCKS]) {
  bool isGroupLeader = threadIdx.x == 0;
  bool isBlockLeader = isGroupLeader && threadIdx.y == 0;
  unsigned int groupId = threadIdx.y;

  __shared__ unsigned long long idx;
  __shared__ unsigned long long blockCounter;
  __shared__ unsigned long long groupVertexIdx[GROUPS_PER_BLOCK];

  __shared__ BitwiseAdjList blockStack[GROUPS_PER_BLOCK][STACK_SIZE];

  BitwiseAdjList *stack = blockStack[groupId];
  AdjMatrix &adj = matrices[blockIdx.x];

  size_t stackListIdx[STACK_SIZE];

  size_t *subgraph_vertices;
  size_t subgraph_vertices_count;

  unsigned long long localRes[MAX_K - 1] = {};

  for (;;) {
    __syncthreads(); // Do not get another job until all groups have finished processing.
    if (isBlockLeader) {
      idx = atomicAdd(counter, 1);
      blockCounter = 0;
    }
    __syncthreads();

    if (idx + 1 >= row_ptr_size) {
      break;
    }

    subgraph_vertices = &col_ind[row_ptr[idx]];
    subgraph_vertices_count = row_ptr[idx + 1] - row_ptr[idx];
    localRes[0] += subgraph_vertices_count;
    localRes[0] %= MOD;

    if (k == 2) {
      continue;
    }

    // Vertices are remapped.
    adj.fill(subgraph_vertices, subgraph_vertices_count, row_ptr, col_ind);
    __syncthreads();

    // From now on the work is split between groups.
    for (;;) {
      __syncwarp();
      if (isGroupLeader) {
        groupVertexIdx[groupId] = atomicAdd(&blockCounter, 1);
      }
      __syncwarp();

      if (groupVertexIdx[groupId] >= subgraph_vertices_count) {
        break;
      }

      localRes[1] += stack[0].copy(adj.matrix[groupVertexIdx[groupId]]);
      localRes[1] %= MOD;
      if (k == 3) {
        continue;
      }
      stackListIdx[0] = 0;

//    𝑛𝑢𝑚𝐶𝑙𝑖𝑞𝑢𝑒𝑠 = 0
//    procedure 𝑡𝑟𝑎𝑣𝑒𝑟𝑠𝑒𝑆𝑢𝑏𝑡𝑟𝑒𝑒 (𝐺, 𝑘, ℓ, 𝐼)
//    for 𝑣 ∈ 𝐼
//      𝐼′ = 𝐼 ∩ 𝐴𝑑𝑗𝐺 (𝑣)
//      if ℓ + 1 == 𝑘
//        𝑛𝑢𝑚𝐶𝑙𝑖𝑞𝑢𝑒𝑠 + = |𝐼′|
//      else if |𝐼′| > 0
//        𝑡𝑟𝑎𝑣𝑒𝑟𝑠𝑒𝑆𝑢𝑏𝑡𝑟𝑒𝑒 (𝐺, 𝑘, ℓ + 1, 𝐼′)

      unsigned long long level = 2;
      while (level != 1) {
        bool recurse = false;
        const bool lastLevel = level == k;
        BitwiseAdjList &frame = stack[level - 2];
        size_t &frame_list_idx = stackListIdx[level - 2];

        while (frame_list_idx < subgraph_vertices_count) {
          size_t vertex = frame_list_idx;
          frame_list_idx++;
          if (!frame.is_neighbour(vertex)) {
            continue;
          }

          int num_neighbours;
          if (lastLevel) {
            num_neighbours = frame.intersect(adj.matrix[vertex]);
          } else {
            num_neighbours = frame.intersect(adj.matrix[vertex], &stack[level - 1]);
          }
          localRes[level] += num_neighbours;
          localRes[level] %= MOD;

          if (!lastLevel && __ballot_sync(FULL_MASK, num_neighbours)) {
            // At least one neighbour found.
            recurse = true;
            break;
          }
        }

        if (recurse) {
          stackListIdx[level - 1] = 0;
          level++;
        } else {
          level--;
        }
      }
    }
  }

  // Reduce local answers to a block answer.
  __shared__ unsigned long long blockRes[THREADS / 2][MAX_K - 2];

  for (unsigned i = threadTotalIdx() % 2; i < MAX_K - 2; i += 2) {
    blockRes[threadTotalIdx() / 2][i] = 0;
  }
  __syncthreads();

  reduceBlockRes(blockRes, localRes, k);

  if (isBlockLeader) {
    reduceRes(blockRes, localRes, out, k);
  }
}

void handle_error(cudaError_t err) {
  if (err != cudaSuccess) {
    throw std::runtime_error("CUDA operations resulted in failure, code: " + std::to_string(err));
  }
}

std::vector<unsigned long long> find_cliques(unsigned long k, const CSR &csr) {
  std::vector<unsigned long long> ret(k);
  if (k > 1) {
    unsigned long long *out;
    size_t *row_ptr;
    size_t *col_ind;

    handle_error(cudaMalloc(&out, (k - 1) * sizeof(unsigned long long))); // TODO scale out down.
    handle_error(cudaMalloc(&row_ptr, csr.row_ptr.size() * sizeof(size_t)));
    handle_error(cudaMalloc(&col_ind, csr.col_ind.size() * sizeof(size_t)));

    handle_error(cudaMemcpy(row_ptr, csr.row_ptr.data(), csr.row_ptr.size() * sizeof(size_t), cudaMemcpyHostToDevice));
    handle_error(cudaMemcpy(col_ind, csr.col_ind.data(), csr.col_ind.size() * sizeof(size_t), cudaMemcpyHostToDevice));

    handle_error(cudaMemset(out, 0, (k - 1) * sizeof(unsigned long long)));

    unsigned long long *counter;
    handle_error(cudaMalloc(&counter, sizeof(unsigned long long)));
    handle_error(cudaMemset(counter, 0, sizeof(unsigned long long)));

    AdjMatrix *matrices;
    handle_error(cudaMalloc(&matrices, sizeof(AdjMatrix) * BLOCKS));
    handle_error(cudaMemset(matrices, 0, sizeof(AdjMatrix) * BLOCKS));

    dim3 threadsDim(WARP_SIZE, GROUPS_PER_BLOCK);
    clique_kernel<<<BLOCKS, threadsDim>>>(row_ptr,
                                          csr.row_ptr.size(),
                                          col_ind,
                                          k,
                                          out,
                                          counter,
                                          matrices);

    handle_error(cudaPeekAtLastError());
    handle_error(cudaDeviceSynchronize());

    handle_error(cudaMemcpy(ret.data() + 1, out, (k - 1) * sizeof(unsigned long long), cudaMemcpyDeviceToHost));

    handle_error(cudaFree(matrices));
    handle_error(cudaFree(out));
    handle_error(cudaFree(row_ptr));
    handle_error(cudaFree(col_ind));
    handle_error(cudaFree(counter));

    for (unsigned long long &el : ret) {
      el %= MOD;
    }
  }

  ret[0] = csr.row_ptr.size() - 1; // Number of 1-cliques is the number of vertices.

  return ret;
}

}

typedef uint32_t vid;

struct UndirectedEdge {
  vid a;
  vid b;

  UndirectedEdge(vid _a, vid _b) : a(_a), b(_b) {}
};

struct Edge {
  vid from;
  vid to;

  Edge(vid _from, vid _to) : from(_from), to(_to) {}
};

struct Graph {
  //    Vertices are numbered by consecutive numbers starting with 1.
  std::vector<UndirectedEdge> edges;
  vid vLimit;

  std::vector<Edge> toOrientedEdges() {
    std::unordered_map<vid, size_t> neighbours_count;
    for (const auto &e : edges) {
      neighbours_count[e.a]++;
      neighbours_count[e.b]++;
    }

    std::vector<Edge> orientedEdges;
    for (const auto &e : edges) {
      if (neighbours_count[e.a] < neighbours_count[e.b]) {
        orientedEdges.emplace_back(e.a, e.b);
      } else if (neighbours_count[e.a] > neighbours_count[e.b]) {
        orientedEdges.emplace_back(e.b, e.a);
      } else {
        if (e.a < e.b) {
          orientedEdges.emplace_back(e.a, e.b);
        } else {
          orientedEdges.emplace_back(e.b, e.a);
        }
      }
    }

    return orientedEdges;
  }

  CSR toCSR() {
    auto orientedEdges = toOrientedEdges();
    std::sort(orientedEdges.begin(),
              orientedEdges.end(),
              [](const Edge &a, const Edge &b) {
                if (a.from == b.from) {
                  return a.to < b.to;
                }
                return a.from < b.from;
              });
    CSR csr{};
    csr.row_ptr.reserve(vLimit + 1);

    size_t curr_ptr = 0;
    for (const auto &e : orientedEdges) {
      while (csr.row_ptr.size() <= e.from) {
        csr.row_ptr.push_back(curr_ptr);
      }
      curr_ptr++;
      csr.col_ind.push_back(e.to);
    }
    while (csr.row_ptr.size() <= vLimit) {
      csr.row_ptr.push_back(curr_ptr);
    }

    return csr;
  }
};

Graph read_graph(std::ifstream &in) {
  Graph g{};
  std::unordered_map<vid, vid> mapping;
  vid currVid = 0;
  vid a, b;
  while (in >> a >> b) {
    auto pA = mapping.find(a);
    if (pA == mapping.end()) {
      pA = mapping.insert({a, currVid}).first;
      currVid++;
    }

    auto pB = mapping.find(b);
    if (pB == mapping.end()) {
      pB = mapping.insert({b, currVid}).first;
      currVid++;
    }

    g.edges.emplace_back(pA->second, pB->second);
  }

  g.vLimit = currVid;

  return g;
}

void kclique(std::ifstream &in, unsigned long k, std::ofstream &out) {
  Graph g = read_graph(in);

  auto ret = gpu::find_cliques(k, g.toCSR());

  for (size_t i = 0; i < k; i++) {
    out << ret[i] << " ";
  }
  out << std::endl;
}

int main(int argc, char **argv) {
  if (argc != 4) {
    std::cerr << "Usage: " << argv[0] << " <Graph input file> <k value> <output file>" << std::endl;
    return 1;
  }
  std::ifstream in{argv[1]};
  if (!in.is_open()) {
    std::cerr << "Input file: " << argv[1] << " doesn't exist." << std::endl;
    return 1;
  }
  char *endptr;
  unsigned long k = strtoul(argv[2], &endptr, 0);
  if (*endptr != '\0') {
    std::cerr << "Invalid k value." << std::endl;
    return 1;
  }
  if (k == 0 || k > 12) {
    std::cerr << "K value must be in range (0, 12]." << std::endl;
    return 1;
  }
  std::ofstream out{argv[3]};
  if (!out.is_open()) {
    std::cerr << "Error opening output file: " << argv[3] << std::endl;
    return 1;
  }

  kclique(in, k, out);

  return 0;
}
