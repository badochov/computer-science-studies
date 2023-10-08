#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <time.h>
#include <stdint.h>
#include <mpi.h>

#define TAG 42
#define SIZE 100

void handle_err(int err)
{
}

double experiment(void *buff, int size, int rank)
{
    double startTime;
    double endTime;

    startTime = MPI_Wtime();

    if (rank == 0)
    {
        handle_err(MPI_Recv(buff,             /* where the message will be saved */
                            size,             /* max number of elements we expect */
                            MPI_BYTE,         /* type of data in the message */
                            1,                /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                            TAG,              /* if not MPI_ANY_TAG, receive only with a certain tag */
                            MPI_COMM_WORLD,   /* communicator to use */
                            MPI_STATUS_IGNORE /* if not MPI_STATUS_IGNORE, write comm info here */
                            ));
    }
    else
    {
        handle_err(MPI_Send(buff,          /* pointer to the message */
                            size,          /* number of items in the message */
                            MPI_BYTE,      /* type of data in the message */
                            0,             /* rank of the destination process */
                            TAG,           /* app-defined message type */
                            MPI_COMM_WORLD /* communicator to use */
                            ));
    }

    if (rank == 1)
    {
        handle_err(MPI_Recv(buff,             /* where the message will be saved */
                            size,             /* max number of elements we expect */
                            MPI_BYTE,         /* type of data in the message */
                            0,                /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                            TAG,              /* if not MPI_ANY_TAG, receive only with a certain tag */
                            MPI_COMM_WORLD,   /* communicator to use */
                            MPI_STATUS_IGNORE /* if not MPI_STATUS_IGNORE, write comm info here */
                            ));
    }
    else
    {
        handle_err(MPI_Send(buff,          /* pointer to the message */
                            size,          /* number of items in the message */
                            MPI_BYTE,      /* type of data in the message */
                            1,             /* rank of the destination process */
                            TAG,           /* app-defined message type */
                            MPI_COMM_WORLD /* communicator to use */
                            ));
    }

    endTime = MPI_Wtime();

    return endTime - startTime;
}

int main(int argc, char *argv[])
{
    int rank;
    char name[32] = {0};
    int sizes[] = {1, 10, 100, 5e4};

    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    sprintf(name, "results-%d.ctx", rank);
    FILE *f = fopen(name, "w");

    for (int i = 0; i < sizeof(sizes) / sizeof(sizes[0]); i++)
    {
        int size = sizes[i];
        void *buff = malloc(size);
        printf("%d\n", size);

        for (int i = 0; i < SIZE; i++)
        {
            double time = experiment(buff, size, rank);
            fprintf(f, "%d %d %f\n", i, size, time);
        }

        free(buff);
    }

    fclose(f);

    MPI_Finalize();

    return 0;
}
