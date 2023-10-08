#include <iostream>
#include <stdexcept>
#include <unistd.h>
#include <utility>
#include <vector>
#include <cassert>
#include <cstring>

#include <mpi.h>

#include "types.hpp"
#include "params.hpp"
#include "densematgen.h"

// Following error checking code is from: https://stackoverflow.com/a/75458495
#define MPI_CHECK(n) check_mpi_error(__FILE__, __LINE__, n)
inline void check_mpi_error(const char *file, const int line, const int n) {
  char errbuffer[MPI::MAX_ERROR_STRING];
  int errlen;

  if (n != MPI_SUCCESS) {
    MPI_Error_string(n, errbuffer, &errlen);
    printf("MPI-error: %s\n", errbuffer);
    printf("Location: %s:%i\n", file, line);
    MPI::COMM_WORLD.Abort(n);
  }
}

enum {
  FRAGMENT_TAG,
  SHIFT_TAG_A,
  SHIFT_TAG_B,
  RESULTS_TAG,
  PRINT_TAG,
};

struct MpiInfo {
  int numProcesses;
  int myRank;

  static constexpr int LEADER_RANK = 0;

  bool isLeader() const {
    return myRank == LEADER_RANK;
  }
};

struct Grid {
  const Dim3 dim;
  MPI_Comm comm;

  explicit Grid(Dim3 _dim, int rank) : dim(_dim), comm(createComm(rank)) {}

  unsigned long replicationFactor(bool isA) const {
    if (isReplicatedA() == isA) {
      return num_cannon_groups();
    }
    return 1;
  }

  unsigned long replicationGroupSize(bool isA) const {
    return pkGroupSize() / replicationFactor(isA);
  }

  unsigned long cannonGroupSize() const {
    return cannonDim() * cannonDim();
  }

  unsigned long num_cannon_groups() const {
    return std::max(dim.m, dim.n) / std::min(dim.m, dim.n);
  }

  unsigned long pkGroupSize() const {
    return dim.m * dim.n;
  }

  unsigned long cannonDim() const {
    return std::min(dim.m, dim.n);
  }

  std::pair<unsigned long, unsigned long> frags(bool isA) const {
    if (isA) {
      unsigned long xFrags = dim.n / replicationFactor(isA) * dim.k;
      unsigned long yFrags = dim.m;

      return {xFrags, yFrags};
    } else {
      unsigned long xFrags = dim.n;
      unsigned long yFrags = dim.m / replicationFactor(isA) * dim.k;

      return {xFrags, yFrags};
    }

  }

  std::pair<unsigned long, unsigned long> pk(int rank) const {
    return {rank / pkGroupSize(), rank % pkGroupSize()};
  }

  std::pair<unsigned long, unsigned long> cannon(int rank) const {
    unsigned long pkId = pk(rank).second;

    return {pkId / cannonGroupSize(), rank % cannonGroupSize()};
  }

  std::pair<unsigned long, unsigned long> xy(int rank, bool isA) const {
    unsigned long pkGroup = rank / pkGroupSize();

    if (isA) {
      unsigned long groupStartX = pkGroup * (dim.n / replicationFactor(isA));
      unsigned long inGroupX = (rank % replicationGroupSize(isA)) / dim.m;

      return {groupStartX + inGroupX, rank % dim.m};
    } else {
      unsigned long x = (rank % replicationGroupSize(isA)) / (dim.m / replicationFactor(isA));
      unsigned long groupStartY = pkGroup * (dim.m / replicationFactor(isA));
      unsigned long inGroupY = rank % (dim.m / replicationFactor(isA));

      return {x, groupStartY + inGroupY};
    }
  }

  unsigned long getRankForFragment(unsigned long x, unsigned long y, bool isA) const {
    unsigned long pkGroup, cannonGroup;
    unsigned long cannonGroupOffset = (x % cannonDim()) * cannonDim() + (y % cannonDim());
    if (isA) {
      pkGroup = x / cannonDim();
      cannonGroup = y / cannonDim();
    } else {
      pkGroup = y / cannonDim();
      cannonGroup = x / cannonDim();
    }
    return pkGroup * pkGroupSize() + cannonGroup * cannonGroupSize() + cannonGroupOffset;
  }

  bool isReplicatedA() const {
    return dim.n > dim.m;
  }

  unsigned long getRankForResultFragment(unsigned long x, unsigned long y) const {
    if (isReplicatedA()) {
      unsigned long cannonGroup = x / cannonDim();
      unsigned long offsetInCannon = (x % cannonDim()) * cannonDim() + y;

      return cannonGroup * cannonGroupSize() + offsetInCannon;
    } else {
      unsigned long cannonGroup = y / cannonDim();
      unsigned long offsetInCannon = x * cannonDim() + (y % cannonDim());

      return cannonGroup * cannonGroupSize() + offsetInCannon;
    }
  }

  unsigned long size() const {
    return dim.m * dim.n * dim.k;
  }

 private:
  MPI_Comm createComm(int rank) const {
    MPI_Comm c;

    int color = rank < size() ? 0 : MPI_UNDEFINED;
    MPI_CHECK(MPI_Comm_split(MPI_COMM_WORLD, color, rank, &c));

    return c;
  }
};

struct MatrixFragment {
  const unsigned long width;
  const unsigned long height;
  const unsigned long x;
  const unsigned long y;

  double *data;

  MatrixFragment(unsigned long _width,
                 unsigned long _height,
                 unsigned long _x,
                 unsigned long _y) : width(_width),
                                     height(_height),
                                     x(_x),
                                     y(_y),
                                     data(new double[width * height]) {}

  unsigned long size() const {
    return width * height;
  }

  ~MatrixFragment() {
    delete[] data;
  }

  double &get(unsigned long col, unsigned long row) const {
    return data[col * height + row];
  }

  void clean() const {
    memset(data, 0, sizeof(double) * size());
  }

  void copyRow(double *buffer, unsigned long row) const {
    for (unsigned long col = 0; col < width; col++) {
      buffer[col] = get(col, row);
    }
  }
};

class Matrix {
 public:
  const unsigned long width;
  const unsigned long height;

  const bool isA;

  MatrixFragment fragment;

 private:
  MatrixFragment createFragment(const Grid &grid, int rank) const {
    auto frags = grid.frags(isA);
    unsigned long xFrags = frags.first, yFrags = frags.second;
    auto xy = grid.xy(rank, isA);
    unsigned long x = xy.first, y = xy.second;

    unsigned long fragX = (width + xFrags - 1) / xFrags;
    unsigned long fragY = (height + yFrags - 1) / yFrags;

    return MatrixFragment{fragX, fragY, x, y};
  }

 public:
  Matrix(unsigned long _width,
         unsigned long _height,
         const Grid &grid,
         int _rank,
         bool _isA)
      : width(_width), height(_height), isA(_isA), fragment(createFragment(grid, _rank)) {}

  void generateMatrixFragment(double *frag, int seed, unsigned long x, unsigned long y) const {
    for (unsigned long i = 0; i < fragment.width; i++) {
      unsigned long globalX = x * fragment.width + i;
      for (unsigned long j = 0; j < fragment.height; j++) {
        unsigned long globalY = y * fragment.height + j;
        if (globalX >= width || globalY >= height) {
          frag[i * fragment.height + j] = 0;
        } else {
          frag[i * fragment.height + j] = generate_double(seed, globalY, globalX);
        }
      }
    }
  }
};

class Result {
 public:
  const unsigned long width;
  const unsigned long height;

  MatrixFragment fragment;

 private:
  MatrixFragment createFragment(const Grid &grid, int rank) const {
    unsigned long fragX = (width + grid.dim.n - 1) / grid.dim.n;
    unsigned long fragY = (height + grid.dim.m - 1) / grid.dim.m;

    auto cannonData = grid.cannon(rank);
    unsigned long cannonGroup = cannonData.first, cannonId = cannonData.second;

    unsigned long xInCannon = cannonId / grid.cannonDim();
    unsigned long yInCannon = cannonId % grid.cannonDim();

    unsigned long xCannon;
    unsigned long yCannon;

    if (grid.isReplicatedA()) {
      xCannon = cannonGroup;
      yCannon = 0;
    } else {
      xCannon = 0;
      yCannon = cannonGroup;
    }

    unsigned long x = xCannon * grid.cannonDim() + xInCannon;
    unsigned long y = yCannon * grid.cannonDim() + yInCannon;

    return MatrixFragment{fragX, fragY, x, y};
  }

 public:
  Result(unsigned long _width,
         unsigned long _height,
         const Grid &grid,
         int _rank)
      : width(_width), height(_height), fragment(createFragment(grid, _rank)) {}

  unsigned long long overGeInFragment(double ge) const {
    unsigned long long res = 0;

    for (unsigned long i = 0; i < fragment.width; i++) {
      unsigned long globalX = fragment.x * fragment.width + i;
      if (globalX >= width) {
        break;
      }
      for (unsigned long j = 0; j < fragment.height; j++) {
        unsigned long globalY = fragment.y * fragment.height + j;
        if (globalY >= height) {
          break;
        }
        if (fragment.get(i, j) > ge) {
          res++;
        }
      }
    }

    return res;
  }
};

class CA3DMM {
  const MpiInfo mpi;
  const Grid grid;

  Matrix a;
  Matrix b;
  Result result;

  MPI_Comm pkComm;
  MPI_Comm cannonComm;
  MPI_Comm replicaComm;
  MPI_Comm resultsComm;

  double *const bufferA;
  double *const bufferB;
  double *const bufferResult;

  Matrix createA(const Dim3 &matrices_dim) const {
    return Matrix{matrices_dim.k, matrices_dim.m, grid, mpi.myRank, true};
  }

  Matrix createB(const Dim3 &matrices_dim) const {
    return Matrix{matrices_dim.n, matrices_dim.k, grid, mpi.myRank, false};
  }

  Result createResult(const Dim3 &matrices_dim) const {
    return Result{matrices_dim.n, matrices_dim.m, grid, mpi.myRank};
  }

  MPI_Comm createPkComm() const {
    MPI_Comm pk;
    auto pkData = grid.pk(mpi.myRank);
    unsigned long pkGroup = pkData.first, pkId = pkData.second;

    MPI_CHECK(MPI_Comm_split(grid.comm, pkGroup, pkId, &pk));

    return pk;
  }

  MPI_Comm createCannonComm() const {
    MPI_Comm cannon;
    auto cannonData = grid.cannon(mpi.myRank);
    unsigned long cannonGroup = cannonData.first, cannonId = cannonData.second;

    MPI_CHECK(MPI_Comm_split(pkComm, cannonGroup, cannonId, &cannon));

    return cannon;
  }

  MPI_Comm createReplicaComm() const {
    MPI_Comm replica;
    auto cannonData = grid.cannon(mpi.myRank);
    unsigned long cannonGroup = cannonData.first, cannonId = cannonData.second;

    MPI_CHECK(MPI_Comm_split(pkComm, cannonId, cannonGroup, &replica));

    return replica;
  }

  MPI_Comm createResultsComm() const {
    MPI_Comm results;

    MPI_CHECK(MPI_Comm_split(grid.comm, mpi.myRank % grid.pkGroupSize(), mpi.myRank, &results));

    return results;
  }

  unsigned long skewRank(unsigned long rank, bool isA) const {
    unsigned long cannonGroupStart = rank - (rank % grid.cannonGroupSize());

    unsigned long cannonId = grid.cannon(rank).second;
    unsigned long cannonDim = grid.cannonDim();

    unsigned long col = cannonId / cannonDim;
    unsigned long row = cannonId % cannonDim;

    unsigned long dstCol, dstRow;
    if (isA) {
      dstRow = row;
      dstCol = ((cannonDim + col) - row) % cannonDim;
    } else {
      dstRow = ((cannonDim + row) - col) % cannonDim;
      dstCol = col;
    }
    return cannonGroupStart + dstCol * cannonDim + dstRow;
  }

  void generateAndDistributeMatrices(std::pair<int, int> seeds) {
    generateAndDistributeMatrix(a, seeds.first);
    generateAndDistributeMatrix(b, seeds.second);
  }

  void generateAndDistributeMatrix(const Matrix &m, int seed) {
    if (mpi.isLeader()) {
      double *buffer = m.isA ? bufferA : bufferB;
      auto frags = grid.frags(m.isA);
      unsigned long xFrags = frags.first, yFrags = frags.second;

      for (unsigned long x = 0; x < xFrags; x++) {
        for (unsigned long y = 0; y < yFrags; y++) {
          unsigned long rank = grid.getRankForFragment(x, y, m.isA);
          int skewedRank = skewRank(rank, m.isA);

          if (skewedRank == mpi.myRank) {
            m.generateMatrixFragment(m.fragment.data, seed, x, y);
          } else {
            m.generateMatrixFragment(buffer, seed, x, y);
            MPI_CHECK(MPI_Send(buffer, m.fragment.size(), MPI_DOUBLE, skewedRank, FRAGMENT_TAG, grid.comm));
          }
        }
      }
    } else {
      unsigned long groupId = grid.pk(mpi.myRank).second;
      if (groupId < grid.replicationGroupSize(m.isA)) {
        MPI_CHECK(MPI_Recv(m.fragment.data,
                           m.fragment.size(),
                           MPI_DOUBLE,
                           MpiInfo::LEADER_RANK,
                           FRAGMENT_TAG,
                           grid.comm,
                           MPI_STATUS_IGNORE));
      }
    }
  }

  void replicateMatrices() {
    replicateMatrix(a);
    replicateMatrix(b);
  }

  void replicateMatrix(const Matrix &m) {
    unsigned long replicationFactor = grid.replicationFactor(m.isA);
    if (replicationFactor == 1) {
      return;
    }
    auto cannonData = grid.cannon(mpi.myRank);
    unsigned long cannonGroup = cannonData.first, cannonId = cannonData.second;

    MPI_CHECK(MPI_Bcast(m.fragment.data, m.fragment.size(), MPI_DOUBLE, 0, replicaComm));
  }

  void multiplyFragments() const {
    assert(a.fragment.width == b.fragment.height);

    for (unsigned long x = 0; x < result.fragment.width; x++) {
      for (unsigned long y = 0; y < result.fragment.height; y++)
        for (unsigned long k = 0; k < a.fragment.width; k++) {
          result.fragment.get(x, y) += a.fragment.get(k, y) * b.fragment.get(x, k);
        }
    }
  }

  void shiftMatrices() {
    std::vector<MPI_Request> reqs;
    reqs.reserve(4);

    shiftMatrix(a, reqs);
    shiftMatrix(b, reqs);

    MPI_CHECK(MPI_Waitall(reqs.size(), reqs.data(), MPI_STATUSES_IGNORE));
  }

  void shiftMatrix(const Matrix &m, std::vector<MPI_Request> &reqs) {
    unsigned long cannonId = grid.cannon(mpi.myRank).second;
    unsigned long cannonDim = grid.cannonDim();

    unsigned long col = cannonId / cannonDim;
    unsigned long row = cannonId % cannonDim;
    unsigned long dstCol, dstRow;
    if (m.isA) {
      dstCol = (col == 0 ? cannonDim : col) - 1;
      dstRow = row;
    } else {
      dstCol = col;
      dstRow = (row == 0 ? cannonDim : row) - 1;
    }
    unsigned long dst = dstCol * cannonDim + dstRow;

    unsigned long srcCol, srcRow;
    if (m.isA) {
      srcRow = row;
      srcCol = (col + 1) % cannonDim;
    } else {
      srcCol = col;
      srcRow = (row + 1) % cannonDim;
    }
    unsigned long src = srcCol * cannonDim + srcRow;

    if (src == mpi.myRank && dst == mpi.myRank) {
      return;
    }

    double *buffer = m.isA ? bufferA : bufferB;

    memcpy(buffer, m.fragment.data, m.fragment.size() * sizeof(double));

    MPI_Request send, recv;

    int tag = m.isA ? SHIFT_TAG_A : SHIFT_TAG_B;

    MPI_CHECK(MPI_Isend(buffer, m.fragment.size(), MPI_DOUBLE, src, tag, cannonComm, &send));
    MPI_CHECK(MPI_Irecv(m.fragment.data, m.fragment.size(), MPI_DOUBLE, dst, tag, cannonComm, &recv));

    reqs.push_back(send);
    reqs.push_back(recv);
  }

  void reduceResults() {
    unsigned long pkGroup = grid.pk(mpi.myRank).first;

    for (unsigned long diff = 1; diff < grid.dim.k; diff *= 2) {
      if (pkGroup % (diff * 2) == 0) {
        if (mpi.myRank + diff * grid.pkGroupSize() >= grid.size()) {
          continue;
        }
        MPI_CHECK(MPI_Recv(bufferResult,
                           result.fragment.size(),
                           MPI_DOUBLE,
                           pkGroup + diff,
                           RESULTS_TAG,
                           resultsComm,
                           MPI_STATUS_IGNORE));

        for (unsigned long i = 0; i < result.fragment.size(); i++) {
          result.fragment.data[i] += bufferResult[i];
        }
      } else {
        MPI_CHECK(MPI_Send(result.fragment.data,
                           result.fragment.size(),
                           MPI_DOUBLE,
                           pkGroup - diff,
                           RESULTS_TAG,
                           resultsComm));
        break;
      }
    }
  }

 public:
  CA3DMM(MpiInfo _mpi, Grid _grid, Dim3 _matrices_dim)
      : mpi(_mpi),
        grid(std::move(_grid)),
        a(createA(_matrices_dim)),
        b(createB(_matrices_dim)),
        pkComm(createPkComm()),
        cannonComm(createCannonComm()),
        resultsComm(createResultsComm()),
        replicaComm(createReplicaComm()),
        bufferA(new double[a.fragment.size()]),
        bufferB(new double[b.fragment.size()]),
        result(createResult(_matrices_dim)),
        bufferResult(new double[result.fragment.size()]) {}

  ~CA3DMM() {
    MPI_CHECK(MPI_Comm_free(&resultsComm));
    MPI_CHECK(MPI_Comm_free(&cannonComm));
    MPI_CHECK(MPI_Comm_free(&replicaComm));
    MPI_CHECK(MPI_Comm_free(&pkComm));

    delete[] bufferA;
    delete[] bufferB;
    delete[] bufferResult;
  }

  void run(std::pair<int, int> seeds) {
    result.fragment.clean();

    generateAndDistributeMatrices(seeds);
    replicateMatrices();

    for (unsigned long i = 0; i < grid.cannonDim(); i++) {
      multiplyFragments();
      if (i != grid.cannonDim() - 1) {
        // Last shift is redundant.
        shiftMatrices();
      }
    }

    reduceResults();
  }

  void printResult() {
    if (grid.pk(mpi.myRank).first != 0) {
      // Results are reduced in the first pk.
      return;
    }

    if (mpi.isLeader()) {
      std::cout << result.height << " " << result.width << std::endl;
      for (unsigned long row = 0; row < result.height; row++) {
        unsigned long col = 0;
        for (unsigned long fragX = 0; fragX < grid.dim.n; fragX++) {
          int rank = grid.getRankForResultFragment(fragX, row / result.fragment.height);

          if (rank == mpi.myRank) {
            result.fragment.copyRow(bufferResult, row);
          } else {
            MPI_CHECK(MPI_Recv(bufferResult,
                               result.fragment.width,
                               MPI_DOUBLE,
                               rank,
                               PRINT_TAG,
                               pkComm,
                               MPI_STATUS_IGNORE));
          }

          for (unsigned long x = 0; x < result.fragment.width && col < result.width; x++, col++) {
            std::cout << bufferResult[x] << " ";
          }
        }
        std::cout << std::endl;
      }
    } else {
      unsigned long globalRowStart = result.fragment.y * result.fragment.height;

      for (unsigned long row = 0; row < result.fragment.height && globalRowStart + row < result.height; row++) {
        result.fragment.copyRow(bufferResult, row);

        MPI_CHECK(MPI_Send(bufferResult,
                           result.fragment.width,
                           MPI_DOUBLE,
                           MpiInfo::LEADER_RANK,
                           PRINT_TAG,
                           pkComm));
      }
    }
  }

  void printOverGe(double ge) {
    if (grid.pk(mpi.myRank).first != 0) {
      // Summed results are in first group.
      return;
    }

    unsigned long long overGe = result.overGeInFragment(ge);

    if (mpi.isLeader()) {
      unsigned long long counts[grid.pkGroupSize()];

      MPI_CHECK(MPI_Gather(&overGe,
                           1,
                           MPI_UNSIGNED_LONG_LONG,
                           counts,
                           1,
                           MPI_UNSIGNED_LONG_LONG,
                           MpiInfo::LEADER_RANK,
                           pkComm));

      unsigned long long sum = 0;
      for (auto const el : counts) {
        sum += el;
      }

      std::cout << sum << std::endl;
    } else {
      MPI_CHECK(MPI_Gather(&overGe,
                           1,
                           MPI_UNSIGNED_LONG_LONG,
                           nullptr,
                           0,
                           MPI_DATATYPE_NULL,
                           MpiInfo::LEADER_RANK,
                           pkComm));
    }
  }
};

void ca3dmm(const Dim3 &matrix_dimensions,
            const MpiInfo &info,
            const std::vector<std::pair<int, int>> &seeds,
            bool printMatrix,
            const double *ge) {

  Dim3 gridDim = choose_grid_dimensions(matrix_dimensions, info.numProcesses).first;
  Grid grid{gridDim, info.myRank};

  if (grid.comm == MPI_COMM_NULL) {
    // I am useless, goodbye!
    return;
  }

  CA3DMM ca_3_dmm{info, grid, matrix_dimensions};

  for (const auto &seed : seeds) {
    ca_3_dmm.run(seed);
    if (ge != nullptr) {
      ca_3_dmm.printOverGe(*ge);
    } else if (printMatrix) {
      ca_3_dmm.printResult();
    }
  }
}

std::vector<std::pair<int, int>> parse_seeds(char *arg) {
  std::vector<std::pair<int, int>> res;

  size_t totalPos = 0;
  for (;;) {
    size_t pos;
    int seedA = std::stoi(arg + totalPos, &pos);
    totalPos += pos;
    if (arg[totalPos] != ',') {
      throw std::runtime_error("Invalid seeds format: " + std::string(arg) + ".");
    }
    totalPos++;
    int seedB = std::stoi(arg + totalPos, &pos);
    totalPos += pos;
    res.emplace_back(seedA, seedB);

    if (arg[totalPos] == ',') {
      totalPos++;
    } else if (arg[totalPos] == '\0') {
      break;
    } else {
      throw std::runtime_error("Invalid seeds format: " + std::string(arg) + ".");
    }
  }

  return res;
}

struct MPIGuard {
  MPIGuard(int *pargc, char ***pargv) {
    MPI_Init(pargc, pargv);
  }

  ~MPIGuard() {
    MPI_Finalize();
  }
};

int main(int argc, char **argv) {
  MpiInfo info;

  MPIGuard mpiGuard{&argc, &argv};

  MPI_Comm_size(MPI_COMM_WORLD, &info.numProcesses);
  MPI_Comm_rank(MPI_COMM_WORLD, &info.myRank);

  MPI_Comm_set_errhandler(MPI_COMM_WORLD, MPI_ERRORS_RETURN);

  if (argc < 6) {
    std::cerr << "Usage: " << argv[0] << " n m k -s seeds [-g ge_value] [-v]" << std::endl;
    return 1;
  }
  Dim3 matrixDimensions;
  std::vector<std::pair<int, int>> seeds;
  bool printMatrix = false;
  double ge;
  double *gePtr = nullptr;

  matrixDimensions.n = std::stoul(argv[1]);
  matrixDimensions.m = std::stoul(argv[2]);
  matrixDimensions.k = std::stoul(argv[3]);

  int c;
  while ((c = getopt(argc, argv, "s:g:v")) != -1) {
    switch (c) {
      case 's':seeds = parse_seeds(optarg);
        break;
      case 'g':ge = std::stod(optarg);
        gePtr = &ge;
        break;
      case 'v':printMatrix = true;
        break;
      default:std::cerr << "Usage: " << argv[0] << " n m k -s seeds [-g ge_value] [-v]" << std::endl;
        return 1;
    }
  }

  if (seeds.empty()) {
    std::cerr << "No seeds provided." << std::endl;
    return 1;
  }

  if (printMatrix && gePtr != nullptr) {
    std::cerr << "Options '-g' and '-v' are mutually exclusive." << std::endl;
    return 1;
  }

  ca3dmm(matrixDimensions, info, seeds, printMatrix, gePtr);

  return 0;
}
