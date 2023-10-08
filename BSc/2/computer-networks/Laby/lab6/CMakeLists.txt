project(lab6)

cmake_minimum_required (VERSION 3.1)

set(CMAKE_C_FLAGS "-g -Wall")

# http://stackoverflow.com/questions/10555706/
macro (add_executable _name)
  _add_executable(${ARGV})
  if (TARGET ${_name})
    target_link_libraries(${_name} err)
  endif()
endmacro()

add_library(err err.c)

add_executable(client client.c)
add_executable(thread_server thread_server.c)
target_link_libraries(thread_server pthread)
add_executable(proc_server proc_server.c)
add_executable(reuse_server reuse_server.c)
add_executable(preforked_server preforked_server.c)

add_executable(file-client-tcp file-client-tcp.c)
add_executable(file-server-tcp file-server-tcp.c)
target_link_libraries(file-server-tcp pthread)

install(TARGETS DESTINATION .)
