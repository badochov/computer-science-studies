#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include "err.h"

void syserr(const char *fmt) {
    fprintf(stderr, "ERROR: %s (%d; %s)\n", fmt, errno, strerror(errno));
    exit(1);
}

void fatal(const char *fmt) {
    fprintf(stderr, "ERROR: %s\n", fmt);
    exit(1);
}
