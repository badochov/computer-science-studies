#ifdef QUEUE_EL_T

#include "queue.h"
#include "util.h"
#include "err.h"
#include <stdlib.h>


struct QUEUE_STRUCT {
    QUEUE_EL_T *arr;
    size_t max_size;
    size_t size;
    size_t read;
    size_t write;
    size_t curr_arr_size;
};


QUEUE_T QUEUE_INIT(size_t max_size) {
    QUEUE_T queue = calloc(1, sizeof(struct QUEUE_STRUCT));
    if (queue == NULL) {
        return NULL;
    }
//    queue->curr_arr_size = 1069;
//    queue->arr=calloc(1069, sizeof(QUEUE_EL_T));
    queue->max_size = max_size;
    return queue;
}

static bool resize_arr(QUEUE_T queue, size_t new_size) {
    queue->curr_arr_size = new_size;
    QUEUE_EL_T *new_arr = realloc(queue->arr, queue->curr_arr_size * sizeof(QUEUE_EL_T));
    if (new_arr == NULL && queue->curr_arr_size > 0) {
        return false;
    }
    queue->arr = new_arr;

    return true;
}

static bool make_queue_bigger(QUEUE_T queue) {
    if (!resize_arr(queue, max(1, queue->curr_arr_size * 2))) {
        return false;
    }
    for (size_t i = 0; i < queue->read; ++i) {
        queue->arr[queue->size - 1 + i] = queue->arr[i];
    }
    if(queue->size != 1){
        queue->write = (queue->read + queue->size - 1) % queue->curr_arr_size;

    }

    return true;
}

int QUEUE_ADD(QUEUE_T queue, QUEUE_EL_T el) {
    if (queue->size == queue->max_size) {
        return 1;
    }

    ++queue->size;
    if (queue->size > queue->curr_arr_size) {
        if (!make_queue_bigger(queue)) {
            return -1;
        }
    }

    queue->arr[queue->write] = el;

    ++queue->write;
    if (queue->write == queue->curr_arr_size) {
        queue->write = 0;
    }

    return 0;
}

QUEUE_EL_T QUEUE_POP(QUEUE_T queue) {
    QUEUE_EL_T el = queue->arr[queue->read];

    ++queue->read;
    --queue->size;
    if (queue->read == queue->curr_arr_size) {
        queue->read = 0;
    }

//    if (4 * queue->size < queue->curr_arr_size) {
//        if (!resize_arr(queue, queue->curr_arr_size / 2)) {
//            syserr("Realloc failed!");
//        }
//    }

    return el;
}

size_t QUEUE_SIZE(QUEUE_T queue) {
    return queue->size;
}


void QUEUE_DESTROY(QUEUE_T queue) {
    if (queue == NULL) {
        return;
    }
    free(queue->arr);
    free(queue);
}

#endif
