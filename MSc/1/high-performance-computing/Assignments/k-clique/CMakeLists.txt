cmake_minimum_required(VERSION 3.10)
project(kcliques CUDA)

set(CMAKE_CUDA_STANDARD 20)

if (CMAKE_BUILD_TYPE STREQUAL "Debug")
    set(CMAKE_CUDA_FLAGS ${CMAKE_CUDA_FLAGS} "-g -G -O0")  # enable cuda-gdb
else ()
    set(CMAKE_CUDA_FLAGS "${CMAKE_CUDA_FLAGS} -O3 -arch=sm_70 -gencode=arch=compute_70,code=sm_70 -gencode=arch=compute_70,code=compute_70")
endif ()

add_executable(kcliques kcliques.cu)

set_target_properties(kcliques PROPERTIES
        CUDA_SEPARABLE_COMPILATION ON)

set_target_properties(kcliques PROPERTIES CUDA_ARCHITECTURES "61;70")
