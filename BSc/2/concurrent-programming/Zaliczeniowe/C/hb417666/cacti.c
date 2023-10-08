#include "cacti.h"
#include <stdbool.h>
#include <stdlib.h>
#include <pthread.h>
#include "util.h"
#include "err.h"

#define INIT_ACTOR_ARR_SIZE 1024

#include "message_queue.h"

#include "thread_pool.h"

#define UNUSED(arg) (void)(arg)

#define GET_ACTOR(i) state.actors[(i)]


#define CHECK(sys) if((sys)){syserr(#sys" returned non zero value!");}
#define LOCK_MUTEX() CHECK(pthread_mutex_lock(&state.mutex))
#define UNLOCK_MUTEX() CHECK(pthread_mutex_unlock(&state.mutex))

_Thread_local actor_id_t threads_actor;

void add_actor_to_jobs(actor_id_t actor_id);

typedef struct actor_data {
    void *state;
    bool active;
    bool is_processed;
    message_t__queue_t messages;
    const role_t *role;
} *actor_data_t;


typedef struct {
    pthread_t sig_handler;
    pthread_mutex_t mutex;
    actor_data_t *actors;
    size_t actor_count;
    size_t dead_actors_count;
    bool can_actors_spawn;
    size_t actors_arr_size;
    pool_t pool;

} actor_system_state_t;

actor_system_state_t state;


bool known_actor(actor_id_t id) {
    return (size_t) id < state.actor_count;
}

int send_message_helper(actor_id_t actor, message_t message) {
    int res = 0;
    if (known_actor(actor)) {
        if (GET_ACTOR(actor)->active) {
            message_t__queue_add(GET_ACTOR(actor)->messages, message);
            if (!GET_ACTOR(actor)->is_processed) {
                add_actor_to_jobs(actor);
            }
        } else {
            res = -1;
        }
    } else {
        res = -2;
    }
    return res;
}


void destroy_actor(actor_id_t actor_id) {
    message_t__queue_destroy(GET_ACTOR(actor_id)->messages);
    GET_ACTOR(actor_id)->messages = NULL;
}

void state_destroy() {
    CHECK(pthread_mutex_destroy(&state.mutex))

    if (state.actors != NULL) {
        for (size_t i = 0; i < state.actor_count; ++i) {
            destroy_actor(i);
            free(GET_ACTOR(i));
            GET_ACTOR(i) = NULL;
        }
        free(state.actors);
        state.actors = NULL;
    }

    pool_destroy(state.pool);
    state.pool = NULL;

}


size_t create_new_actor_helper(role_t *const role) {
    size_t id = state.actor_count;


    GET_ACTOR(id) = calloc(1, sizeof(struct actor_data));

    GET_ACTOR(id)->role = role;
    GET_ACTOR(id)->messages = message_t__queue_init(ACTOR_QUEUE_LIMIT);
    GET_ACTOR(id)->active = true;

    ++state.actor_count;
    return id;
}

bool is_enough_space_for_new_actor() {
    return state.actor_count < state.actors_arr_size;
}


actor_id_t create_new_actor(role_t *role) {
    if (!is_enough_space_for_new_actor()) {
        if (state.actors_arr_size == CAST_LIMIT) {
            fatal("The limit of actors in system has been reached!");
        }
        state.actors_arr_size = min(CAST_LIMIT, 2 * state.actors_arr_size);
        actor_data_t *new_arr = realloc(state.actors, sizeof(actor_data_t) * state.actors_arr_size);
        if (new_arr != NULL) {
            state.actors = new_arr;
        }
    }

    return create_new_actor_helper(role);
}


void process_actor(void *arg);

void spawn(actor_id_t actor_id, role_t *role) {
    if (state.can_actors_spawn) {
        actor_id_t new_actor_id = create_new_actor(role);
        message_t message = {.message_type=MSG_HELLO, .nbytes = sizeof(void *), .data=(void *) actor_id};
        send_message_helper(new_actor_id, message);
    }
}

void die(actor_id_t actor_id) {
    GET_ACTOR(actor_id)->active = false;
}

void handle_message(actor_id_t actor_id, message_t message) {
    act_t fn = GET_ACTOR(actor_id)->role->prompts[message.message_type];

    UNLOCK_MUTEX()

    fn(&GET_ACTOR(actor_id)->state, message.nbytes, message.data);

    LOCK_MUTEX()
}

void change_is_processed(actor_id_t actor_id, bool new_state) {
    GET_ACTOR(actor_id)->is_processed = new_state;
}

bool add_new_actor_message_to_execute(actor_id_t actor_id) {
    if (message_t__queue_size(GET_ACTOR(actor_id)->messages) > 0) {
        pool_add_job(state.pool, process_actor, (void *) actor_id);
        return true;
    }

    change_is_processed(actor_id, false);

    return false;
}

void end_sig_handler() {
    CHECK(pthread_cancel(state.sig_handler))
    CHECK(pthread_join(state.sig_handler, NULL))
}

void kill_actor_system() {
    pool_end(state.pool);
    end_sig_handler();
}

void kill_actor(actor_id_t actor_id) {
    state.dead_actors_count++;

    destroy_actor(actor_id);
    if (state.dead_actors_count == state.actor_count) {
        kill_actor_system();
    }
}

void process_actor(void *arg) {
    LOCK_MUTEX()
    actor_id_t actor_id = (actor_id_t) arg;
    threads_actor = actor_id;

    message_t message = message_t__queue_pop(GET_ACTOR(actor_id)->messages);

    switch (message.message_type) {
        case MSG_SPAWN:
            spawn(actor_id, message.data);
            break;
        case MSG_GODIE:
            die(actor_id);
            break;
        default:
            handle_message(actor_id, message);
    }

    if (!add_new_actor_message_to_execute(actor_id) &&
        !GET_ACTOR(actor_id)->active &&
        message_t__queue_size(GET_ACTOR(actor_id)->messages) == 0) {
        kill_actor(actor_id);
    }

    UNLOCK_MUTEX()
}


int init_state_actor_data(size_t size) {
    state.actors_arr_size = size;
    state.actors = calloc(size, sizeof(actor_data_t));
    if (state.actors == NULL) {
        return -1;
    }
    return 0;
}


void *sigint_handler(int sig) {
    UNUSED(sig);

    for (size_t i = 0; i < state.actor_count; i++) {
        die(i);
    }
    state.can_actors_spawn = false;

    return NULL;
}

void *thread(void *arg) {
    UNUSED(arg);

    sigset_t set;
    sigemptyset(&set);
    sigaddset(&set, SIGINT);
    int sig;
    sigwait(&set, &sig);
    return sigint_handler(sig);
}

int handle_sigint() {
    disable_signal(SIGINT, NULL);
    return pthread_create(&state.sig_handler, NULL, thread, NULL);
}

bool init_state() {
    sigset_t sig;
    disable_all_signals(&sig);
    state.pool = pool_create(POOL_SIZE, CAST_LIMIT);
    if (init_state_actor_data(INIT_ACTOR_ARR_SIZE) ||
        pthread_mutex_init(&state.mutex, NULL) ||
        state.pool == NULL) {
        state_destroy();
        return false;
    }
    state.can_actors_spawn = true;
    state.actor_count = state.dead_actors_count = 0;

    set_mask(&sig, NULL);

    return true;
}


actor_id_t actor_id_self() {
    return threads_actor;
}


int actor_system_create(actor_id_t *actor, role_t *const role) {
    if (CAST_LIMIT <= 0 || ACTOR_QUEUE_LIMIT <= 0 || POOL_SIZE <= 0) {
        return -1;
    }
    if (!init_state()) {
        return -1;
    }

    *actor = create_new_actor_helper(role);
    message_t message = {.message_type=MSG_HELLO, .nbytes = sizeof(void *), .data=NULL};
    send_message_helper(*actor, message);
    return handle_sigint();
}

void actor_system_join(actor_id_t actor) {
    bool correct_actor = known_actor(actor);

    if (correct_actor) {
        pool_join(state.pool);
        state_destroy();

        enable_signal(SIGINT, NULL);
    } else {
        fatal("Wrong id of actor called in join!");
    }
}

int send_message(actor_id_t actor, message_t message) {
    LOCK_MUTEX()

    int res = send_message_helper(actor, message);

    UNLOCK_MUTEX()

    return res;
}

void add_actor_to_jobs(actor_id_t actor_id) {
    pool_add_job(state.pool, process_actor, (void *) actor_id);
    change_is_processed(actor_id, true);
}