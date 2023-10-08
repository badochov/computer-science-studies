#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include "err.h"

void change_in_and_out_files_in_parent(int pipe_dsc[2]) {
    if (close(STDIN_FILENO) == -1) syserr("Error in child, close (0)\n");
    if (dup2(pipe_dsc[0], STDIN_FILENO) != STDIN_FILENO) syserr("Error in child, dup (pipe_dsc [0])\n");
    if (close(pipe_dsc[0]) == -1) syserr("Error in child, close (pipe_dsc [0])\n");
    if (close(pipe_dsc[1]) == -1) syserr("Error in child, close (pipe_dsc [1])\n");
}

char *new_name(char *argv[]) {
    char *res = malloc(strlen(argv[0]) + 3);
    sprintf(res, "./%s", argv[0]);
    return res;
}

char **get_new_args(int argc, char *argv[]) {
    char **res = malloc(argc);
    res[0] = argv[0];
    for (int i = 0; i < argc - 2; i++) {
        res[i + 1] = argv[i + 2];
    }
    res[argc - 1] = NULL;

    return res;
}


void change_in_and_out_files_in_child(int pipe_dsc[2]) {
    if (close(STDOUT_FILENO) == -1) syserr("Error in child, close (0)\n");
    if (dup2(pipe_dsc[1], STDOUT_FILENO) != STDOUT_FILENO) syserr("Error in child, dup (pipe_dsc [0])\n");
    if (close(pipe_dsc[0]) == -1) syserr("Error in child, close (pipe_dsc [0])\n");
    if (close(pipe_dsc[1]) == -1) syserr("Error in child, close (pipe_dsc [1])\n");
}

void handle_child(int pipe_dsc[2], int argc, char *argv[]) {

    change_in_and_out_files_in_child(pipe_dsc);

    char *program_name = argv[1];
    char *new_args[] = {program_name, NULL};
    execvp(program_name, new_args);
    syserr(program_name);

}

void handle_parent(int pipe_dsc[2], int argc, char *argv[]) {
    change_in_and_out_files_in_parent(pipe_dsc);

    char *command = new_name(argv);
    char **new_args = get_new_args(argc, argv);
    execvp(command, new_args);
    syserr("Error in execvp\n");

}

int main(int argc, char *argv[]) {

    if (argc == 1) {
        exit(0);
    }
    if (argc == 2) {
        execvp(argv[1], argv + 1);

        syserr("Error in execvp\n");
    }

    int pipe_dsc[2];

    if (pipe(pipe_dsc) == -1)
        syserr("Error in pipe\n");
    switch (fork()) {
        case -1:
            syserr("Error in fork\n");
        case 0:
            handle_child(pipe_dsc, argc, argv);
        default:
            handle_parent(pipe_dsc, argc, argv);
    }
}
