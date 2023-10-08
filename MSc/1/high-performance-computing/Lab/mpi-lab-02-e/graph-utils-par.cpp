/*
 * A template for the 2019 MPI lab at the University of Warsaw.
 * Copyright (C) 2016, Konrad Iwanicki.
 * Refactoring 2019, Łukasz Rączkowski
 */

#include <cassert>
#include <algorithm>
#include <mpi.h>
#include "graph-base.h"
#include "graph-utils.h"

#define TAG 42

void handle_error(int err)
{
    assert(err == MPI_SUCCESS);
}

int getFirstGraphRowOfProcess(int numVertices, int numProcesses, int myRank)
{
    int perProc = numVertices / numProcesses;
    int ret = perProc * myRank;
    int leftover = numVertices - perProc * numProcesses;

    ret += std::min(leftover, myRank);

    return ret;
}

Graph *createAndDistributeGraph(int numVertices, int numProcesses, int myRank)
{
    assert(numProcesses >= 1 && myRank >= 0 && myRank < numProcesses);

    auto graph = allocateGraphPart(
        numVertices,
        getFirstGraphRowOfProcess(numVertices, numProcesses, myRank),
        getFirstGraphRowOfProcess(numVertices, numProcesses, myRank + 1));

    if (graph == nullptr)
    {
        return nullptr;
    }

    assert(graph->numVertices > 0 && graph->numVertices == numVertices);
    assert(graph->firstRowIdxIncl >= 0 && graph->lastRowIdxExcl <= graph->numVertices);

    if (myRank == 0)
    {
        int *row = (int *)malloc(numVertices * sizeof(int));
        assert(row != nullptr);

        for (int i = 0; i < graph->lastRowIdxExcl; ++i)
        {
            initializeGraphRow(graph->data[i], graph->firstRowIdxIncl + i, graph->numVertices);
        }

        for (int proc = 1; proc < numProcesses; proc++)
        {
            for (int i = getFirstGraphRowOfProcess(numVertices, numProcesses, proc); i < getFirstGraphRowOfProcess(numVertices, numProcesses, proc + 1); ++i)
            {
                initializeGraphRow(row, i, graph->numVertices);
                handle_error(MPI_Send(row,           /* where the message will be saved */
                                      numVertices,   /* max number of elements we expect */
                                      MPI_INT,       /* type of data in the message */
                                      proc,          /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                                      TAG,           /* if not MPI_ANY_TAG, receive only with a certain tag */
                                      MPI_COMM_WORLD /* communicator to use */
                                      ));
            }
        }

        free(row);
    }
    else
    {
        for (int i = 0; i < graph->lastRowIdxExcl - graph->firstRowIdxIncl; i++)
        {
            handle_error(MPI_Recv(graph->data[i],   /* where the message will be saved */
                                  numVertices,      /* max number of elements we expect */
                                  MPI_INT,          /* type of data in the message */
                                  0,                /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                                  TAG,              /* if not MPI_ANY_TAG, receive only with a certain tag */
                                  MPI_COMM_WORLD,   /* communicator to use */
                                  MPI_STATUS_IGNORE /* if not MPI_STATUS_IGNORE, write comm info here */
                                  ));
        }
    }

    return graph;
}

void collectAndPrintGraph(Graph *graph, int numProcesses, int myRank)
{
    assert(numProcesses >= 1 && myRank >= 0 && myRank < numProcesses);
    assert(graph->numVertices > 0);
    assert(graph->firstRowIdxIncl >= 0 && graph->lastRowIdxExcl <= graph->numVertices);

    if (myRank == 0)
    {
        int *row = (int *)malloc(graph->numVertices * sizeof(int));
        assert(row != nullptr);

        for (int i = 0; i < graph->lastRowIdxExcl; ++i)
        {
            printGraphRow(graph->data[i], i, graph->numVertices);
        }

        for (int proc = 1; proc < numProcesses; proc++)
        {
            for (int i = getFirstGraphRowOfProcess(graph->numVertices, numProcesses, proc); i < getFirstGraphRowOfProcess(graph->numVertices, numProcesses, proc + 1); ++i)
            {
                handle_error(MPI_Recv(row,                /* where the message will be saved */
                                      graph->numVertices, /* max number of elements we expect */
                                      MPI_INT,            /* type of data in the message */
                                      proc,               /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                                      TAG,                /* if not MPI_ANY_TAG, receive only with a certain tag */
                                      MPI_COMM_WORLD,     /* communicator to use */
                                      MPI_STATUS_IGNORE   /* if not MPI_STATUS_IGNORE, write comm info here */
                                      ));

                printGraphRow(row, i, graph->numVertices);
            }
        }

        free(row);
    }
    else
    {
        for (int i = 0; i < graph->lastRowIdxExcl - graph->firstRowIdxIncl; i++)
        {
            handle_error(MPI_Send(graph->data[i],     /* where the message will be saved */
                                  graph->numVertices, /* max number of elements we expect */
                                  MPI_INT,            /* type of data in the message */
                                  0,                  /* if not MPI_ANY_SOURCE, receive only from source with the given rank  */
                                  TAG,                /* if not MPI_ANY_TAG, receive only with a certain tag */
                                  MPI_COMM_WORLD      /* communicator to use */
                                  ));
        }
    }
}

void destroyGraph(Graph *graph, int numProcesses, int myRank)
{
    freeGraphPart(graph);
}
