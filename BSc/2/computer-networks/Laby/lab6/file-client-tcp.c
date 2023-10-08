#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "err.h"

int main(int argc, char *argv[]) {
  int rc;
  int sock;
  struct addrinfo addr_hints, *addr_result;

  /* Kontrola dokumentów ... */
  if (argc != 4) {
    fatal("Usage: %s host port file_name", argv[0]);
  }

  sock = socket(PF_INET, SOCK_STREAM, IPPROTO_TCP);
  if (sock < 0) {
    syserr("socket");
  }

  /* Trzeba się dowiedzieć o adres internetowy serwera. */
  memset(&addr_hints, 0, sizeof(struct addrinfo));
  addr_hints.ai_flags = 0;
  addr_hints.ai_family = AF_INET;
  addr_hints.ai_socktype = SOCK_STREAM;
  addr_hints.ai_protocol = IPPROTO_TCP;

  rc = getaddrinfo(argv[1], argv[2], &addr_hints, &addr_result);
  if (rc != 0) {
    fprintf(stderr, "rc=%d\n", rc);
    syserr("getaddrinfo: %s", gai_strerror(rc));
  }

  if (connect(sock, addr_result->ai_addr, addr_result->ai_addrlen) != 0) {
    syserr("connect");
  }
  freeaddrinfo(addr_result);

  FILE *f = fopen(argv[3], "r");

  fseek(f, 0, SEEK_END);
  size_t size = ftell(f);
  rewind(f);

  char *file_content = malloc(size + 1);

  fread(file_content, 1, size, f);

  fclose(f);

  if (write(sock, argv[3], strlen(argv[3])) < 0)
    syserr("writing on stream socket");

  char size_str[21] = "";

  sprintf(size_str, "\n%zu\n", size);

  if (write(sock, size_str, strlen(size_str)) < 0)
    syserr("writing on stream socket");

  if (write(sock, file_content, size) < 0)
    syserr("writing on stream socket");

  if (close(sock) < 0)
    syserr("closing stream socket");

  free(file_content);

  return 0;
}

