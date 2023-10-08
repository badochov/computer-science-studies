cmake_minimum_required(VERSION 3.10)
project(converter)

set(CMAKE_CXX_STANDARD 20)

include(FindPkgConfig)

pkg_check_modules (CAPSTONE REQUIRED capstone)
pkg_check_modules (KEYSTONE REQUIRED keystone)

set(CMAKE_CXX_FLAGS  "${CMAKE_CXX_FLAGS}")

add_executable(converter main.cpp)

target_link_libraries(converter ${CAPSTONE_LIBRARIES})
target_link_libraries(converter ${KEYSTONE_LIBRARIES})

target_include_directories(converter PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/ELFIO)
