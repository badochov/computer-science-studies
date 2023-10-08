#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <errno.h>
#include <unistd.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "err.h"

#define LINE_SIZE 1024

size_t file_size_sum = 0;

pthread_mutex_t mutex;

void *handle_connection(void *s_ptr) {
  int ret, s;
  socklen_t len;
  char line[LINE_SIZE + 1], peername[LINE_SIZE + 1], peeraddr[LINE_SIZE + 1];
  struct sockaddr_in addr;

  s = *(int *) s_ptr;
  free(s_ptr);

  len = sizeof(addr);

  /* Któż to do nas dzwoni (adres)?  */
  ret = getpeername(s, (struct sockaddr *) &addr, &len);
  if (ret == -1)
    syserr("getsockname");

  inet_ntop(AF_INET, &addr.sin_addr, peeraddr, LINE_SIZE);
  snprintf(peername, 2 * LINE_SIZE, "%s:%d", peeraddr, ntohs(addr.sin_port));

  size_t file_size;
  char *filename = NULL;

  FILE *f = fdopen(s, "r");

  size_t n = 0;
  char *tmp = NULL;
  getline(&filename, &n, f);
  n = 0;
  getline(&tmp, &n, f);

  filename[strlen(filename) - 1] = '\0';



  file_size = strtoll(tmp, NULL, 10);

  printf("new client %s size=%zu file=%s", peername, file_size, filename);

  sleep(1);

  FILE *incoming_file = fopen(filename, "w");

  for (;;) {
    memset(line, 0, sizeof(line));
    ret = fread(line,1, sizeof(line) - 1, f);
    if (ret == -1)
      syserr("read");
    else if (ret == 0)
      break;
    fputs(line, incoming_file);
  }

  fclose(incoming_file);

  fclose(f);
  close(s);

  pthread_mutex_lock(&mutex);
  file_size_sum += file_size;
  pthread_mutex_unlock(&mutex);

  printf("client %s has sent its file of size=%zu\n"
         "total size of uploaded files %zu\n",
         peername,
         file_size,
         file_size_sum);

  return 0;
}

int main(int argc, char **argv) {
  int ear, rc;
  socklen_t len;
  struct sockaddr_in server;

  if (argc != 2)
    syserr("Usage: %s port", argv[2]);

  int port = atoi(argv[1]);

  /* Tworzymy gniazdko */
  ear = socket(PF_INET, SOCK_STREAM, 0);
  if (ear == -1)
    syserr("socket");

  /* Podłączamy do centrali */
  server.sin_family = AF_INET;
  server.sin_addr.s_addr = htonl(INADDR_ANY);
  server.sin_port = htons(port);
  rc = bind(ear, (struct sockaddr *) &server, sizeof(server));
  if (rc == -1)
    syserr("bind");


  /* Każdy chce wiedzieć jaki to port */
  len = (socklen_t) sizeof(server);
  rc = getsockname(ear, (struct sockaddr *) &server, &len);
  if (rc == -1)
    syserr("getsockname");

  printf("Listening at port %d\n", (int) ntohs(server.sin_port));

  rc = listen(ear, 5);
  if (rc == -1)
    syserr("listen");

  pthread_mutex_init(&mutex, NULL);

  /* No i do pracy */
  for (;;) {
    int msgsock;
    int *con;
    pthread_t t;

    msgsock = accept(ear, (struct sockaddr *) NULL, NULL);
    if (msgsock == -1) {
      syserr("accept");
    }

    /* Tylko dla tego wątku */
    con = malloc(sizeof(int));
    if (!con) {
      syserr("malloc");
    }
    *con = msgsock;

    rc = pthread_create(&t, 0, handle_connection, con);
    if (rc == -1) {
      syserr("pthread_create");
    }

    /* No przecież nie będę na niego czekał ... */
    rc = pthread_detach(t);
    if (rc == -1) {
      syserr("pthread_detach");
    }
  }
}

