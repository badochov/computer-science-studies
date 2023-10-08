/*
 Ten program używa poll(), aby równocześnie obsługiwać wielu klientów
 bez tworzenia procesów ani wątków.
*/

#include <limits.h>
#include <poll.h>
#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <netdb.h>
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>
#include "err.h"

#define TRUE 1
#define FALSE 0
#define BUF_SIZE 1024

static int finish = FALSE;

/* Obsługa sygnału kończenia */
static void catch_int(int sig) {
  finish = TRUE;
  fprintf(stderr,
          "Signal %d catched. No new connections will be accepted.\n", sig);
}

int main(int argc, char **argv) {
  if (argc != 3) {
    fatal("Usage: %s client_port control_port", argv[0]);
  }

  int client_port = atoi(argv[1]);
  int control_port = atoi(argv[2]);

  struct pollfd client[_POSIX_OPEN_MAX];
  _Bool is_telnet[_POSIX_OPEN_MAX] = {0};
  struct sockaddr_in client_server;
  struct sockaddr_in telnet_server;
  char buf[BUF_SIZE];
  size_t length;
  ssize_t rval;
  int msgsock, activeClients, totalClients = 0, i, ret;
  struct sigaction action;
  sigset_t block_mask;

  fprintf(stderr, "_POSIX_OPEN_MAX = %d\n", _POSIX_OPEN_MAX);

  /* Po Ctrl-C kończymy */
  sigemptyset(&block_mask);
  action.sa_handler = catch_int;
  action.sa_mask = block_mask;
  action.sa_flags = SA_RESTART;

  if (sigaction(SIGINT, &action, 0) == -1)
    syserr("sigaction");

  /* Inicjujemy tablicę z gniazdkami klientów, client[0] to gniazdko centrali */
  for (i = 0; i < _POSIX_OPEN_MAX; ++i) {
    client[i].fd = -1;
    client[i].events = POLLIN;
    client[i].revents = 0;
  }
  activeClients = 0;

  /* Tworzymy gniazdko centrali */
  client[0].fd = socket(PF_INET, SOCK_STREAM, 0);
  if (client[0].fd == -1)
    syserr("Opening stream socket");

  /* Co do adresu nie jesteśmy wybredni */
  client_server.sin_family = AF_INET;
  client_server.sin_addr.s_addr = htonl(INADDR_ANY);
  client_server.sin_port = htons(client_port);
  if (bind(client[0].fd, (struct sockaddr *) &client_server,
           (socklen_t) sizeof(client_server)) == -1)
    syserr("Binding stream socket");

  /* Dowiedzmy się, jaki to port i obwieśćmy to światu */
  length = sizeof(client_server);
  if (getsockname(client[0].fd, (struct sockaddr *) &client_server,
                  (socklen_t *) &length) == -1)
    syserr("Getting socket name");
  printf("Socket port #%u\n", (unsigned) ntohs(client_server.sin_port));

  /* Zapraszamy klientów */
  if (listen(client[0].fd, 5) == -1)
    syserr("Starting to listen");




  /* Tworzymy gniazdko TELNETU */
  client[1].fd = socket(PF_INET, SOCK_STREAM, 0);
  if (client[1].fd == -1)
    syserr("Opening stream socket");

  /* Co do adresu nie jesteśmy wybredni */
  telnet_server.sin_family = AF_INET;
  telnet_server.sin_addr.s_addr = htonl(INADDR_ANY);
  telnet_server.sin_port = htons(control_port);
  if (bind(client[1].fd, (struct sockaddr *) &telnet_server,
           (socklen_t) sizeof(telnet_server)) == -1)
    syserr("Binding stream socket");

  /* Dowiedzmy się, jaki to port i obwieśćmy to światu */
  length = sizeof(telnet_server);
  if (getsockname(client[1].fd, (struct sockaddr *) &telnet_server,
                  (socklen_t *) &length) == -1)
    syserr("Getting socket name");
  printf("Control socket port #%u\n", (unsigned) ntohs(telnet_server.sin_port));

  /* Zapraszamy klientów */
  if (listen(client[1].fd, 5) == -1)
    syserr("Starting to listen");

  /* Do pracy */
  do {
    for (i = 0; i < _POSIX_OPEN_MAX; ++i)
      client[i].revents = 0;

    /* Po Ctrl-C zamykamy gniazdko centrali */
    if (finish == TRUE && client[0].fd >= 0) {
      if (close(client[0].fd) < 0)
        syserr("close");
      client[0].fd = -1;
    }

    /* Po Ctrl-C zamykamy gniazdko telentu */
    if (finish == TRUE && client[1].fd >= 0) {
      if (close(client[1].fd) < 0)
        syserr("close");
      client[1].fd = -1;
    }

    /* Czekamy przez 5000 ms */
    ret = poll(client, _POSIX_OPEN_MAX, 5000);
    if (ret == -1)
      if (errno == EINTR) fprintf(stderr, "Interrupted system call\n");
      else syserr("poll");
    else if (ret > 0) {
      if (finish == FALSE && (client[0].revents & POLLIN)) {
        /* Przyjmuję nowe połączenie */
        msgsock = accept(client[0].fd, (struct sockaddr *) 0, (socklen_t *) 0);
        if (msgsock == -1)
          syserr("accept");
        else {
          for (i = 2; i < _POSIX_OPEN_MAX; ++i) {
            if (client[i].fd == -1) {
              fprintf(stderr, "Received new connection (%d)\n", i);
              client[i].fd = msgsock;
              client[i].events = POLLIN;
              activeClients += 1;

              totalClients += 1;
              break;
            }
          }
          if (i >= _POSIX_OPEN_MAX) {
            fprintf(stderr, "Too many clients\n");
            if (close(msgsock) < 0)
              syserr("close");
          }
        }
      }

      if (finish == FALSE && (client[1].revents & POLLIN)) {
        /* Przyjmuję nowe połączenie */
        msgsock = accept(client[1].fd, (struct sockaddr *) 0, (socklen_t *) 0);
        if (msgsock == -1)
          syserr("accept");
        else {
          for (i = 2; i < _POSIX_OPEN_MAX; ++i) {
            if (client[i].fd == -1) {
              fprintf(stderr, "Received new connection (%d)\n", i);
              client[i].fd = msgsock;
              client[i].events = POLLIN;

              is_telnet[i] = 1;
              break;
            }
          }
          if (i >= _POSIX_OPEN_MAX) {
            fprintf(stderr, "Too many clients\n");
            if (close(msgsock) < 0)
              syserr("close");
          }
        }
      }

      for (i = 2; i < _POSIX_OPEN_MAX; ++i) {
        if (!is_telnet[i] && client[i].fd != -1 && (client[i].revents & (POLLIN | POLLERR))) {
          rval = read(client[i].fd, buf, BUF_SIZE);
          if (rval < 0) {
            fprintf(stderr, "Reading message (%d, %s)\n", errno, strerror(errno));
            if (close(client[i].fd) < 0)
              syserr("close");
            client[i].fd = -1;
            activeClients -= 1;
          } else if (rval == 0) {
            fprintf(stderr, "Ending connection\n");
            if (close(client[i].fd) < 0)
              syserr("close");
            client[i].fd = -1;
            activeClients -= 1;
          } else
            printf("-->%.*s\n", (int) rval, buf);
        }
      }

      for (i = 2; i < _POSIX_OPEN_MAX; ++i) {
        if (is_telnet[i] && client[i].fd != -1 && (client[i].revents & (POLLIN | POLLERR))) {
          rval = read(client[i].fd, buf, BUF_SIZE);
          if (rval < 0) {
            fprintf(stderr, "Reading message (%d, %s)\n", errno, strerror(errno));
            if (close(client[i].fd) < 0)
              syserr("close");
            client[i].fd = -1;
            is_telnet[i] = 0;
          } else if (rval == 0) {
            fprintf(stderr, "Ending connection\n");
            if (close(client[i].fd) < 0)
              syserr("close");
            client[i].fd = -1;
            is_telnet[i] = 0;
          } else {
            printf("-->%d %.*s\n", (int) rval, (int) rval, buf);

            if (strncmp("count\r\n", buf, rval) == 0) {
              char resp[69] = "";
              sprintf(resp, "Number of active clients: %d\nTotal number of clients: %d\n", activeClients, totalClients);
              if (write(client[i].fd, resp, strlen(resp)) < 0) {
                syserr("write failed");
              }
            }

            fprintf(stderr, "Ending connection\n");
            if (close(client[i].fd) < 0)
              syserr("close");
            client[i].fd = -1;
            is_telnet[i] = 0;
          }
        }
      }
    } else
      fprintf(stderr, "Do something else\n");
  } while (finish == FALSE || activeClients > 0);

  if (client[0].fd >= 0)
    if (close(client[0].fd) < 0)
      syserr("Closing main socket");

  if (client[1].fd >= 0)
    if (close(client[1].fd) < 0)
      syserr("Closing telnet socket");
  exit(EXIT_SUCCESS);
}
