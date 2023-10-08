#ifndef CA3DMM__PARAMS_HPP_
#define CA3DMM__PARAMS_HPP_

#include "types.hpp"

std::pair<Dim3, unsigned long> choose_grid_dimensions(const Dim3 &dim, unsigned long numProcesses);

#endif //CA3DMM__PARAMS_HPP_
