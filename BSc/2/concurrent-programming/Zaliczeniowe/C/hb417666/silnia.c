#include "cacti.h"
#include <stdio.h>
#include <stdlib.h>

#define MSG_CALC 0x01
#define MSG_SPAWNED 0x02

#define SIZEOF_ARR(arr) sizeof((arr)) / sizeof((arr)[0])

#define UNUSED(arg) (void)(arg)

typedef struct {
    unsigned long long prev;
    unsigned long long k;
    unsigned long long n;
} calc_t;

typedef struct {
    calc_t *calc;
    actor_id_t parent_id;
} state_t;

void first_hello_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    UNUSED(data);
    state_t *state = calloc(1, sizeof(state_t));
    *stateptr = state;
}

void hello_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    state_t *state = malloc(sizeof(state_t));
    state->calc = NULL;
    state->parent_id = (actor_id_t) data;
    *stateptr = state;
    message_t message = {.message_type= MSG_SPAWNED, .nbytes= sizeof(actor_id_t), .data=(void *) actor_id_self()};
    send_message(state->parent_id, message);
}

message_t godie_msg = {.message_type= MSG_GODIE, .nbytes= 0, .data=NULL};

void calc_handler(void **stateptr, size_t nbytes, void *data);

void spawned_handler(void **stateptr, size_t nbytes, void *data);

act_t act[] = {hello_handler, calc_handler, spawned_handler,};
role_t role = {.nprompts= SIZEOF_ARR(act), .prompts=act};

void calc_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    state_t *state = *stateptr;
    state->calc = data;
    calc_t *calc = state->calc;
    ++calc->k;
    calc->prev *= calc->k;
    if (calc->k < calc->n) {
        message_t message = {.message_type= MSG_SPAWN, .nbytes= sizeof(role_t *), .data=&role};
        send_message(actor_id_self(), message);
    } else {
        free(state);
        send_message(actor_id_self(), godie_msg);
    }
}

void spawned_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    state_t *state = *stateptr;
    actor_id_t actor_id = (actor_id_t) data;
    message_t message = {.message_type= MSG_CALC, .nbytes= sizeof(calc_t *), .data=state->calc};
    send_message(actor_id, message);

    free(state);
    send_message(actor_id_self(), godie_msg);
}


act_t act_first[] = {first_hello_handler, calc_handler, spawned_handler};
role_t role_first = {.nprompts= SIZEOF_ARR(act_first), .prompts=act_first};

int main() {

    calc_t calc = {1, 0, 0};
    scanf("%llu", &calc.n);
    actor_id_t actor_id;
    if (actor_system_create(&actor_id, &role_first) < 0) {
        printf("Error: Couldn't create actor system!");
        return -1;

    }
    message_t message = {.message_type= MSG_CALC, .nbytes= sizeof(calc_t *), .data=&calc};
    send_message(actor_id, message);

    actor_system_join(actor_id);
    printf("%llu", calc.prev);

    return 0;
}
