#ifndef CACTI_THREAD_POOL_H
#define CACTI_THREAD_POOL_H

#include <pthread.h>
#include <stdbool.h>

typedef struct pool *pool_t;


typedef void(*fn_t)(void *);


typedef struct {
    fn_t fn;
    void *arg;
} job_t;


pool_t pool_create(size_t size, size_t job_queue_size);

bool pool_add_job(pool_t pool, fn_t fn, void *arg);

void pool_join(pool_t pool);

void pool_end(pool_t pool);

void pool_destroy(pool_t pool);

#endif //CACTI_THREAD_POOL_H
