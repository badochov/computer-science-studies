#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <linux/if_alg.h>
#include <linux/socket.h>

// #define DIGEST_SZ 32
// #define DIGEST_NAME "sha256"

#define DIGEST_SZ 4
#define DIGEST_NAME "adler32"

int main(int argc, char **argv) {
	struct sockaddr_alg sa = {
		.salg_family = AF_ALG,
		.salg_type = "hash",
		.salg_name = DIGEST_NAME,
	};
	int fd = open(argv[1], O_RDONLY);
	if (fd == -1) {
		perror("open");
		return 1;
	}
	struct stat stat;
	if (fstat(fd, &stat)) {
		perror("fstat");
		return 1;
	}
	void *ptr = mmap(0, stat.st_size, PROT_READ, MAP_SHARED, fd, 0);
	if (ptr == MAP_FAILED) {
		perror("mmap");
		return 1;
	}
	int sfd = socket(AF_ALG, SOCK_SEQPACKET, 0);
	if (sfd == -1) {
		perror("socket");
		return 1;
	}
	if (bind(sfd, (struct sockaddr *)&sa, sizeof(sa))) {
		perror("bind");
		return 1;
	}
	int opfd = accept(sfd, NULL, 0);
	if (opfd == -1) {
		perror("accept");
		return 1;
	}
	if (send(opfd, ptr, stat.st_size, 0) != stat.st_size) {
		perror("send");
		return 1;
	}
	unsigned char buf[DIGEST_SZ];
	if (read(opfd, buf, sizeof buf) != sizeof buf) {
		perror("read");
		return 1;
	}
	for (int i = 0; i < sizeof buf; i++)
		printf("%02x", buf[i]);
	printf("\n");
	return 0;
}
