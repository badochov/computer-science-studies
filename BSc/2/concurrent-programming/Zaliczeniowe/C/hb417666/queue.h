#ifndef CACTI_QUEUE_H
#define CACTI_QUEUE_H
//#define  QUEUE_EL_T int
#ifdef QUEUE_EL_T

#include <stdbool.h>
#include <stddef.h>
#include "template.h"

#define QUEUE_T TEMPLATE(QUEUE_EL_T, queue_t)
#define QUEUE_STRUCT TEMPLATE(QUEUE_EL_T, queue)

#define QUEUE_DESTROY TEMPLATE(QUEUE_EL_T, queue_destroy)
#define QUEUE_INIT TEMPLATE(QUEUE_EL_T, queue_init)
#define QUEUE_ADD TEMPLATE(QUEUE_EL_T, queue_add)
#define QUEUE_POP TEMPLATE(QUEUE_EL_T, queue_pop)
#define QUEUE_SIZE TEMPLATE(QUEUE_EL_T, queue_size)

typedef struct QUEUE_STRUCT *QUEUE_T;


QUEUE_T QUEUE_INIT(size_t max_size);

int QUEUE_ADD(QUEUE_T queue, QUEUE_EL_T message);

QUEUE_EL_T QUEUE_POP(QUEUE_T queue);

size_t QUEUE_SIZE(QUEUE_T queue);

void QUEUE_DESTROY(QUEUE_T queue);


inline void dupa() {}

#endif

#endif //CACTI_QUEUE_H
