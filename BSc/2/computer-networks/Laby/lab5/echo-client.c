#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>

#include "err.h"

#define BUFFER_SIZE 100000


#define PORT "42069"
#define HOST "students.mimuw.edu.pl"

int main(int argc, char *argv[])
{
  int sock;
  struct addrinfo addr_hints;
  struct addrinfo *addr_result;

  int i, err;
  char buffer[BUFFER_SIZE];
  ssize_t len;

  if (argc != 3) {
    fatal("Usage: %s number_of_requests request_size\n", argv[0]);
  }

  // 'converting' host/port in string to struct addrinfo
  memset(&addr_hints, 0, sizeof(struct addrinfo));
  addr_hints.ai_family = AF_INET; // IPv4
  addr_hints.ai_socktype = SOCK_STREAM;
  addr_hints.ai_protocol = IPPROTO_TCP;
  err = getaddrinfo(HOST, PORT, &addr_hints, &addr_result);
  if (err == EAI_SYSTEM) { // system error
    syserr("getaddrinfo: %s", gai_strerror(err));
  }
  else if (err != 0) { // other error (host not found, etc.)
    fatal("getaddrinfo: %s", gai_strerror(err));
  }

  // initialize socket according to getaddrinfo results
  sock = socket(addr_result->ai_family, addr_result->ai_socktype, addr_result->ai_protocol);
  if (sock < 0)
    syserr("socket");

  // connect socket to the server
  if (connect(sock, addr_result->ai_addr, addr_result->ai_addrlen) < 0)
    syserr("connect");

  freeaddrinfo(addr_result);

    long n = strtol(argv[1], NULL, 0);
    long k = strtol(argv[2], NULL, 0);

    char msg[k];

    for (i = 0; i < n; i++) {
        len = k;
        memset(msg, 'a' + (i % 24), len);
        (void) printf("sending to socket: %zd bytes\n", len);
        if (write(sock, buffer, len) != len) {
            syserr("partial / failed write");
        }
    }


  (void) close(sock); // socket would be closed anyway when the program ends

  return 0;
}
