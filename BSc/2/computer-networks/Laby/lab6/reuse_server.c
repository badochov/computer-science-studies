#include <stdio.h>
#include <stdlib.h>

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/ip.h>
#include <netinet/tcp.h>
#include <linux/net.h>
#include <unistd.h>
#include <arpa/inet.h>

#include "err.h"

#define BUFSIZE 1024
#define QUEUELEN 5

int main(int argc, char *argv[])
{
    if (argc != 2)
        fatal("Usage: %s port_number", argv[0]);

    int port = atoi(argv[1]);
    if (port < 1 || port > 65535)
        fatal("Usage: %s port_number", argv[0]);

    int sock = socket(PF_INET, SOCK_STREAM, 0);
    if (sock == -1)
        syserr("opening stream socket");

    // Tu ustawiamy opcję gniazda, SO_REUSEPORT trzeba koniecznie ustawić
    // przed wywołaniem bind()
    int optval = 1;
    if (setsockopt(sock, SOL_SOCKET, SO_REUSEPORT, &optval, sizeof optval) == -1)
        syserr("setting SO_REUSEPORT option");

    struct sockaddr_in saddr;
    saddr.sin_family = AF_INET;
    saddr.sin_addr.s_addr = htonl(INADDR_ANY);
    saddr.sin_port = htons(port);

    if (bind(sock, (struct sockaddr *)&saddr, sizeof saddr) == -1)
        syserr("bind failed");
    if (listen(sock, QUEUELEN) == -1)
        syserr("listen failed on bound socket");

    printf("Process %d starts serving clients on socket %d.\n", getpid(), port);

    do
    {
        struct sockaddr_in peer;
        unsigned peer_addr_size = sizeof peer;

        int cs = accept(sock, (struct sockaddr *)&peer, &peer_addr_size);
        if (cs == -1)
            syserr("accept failed");

        char addr_buf[64];
        inet_ntop(AF_INET, (void *)&peer.sin_addr, addr_buf, 64);
        printf("[%d] Connection from %s:%d.\n", getpid(), addr_buf, ntohs(peer.sin_port));

        int cnt;
        do
        {
            char buf[BUFSIZE];

            cnt = read(cs, buf, BUFSIZE);
            if (cnt == -1)
                syserr("read failed");
            else if (cnt == 0)
                printf("[%d] closing connection.\n", getpid());
            else
                printf("[%d] -> \"%.*s\"\n", getpid(), (int)cnt, buf);
        } while (cnt > 0);

        if (close(cs) == -1)
            syserr("close failed");
    } while (1);

    return 0;
}
