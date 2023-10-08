#include "thread_pool.h"
#include <stdlib.h>


#include <stdio.h>


#include "job_queue.h"


struct pool {
    pthread_cond_t cond;
    pthread_mutex_t mutex;
    bool end;
    job_t__queue_t job_queue;
    size_t pool_size;
    pthread_t *threads;
};


static void *thread(void *arg) {

    pool_t pool = arg;
    while (1) {
        pthread_mutex_lock(&pool->mutex);
        while (!pool->end && job_t__queue_size(pool->job_queue) == 0) {
            pthread_cond_wait(&pool->cond, &pool->mutex);
        }
        if (pool->end) {
            pthread_mutex_unlock(&pool->mutex);
            break;
        } else {
            job_t job = job_t__queue_pop(pool->job_queue);
            if (job_t__queue_size(pool->job_queue) > 0) {
                pthread_cond_signal(&pool->cond);
            }
            pthread_mutex_unlock(&pool->mutex);
            job.fn(job.arg);
        }
    }

    return NULL;
}

pool_t pool_create(size_t size, size_t job_queue_size) {
    pool_t pool = calloc(1, sizeof(struct pool));
    if (pool == NULL) {
        return NULL;
    }
    pthread_cond_init(&pool->cond, NULL);
    pthread_mutex_init(&pool->mutex, NULL);
    pool->job_queue = job_t__queue_init(job_queue_size);
    pool->threads = calloc(size, sizeof(pthread_t));
    pool->pool_size = size;
    for (size_t i = 0; i < size; i++) {
        if (pthread_create(&pool->threads[i], NULL, thread, pool)) {
            return NULL;
        }
    }

    return pool;
}

bool pool_add_job(pool_t pool, fn_t fn, void *arg) {
    pthread_mutex_lock(&pool->mutex);
    bool res = job_t__queue_add(pool->job_queue, (job_t) {fn, arg});

    pthread_cond_signal(&pool->cond);
    pthread_mutex_unlock(&pool->mutex);
    return res;
}

void pool_join(pool_t pool) {
    if (!pool) {
        return;
    }
    if (!pool->end) {
        for (size_t i = 0; i < pool->pool_size; i++) {
            pthread_join(pool->threads[i], NULL);
        }
    }
}

void pool_end(pool_t pool) {
    pthread_mutex_lock(&pool->mutex);
    pool->end = true;
    pthread_cond_broadcast(&pool->cond);
    pthread_mutex_unlock(&pool->mutex);
}

void pool_destroy(pool_t pool) {
    if (pool == NULL) {
        return;
    }
    pthread_cond_destroy(&pool->cond);
    pthread_mutex_destroy(&pool->mutex);
    job_t__queue_destroy(pool->job_queue);
    free(pool->threads);

    free(pool);
}