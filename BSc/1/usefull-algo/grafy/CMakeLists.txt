cmake_minimum_required(VERSION 3.15)
project(grafy)

set(CMAKE_CXX_STANDARD 17)

include_directories(graph)
add_executable(graf graph/graph.h main.cpp)
#target_include_directories(graf PUBLIC ${graph})

add_executable(dijkstra graph/graph.h dijkstra/main.cpp)
#target_include_directories(dijkstra PRIVATE ${graph})