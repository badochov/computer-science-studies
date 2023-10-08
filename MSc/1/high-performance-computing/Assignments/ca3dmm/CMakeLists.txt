cmake_minimum_required(VERSION 3.10)
SET(CMAKE_CXX_COMPILER mpicxx)   # use CC for okeanos

project(ca3dmm CXX)

set(CMAKE_CXX_STANDARD 14)

include_directories(.)

add_executable(ca3dmm
        main.cpp types.hpp)

add_executable(param_test
        param_test.cpp)

add_library(densematgen
        densematgen.cpp
        densematgen.h)

add_library(params
        params.cpp
        params.hpp
        types.hpp)

find_package(MPI REQUIRED)

target_link_libraries(ca3dmm
        PRIVATE MPI::MPI_CXX
        PRIVATE densematgen
        PRIVATE params)

target_link_libraries(param_test
        PRIVATE params)
