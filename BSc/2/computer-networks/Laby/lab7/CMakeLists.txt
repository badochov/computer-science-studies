project("Lab7")

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
add_executable(poll_server poll_server.c)

add_executable(poll_server_telnet poll_server_telnet.c)

add_executable(nonblocking_client nonblocking/echo-client-timeout.c)
add_executable(nonblocking_poll_server nonblocking/echo-server-nonblocking.c)

install(TARGETS DESTINATION .)
