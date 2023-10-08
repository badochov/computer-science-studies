#include "thread_pool.h"
#define QUEUE_EL_T job_t
#include "queue.c"
#undef QUEUE_EL_T
