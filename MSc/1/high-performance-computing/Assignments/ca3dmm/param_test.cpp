#include <iostream>
#include <fstream>

#include "params.hpp"
#include "types.hpp"

int main(int argc, char **argv) {
  if (argc != 2) {
    std::cerr << "Usage: " << argv[0] << "<test_file>" << std::endl;
    return 1;
  }
  unsigned long numProcesses;
  Dim3 in, out;
  std::ifstream fin{argv[1]};
  bool ok = true;
  unsigned long work;

  while (fin >> numProcesses >> in.n >> in.m >> in.k >> out.n >> out.m >> out.k >> work) {
    auto r = choose_grid_dimensions(in, numProcesses);
    Dim3 res = r.first;
    unsigned long res_work = r.second;

    if (work != res_work) {
      ok = false;
      std::cerr << "For " << numProcesses << " processes on dim3(" << in.n << ", " << in.m << ", " << in.k <<
                "), got: dim3(" << res.n << ", " << res.m << ", " << res.k << ")[" << res_work << "], expected: dim3("
                << out.n << ", "
                << out.m
                << ", " << out.k << ")[" << work << "]" << std::endl;
    }
  }

  return ok ? 0 : 1;
}