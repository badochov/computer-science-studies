#ifndef CACTI_UTIL_H
#define CACTI_UTIL_H

#include <signal.h>
#include <pthread.h>


void disable_all_signals(sigset_t *old);

void set_mask(sigset_t *mask, sigset_t *old);

void enable_signal(int signo, sigset_t *old);

void disable_signal(int signo, sigset_t *old);

size_t min(size_t a, size_t b);

size_t max(size_t a, size_t b);


#endif //CACTI_UTIL_H
