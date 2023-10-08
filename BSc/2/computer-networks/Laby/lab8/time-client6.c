#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <sys/types.h>
#include <sys/socket.h>

#include <poll.h>
#include <stdio.h>

#include "err.h"

#define BSIZE         256
#define HOPS_VALUE    4
#define REPEAT_COUNT  3
#define SLEEP_TIME    1

int main(int argc, char *argv[]) {
  /* argumenty wywołania programu */
  char *remote_colon_address;
  in_port_t remote_port;

  /* zmienne i struktury opisujące gniazda */
  int sock, optval;
  struct sockaddr_in6 local_address;
  struct sockaddr_in6 remote_address;

  /* zmienne obsługujące komunikację */
  char buffer[BSIZE];
  size_t length;
  time_t time_buffer;
  int i;

  /* parsowanie argumentów programu */
  if (argc != 3)
    fatal("Usage: %s remote_address remote_port\n", argv[0]);
  remote_colon_address = argv[1];
  remote_port = (in_port_t) atoi(argv[2]);

  /* otwarcie gniazda */
  sock = socket(AF_INET6, SOCK_DGRAM, 0);
  if (sock < 0)
    syserr("socket");

  /* ustawienie liczby skoków  dla datagramów rozsyłanych do grupy */
  optval = HOPS_VALUE;
  if (setsockopt(sock, SOL_IPV6, IPV6_MULTICAST_HOPS, (void *) &optval, sizeof optval) < 0)
    syserr("setsockopt multicast hops");

  /* związanie z gniazdem adresu i portu odbiorcy, aby móc użyć write zamiast sendto */
  remote_address.sin6_family = AF_INET6;
  remote_address.sin6_port = htons(remote_port);
  if (inet_pton(AF_INET6, remote_colon_address, &remote_address.sin6_addr) < 0)
    syserr("inet_pton");

  struct pollfd client[1];

  socklen_t len;
  getsockname(sock, (struct sockaddr *) &local_address, &len);

  bind(sock, (struct sockaddr *) &local_address, len);

  for (i = 0; i < REPEAT_COUNT; ++i) {
    printf("Sending request %d\n", i + 1);

    if (sendto(sock, "GET_TIME", 8, 0, (struct sockaddr *) &remote_address, sizeof remote_address) != 8)
      syserr("write");

    client[0].revents = 0;
    client[0].fd = sock;
    client[0].events = POLLIN;

    /* Czekamy przez 3000 ms */
    int ret = poll(client, 1, 3000);
    if (ret > 0) {
      if ((client[0].revents & POLLIN)) {
        /* Przyjmuję nowe połączenie */
        char buff[420];

        struct sockaddr_in6 addr;
        socklen_t l = sizeof addr;
        int msg_len = recvfrom(sock, buff, 420, 0, (struct sockaddr *) &addr, &l);

        char peeraddr[42];
        inet_ntop(AF_INET6, &addr.sin6_addr, peeraddr, 42);
        printf("Response from: %s\n", peeraddr);
        printf("Received time: %.*s\n", msg_len, buff);

        break;

      }
    }
  }

  if (i == REPEAT_COUNT) {
    syserr("Timeout: unable to receive response.");
  }

  /* koniec */
  close(sock);
  exit(EXIT_SUCCESS);
}
