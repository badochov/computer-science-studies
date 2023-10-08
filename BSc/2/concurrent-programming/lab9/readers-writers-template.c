#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>
#include "err.h"

#define READERS 3
#define WRITERS 2
#define NAP 2
#define BSIZE 32

struct readwrite {
    pthread_mutex_t lock;
    pthread_cond_t readers;
    pthread_cond_t writers;
    int rcount, wcount, rwait, wwait, rin;
    int change;
};

struct readwrite library;
char book[BSIZE];
int working = 1;

/* Initialize a buffer */

void init(struct readwrite *rw) {
    int err;

    if ((err = pthread_mutex_init(&rw->lock, 0)) != 0)
        syserr(err, "mutex init failed");
    if ((err = pthread_cond_init(&rw->readers, 0)) != 0)
        syserr(err, "cond init 1 failed");
    if ((err = pthread_cond_init(&rw->writers, 0)) != 0)
        syserr(err, "cond init 2 failed");

    rw->rcount = rw->wcount = rw->rwait = rw->wwait = rw->change = rw->rin = 0;
}

/* Destroy the buffer */

void destroy(struct readwrite *b) {
    int err;

    if ((err = pthread_cond_destroy(&b->readers)) != 0)
        syserr(err, "cond destroy 1 failed");
    if ((err = pthread_cond_destroy(&b->writers)) != 0)
        syserr(err, "cond destroy 2 failed");
    if ((err = pthread_mutex_destroy(&b->lock)) != 0)
        syserr(err, "mutex destroy failed");
}

void reader_enter(struct readwrite *lib) {
    int err;

    if ((err = pthread_mutex_lock(&lib->lock)) != 0)
        syserr(err, "lock failed");


    if (lib->rwait) {
        lib->rcount++;
        pthread_cond_wait(&lib->readers, &lib->lock);
        lib->rwait = 0;
        lib->rcount--;
    }
    lib->rin++;

    lib->wwait = 1;
    if (lib->rcount > 0) {
        if ((err = pthread_cond_signal(&lib->readers)) != 0)
            syserr(err, "cond signal failed");
    } else {
        if ((err = pthread_mutex_unlock(&lib->lock)) != 0)
            syserr(err, "unlock failed");
    }
}

void writer_enter(struct readwrite *lib) {
    int err;

    if ((err = pthread_mutex_lock(&lib->lock)) != 0)
        syserr(err, "lock failed");


    if (lib->wwait) {
        lib->wcount++;
        pthread_cond_wait(&lib->writers, &lib->lock);
        lib->wcount--;
    }
    lib->wwait = 1;
    lib->rwait = 1;

    if ((err = pthread_mutex_unlock(&lib->lock)) != 0)
        syserr(err, "unlock failed");

}

void reader_exit(struct readwrite *lib) {
    int err;

    if ((err = pthread_mutex_lock(&lib->lock)) != 0)
        syserr(err, "lock failed");

    lib->rin--;
    if (lib->wcount != 0 && lib->rin == 0) {
        lib->rwait = 1;
        if ((err = pthread_cond_signal(&lib->writers)) != 0)
            syserr(err, "cond signal failed");
    }

    if ((err = pthread_mutex_unlock(&lib->lock)) != 0)
        syserr(err, "unlock failed");
}

void writer_exit(struct readwrite *lib) {
    int err;

    if ((err = pthread_mutex_lock(&lib->lock)) != 0)
        syserr(err, "lock failed");

    if (lib->rcount > 0) {
        if ((err = pthread_cond_signal(&lib->readers)) != 0)
            syserr(err, "cond signal failed");
    } else {
        if ((err = pthread_cond_signal(&lib->writers)) != 0)
            syserr(err, "cond signal failed");
    }

    if ((err = pthread_mutex_unlock(&lib->lock)) != 0)
        syserr(err, "unlock failed");
}

void *reader(void *data) {
    while (working) {
        reader_enter(&library);
        printf("reader read: %s\n", book); /* reading */
        reader_exit(&library);
    }
    return 0;
}

void *writer(void *data) {
    int l;
    while (working) {
        writer_enter(&library);
        l = rand() % 10;
        snprintf(book, BSIZE, "6 times a number %d %d %d %d %d %d", l, l, l, l, l, l);
        writer_exit(&library);
    }
    return 0;
}


int main() {
    pthread_t th[READERS + WRITERS];
    pthread_attr_t attr;
    int i, err;
    void *retval;

    srand((unsigned) time(0));

    init(&library);
    if ((err = pthread_attr_init(&attr)) != 0)
        syserr(err, "attr_init failed");
    if ((err = pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_JOINABLE)) != 0)
        syserr(err, "attr_setdetachstate failed");

    for (i = 0; i < READERS + WRITERS; i++) {
        if (i < READERS) {
            if ((err = pthread_create(&th[i], &attr, reader, 0)) != 0)
                syserr(err, "create failed");
        } else if ((err = pthread_create(&th[i], &attr, writer, 0)) != 0)
            syserr(err, "create failed");
    }

    sleep(NAP);
    working = 0;

    for (i = 0; i < READERS + WRITERS; i++) {
        if ((err = pthread_join(th[i], &retval)) != 0)
            syserr(err, "join failed");
    }

    if ((err = pthread_attr_destroy(&attr)) != 0)
        syserr(err, "cond destroy failed");
    destroy(&library);
    return 0;
}
