#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include "err.h"

#define BUFFER_SIZE   100000
#define PORT_NUM     10001

int main(int argc, char *argv[]) {
    int sock;
    int flags, sflags;
    struct sockaddr_in server_address;
    struct sockaddr_in client_address;

    FILE *f = fopen("rec.txt", "w");

    char buffer[BUFFER_SIZE];
    socklen_t snda_len, rcva_len;
    ssize_t len, snd_len;

    sock = socket(AF_INET, SOCK_DGRAM, 0); // creating IPv4 UDP socket
    if (sock < 0)
        syserr("socket");
    // after socket() call; we should close(sock) on any execution path;

    server_address.sin_family = AF_INET; // IPv4
    server_address.sin_addr.s_addr = htonl(INADDR_ANY); // listening on all interfaces
    server_address.sin_port = htons(PORT_NUM); // default port for receiving is PORT_NUM

    // bind the socket to a concrete address
    if (bind(sock, (struct sockaddr *) &server_address,
             (socklen_t) sizeof(server_address)) < 0)
        syserr("bind");

    snda_len = (socklen_t) sizeof(client_address);
    for (;;) {
        do {
            rcva_len = (socklen_t) sizeof(client_address);
            flags = 0; // we do not request anything special
            len = recvfrom(sock, buffer, sizeof(buffer), flags,
                           (struct sockaddr *) &client_address, &rcva_len);
            if (len < 0)
                syserr("error on datagram from client socket");
            else {
                (void) printf("read from socket: %zd bytes: %.*s\n", len,
                              (int) len, buffer);

                fwrite(buffer, 1, len, f);
                fclose(f);
                fopen("rec.txt", "acma  ");

                sprintf(buffer, "%zd", len);
                int res_len = 0;
                int len_h = len;


                while (len_h > 0) {
                    ++res_len;
                    len_h /= 10;
                }

                sflags = 0;
                snd_len = sendto(sock, buffer, (size_t) res_len, sflags,
                                 (struct sockaddr *) &client_address, snda_len);
                if (snd_len != res_len)
                    syserr("error on sending datagram to client socket");
            }
        } while (len > 0);
        (void) printf("finished exchange\n");
    }

    fclose(f);

    if (close(sock) == -1) { //very rare errors can occur here, but then
        syserr("close"); //it's healthy to do the check
    }

    return 0;
}
