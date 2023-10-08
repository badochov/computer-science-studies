#include <vector>
#include <algorithm>
#include <cmath>
#include <limits>

#include "types.hpp"
#include "params.hpp"

static constexpr double L = 0.95;

static unsigned long get_score(const Dim3 &dim, const Dim3 &p) {
  if (std::max(p.m, p.n) % std::min(p.n, p.m) != 0) {
    return std::numeric_limits<unsigned long>::max();
  }

  return p.m * dim.k * dim.n + p.k * dim.m * dim.n + p.n * dim.m * dim.k;
}

static std::vector<unsigned long> get_dividers(unsigned long n) {
  std::vector<unsigned long> dividers;

  for (unsigned long i = 2; i <= n && n != 1; i++) {
    while (n % i == 0) {
      dividers.push_back(i);
      n /= i;
    }
  }

  return dividers;
}

class CombinationsIterator {
  const std::vector<unsigned long> &dividers;
  size_t iteration = 0;

 public:
  explicit CombinationsIterator(const std::vector<unsigned long> &div)
      : dividers(div), iteration(static_cast<unsigned long>(pow(3, static_cast<double>(dividers.size())))) {

  }

  bool next() {
    return iteration--;
  }

  Dim3 get() {
    Dim3 p{1, 1, 1};

    size_t it = iteration;
    for (const auto divider : dividers) {
      switch (it % 3) {
        case 0:p.m *= divider;
          break;
        case 1:p.k *= divider;
          break;
        case 2:p.n *= divider;
          break;
      }

      it /= 3;
    }

    return p;
  }
};

static std::pair<Dim3, unsigned long> get_best_grid(const Dim3 &dim, unsigned long processes) {
  unsigned long current_min = std::numeric_limits<unsigned long>::max();
  Dim3 res{};

  std::vector<unsigned long> dividers = get_dividers(processes);
  CombinationsIterator it(dividers);

  while (it.next()) {
    Dim3 p = it.get();

    unsigned long score = get_score(dim, p);
    if (score < current_min) {
      current_min = score;
      res = p;
    }
  }

  return {res, current_min};
}

static unsigned long min_num_processes(unsigned long numProcesses) {
  return static_cast<unsigned long>(std::ceil(L * static_cast<double>(numProcesses)));
}

std::pair<Dim3, unsigned long> choose_grid_dimensions(const Dim3 &dim, unsigned long numProcesses) {
  unsigned long global_min = std::numeric_limits<unsigned long>::max();
  Dim3 grid{};

  for (unsigned long p = min_num_processes(numProcesses); p <= numProcesses; p++) {
    auto curr = get_best_grid(dim, p);
    Dim3 curr_grid = curr.first;
    unsigned long curr_min = curr.second;

    if (curr_min < global_min) {
      global_min = curr_min;
      grid = curr_grid;
    }
  }

  return {grid, global_min};
}