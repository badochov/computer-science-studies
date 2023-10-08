#include "util.h"

void disable_all_signals(sigset_t *old) {
    sigset_t set;
    sigfillset(&set);
    pthread_sigmask(SIG_BLOCK, &set, old);
}

void enable_all_signals(sigset_t *old) {
    sigset_t set;
    sigemptyset(&set);
    set_mask(&set, old);
}

void set_mask(sigset_t *mask, sigset_t *old) {
    pthread_sigmask(SIG_SETMASK, mask, old);
}

void enable_signal(int signo, sigset_t *old) {
    sigset_t set;
    sigemptyset(&set);
    sigaddset(&set, signo);
    sigprocmask(SIG_UNBLOCK, &set, old);
}

void disable_signal(int signo, sigset_t *old) {
    sigset_t set;
    sigemptyset(&set);
    sigaddset(&set, signo);
    sigprocmask(SIG_BLOCK, &set, old);
}

size_t min(size_t a, size_t b) {
    return a > b ? b : a;
}

size_t max(size_t a, size_t b) {
    return a > b ? a : b;
}
