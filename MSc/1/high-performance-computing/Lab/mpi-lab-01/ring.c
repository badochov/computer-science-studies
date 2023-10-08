#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <time.h>
#include <stdint.h>
#include <mpi.h>

#define TAG 42

void handle_err(int err)
{
}

int main(int argc, char *argv[])
{
    MPI_Init(&argc, &argv);

    int64_t res = 1;
    int numProcesses, rank;

    MPI_Comm_size(MPI_COMM_WORLD, &numProcesses);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    if (rank != 0)
    {
        handle_err(MPI_Recv(&res, /* where the message will be saved */
                            1, /* max number of elements we expect */
                            MPI_INT64_T, /* type of data in the message */
                            rank - 1, /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                            TAG, /* if not MPI_ANY_TAG, receive only with a certain tag */
                            MPI_COMM_WORLD, /* communicator to use */
                            MPI_STATUS_IGNORE /* if not MPI_STATUS_IGNORE, write comm info here */
                            ));
        res *= rank;
    }

    handle_err(MPI_Send(&res, /* pointer to the message */
                        1, /* number of items in the message */
                        MPI_INT64_T, /* type of data in the message */
                        (rank + 1) % numProcesses, /* rank of the destination process */
                        TAG, /* app-defined message type */
                        MPI_COMM_WORLD /* communicator to use */
                        ));

    if (rank == 0)
    {
        handle_err(MPI_Recv(&res, /* where the message will be saved */
                            1, /* max number of elements we expect */
                            MPI_INT64_T, /* type of data in the message */
                            numProcesses - 1, /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                            TAG, /* if not MPI_ANY_TAG, receive only with a certain tag */
                            MPI_COMM_WORLD, /* communicator to use */
                            MPI_STATUS_IGNORE /* if not MPI_STATUS_IGNORE, write comm info here */
                            ));

        printf("Num: %lld\n", res);
    }

    MPI_Finalize();

    return 0;
}
