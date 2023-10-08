#include "cacti.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include "unistd.h"

#define MSG_COL 0x01
#define MSG_SPAWNED 0x02
#define MSG_CALC 0x03
#define MSG_CREATED 0x04
#define MSG_INIT 0x05

#define SIZEOF_ARR(arr) sizeof((arr)) / sizeof((arr)[0])
#define UNUSED(arg) (void)(arg)

typedef struct {
    size_t columns;
    long current;
    size_t rows;
} matrix_data_t;

typedef struct {
    long value;
    size_t time;
} field_t;

typedef struct {
    long prev;
    size_t row;
    field_t **matrix;
} calc_t;

typedef struct {
    matrix_data_t matrix_data;
    actor_id_t father;
    actor_id_t next;
    size_t rows_calculated;
    bool init;
    calc_t *rows;
} state_t;

message_t godie_msg = {.message_type= MSG_GODIE, .nbytes= 0, .data=NULL};


void hello_handler(void **stateptr, size_t nbytes, void *data);

void spawned_handler(void **stateptr, size_t nbytes, void *data);

void column_handler(void **stateptr, size_t nbytes, void *data);

void calc_handler(void **stateptr, size_t nbytes, void *data);

void created_handler(void **stateptr, size_t nbytes, void *data);

act_t act[] = {&hello_handler, &column_handler, &spawned_handler, &calc_handler, &created_handler};
role_t role = {.nprompts= SIZEOF_ARR(act), .prompts=act};

void hello_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    state_t *state = calloc(1, sizeof(state_t));
    state->father = (actor_id_t) data;
    *stateptr = state;
    actor_id_t *id = malloc(sizeof(actor_id_self()));
    *id = actor_id_self();
    message_t message = {.message_type= MSG_SPAWNED, .nbytes= sizeof(actor_id_t *), .data=id};
    send_message(state->father, message);
}

void spawned_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    state_t *state = *stateptr;
    actor_id_t actor_id = *(actor_id_t *) data;
    state->next = actor_id;
    free(data);
    message_t message = {.message_type= MSG_COL, .nbytes= sizeof(matrix_data_t *), .data=&state->matrix_data};
    send_message(actor_id, message);
}

void created_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    UNUSED(data);
    state_t *state = *stateptr;
    message_t message = {.message_type= MSG_CREATED, .nbytes= sizeof(void *), .data=NULL};
    send_message(state->father, message);
}

void column_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    matrix_data_t *columns = data;
    state_t *state = *stateptr;
    state->matrix_data = *columns;
    ++state->matrix_data.current;
    if ((size_t)state->matrix_data.current != columns->columns - 1) {
        message_t message = {.message_type= MSG_SPAWN, .nbytes= sizeof(role_t *), .data=&role};
        send_message(actor_id_self(), message);
    } else {
        message_t message = {.message_type= MSG_CREATED, .nbytes= sizeof(void *), .data=NULL};
        send_message(state->father, message);
    }
}

void calc_handler(void **stateptr, size_t nbytes, void *data) {
    calc_t *calc = data;
    state_t *state = *stateptr;
    field_t field = calc->matrix[calc->row][state->matrix_data.current];
    usleep(field.time * 1000);
    calc->prev += field.value;
    if ((size_t)state->matrix_data.current != state->matrix_data.columns - 1) {
        message_t message = {.message_type= MSG_CALC, .nbytes= nbytes, .data=data};
        send_message(state->next, message);
    }

    ++state->rows_calculated;
    if (state->rows_calculated == state->matrix_data.rows) {
        free(state);
        send_message(actor_id_self(), godie_msg);
    }
}

void init_calc(state_t *state) {
    for (size_t i = 0; i < state->matrix_data.rows; i++) {
        message_t message_calc = {.message_type= MSG_CALC, .nbytes= sizeof(calc_t *), .data=&state->rows[i]};
        send_message(actor_id_self(), message_calc);
    }
}

void init_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);

    calc_t *rows = data;
    state_t *state = *stateptr;
    state->rows = rows;
    if (state->init) {
        init_calc(state);
    }
}


void first_hello_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    UNUSED(data);
    state_t *state = calloc(1, sizeof(state_t));
    *stateptr = state;
}

void first_created_handler(void **stateptr, size_t nbytes, void *data) {
    UNUSED(nbytes);
    UNUSED(data);
    state_t *state = *stateptr;
    state->init = true;
    if (state->rows != NULL) {
        init_calc(state);
    }
}


act_t act_first[] = {&first_hello_handler, &column_handler, &spawned_handler, &calc_handler, &first_created_handler,
                     &init_handler};
role_t role_first = {.nprompts= SIZEOF_ARR(act_first), .prompts=act_first};

int main() {
    size_t k;
    size_t n;

    scanf("%ld", &k);
    scanf("%ld", &n);
    field_t **matrix = calloc(k, sizeof(field_t *));

    calc_t rows[n];
    for (size_t i = 0; i < k; i++) {
        rows[i] = (calc_t) {.matrix=matrix, .prev=0, .row=i};
        matrix[i] = calloc(n, sizeof(field_t));
    }

    for (size_t i = 0; i < k * n; i++) {
        scanf("%ld %ld", &matrix[i / n][i % n].value, &matrix[i / n][i % n].time);
    }
    actor_id_t actor_id;
    if (actor_system_create(&actor_id, &role_first) < 0) {
        printf("Error: Couldn't create actor system!");
        return -1;

    }
    matrix_data_t data = {.current=-1, .rows = k, .columns=n};
    message_t message = {.message_type= MSG_COL, .nbytes= sizeof(matrix_data_t *), .data=&data};
    send_message(actor_id, message);
    message_t message_ini = {.message_type= MSG_INIT, .nbytes= sizeof(calc_t **), .data=&rows};
    send_message(actor_id, message_ini);


    actor_system_join(actor_id);
    for (size_t i = 0; i < k; i++) {

        printf("%ld\n", rows[i].prev);
    }

    for (size_t i = 0; i < k; i++) {
        free(matrix[i]);
    }
    free(matrix);

    return 0;
}
