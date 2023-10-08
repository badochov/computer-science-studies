/*
 * A template for the 2019 MPI lab at the University of Warsaw.
 * Copyright (C) 2016, Konrad Iwanicki.
 * Refactoring 2019, Łukasz Rączkowski
 */

#include <iostream>
#include <string>
#include <cassert>
#include <mpi.h>
#include "graph-base.h"
#include "graph-utils.h"

void handle_error(int err);

static void runFloydWarshallParallel(Graph *graph, int numProcesses, int myRank)
{
    assert(numProcesses <= graph->numVertices);
    int m = graph->numVertices;

    int *row = (int *)malloc(graph->numVertices * sizeof(int));
    assert(row != nullptr);

    int curr_rank = 0;
    int row_max = getFirstGraphRowOfProcess(graph->numVertices, numProcesses, curr_rank + 1);
    // std::cout << myRank << std::endl;

    for (int k = 0; k < m; ++k)
    {
        if (k == row_max)
        {
            curr_rank++;
            row_max = getFirstGraphRowOfProcess(graph->numVertices, numProcesses, curr_rank + 1);
        }
        // std::cout << myRank << std::endl;
        if (curr_rank == myRank)
        {
            // std::cout << "RUF:" << curr_rank << std::endl;
            memcpy(row, graph->data[k - graph->firstRowIdxIncl], graph->numVertices * sizeof(int));
            // std::cout << "FOO" << std::endl;
            handle_error(MPI_Bcast(row, /* the message will be written here */
                                   /* if my_rank==root, the message will be read from here */
                                   graph->numVertices, /* number of items in the message */
                                   MPI_INT,            /* type of data in the message */
                                   myRank,             /* if my_rank==root, I'm sending, otherwise I'm receiving */
                                   MPI_COMM_WORLD      /* communicator to use */
                                   ));
            // std::cout << "BAR" << std::endl;
        }
        else
        {
            // std::cout << "R:" << curr_rank << std::endl;
            handle_error(MPI_Bcast(row, /* the message will be written here */
                                   /* if my_rank==root, the message will be read from here */
                                   graph->numVertices, /* number of items in the message */
                                   MPI_INT,            /* type of data in the message */
                                   curr_rank,          /* if my_rank==root, I'm sending, otherwise I'm receiving */
                                   MPI_COMM_WORLD      /* communicator to use */
                                   ));
            // std::cout << "BAZ" << std::endl;
        }
        // std::cout << myRank << std::endl;
        for (int i = 0; i < graph->lastRowIdxExcl - graph->firstRowIdxIncl; ++i)
        {
            for (int j = 0; j < m; ++j)
            {
                int pathSum = graph->data[i][k] + row[j];
                if (graph->data[i][j] > pathSum)
                {
                    graph->data[i][j] = pathSum;
                }
            }
        }
    }

    free(row);
}

int main(int argc, char *argv[])
{
    int numVertices = 0;
    int numProcesses = 0;
    int myRank = 0;
    int showResults = 0;

    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &numProcesses);
    MPI_Comm_rank(MPI_COMM_WORLD, &myRank);

#ifdef USE_RANDOM_GRAPH
#ifdef USE_RANDOM_SEED
    srand(USE_RANDOM_SEED);
#endif
#endif

    for (int i = 1; i < argc; ++i)
    {
        if (std::string(argv[i]).compare("--show-results") == 0)
        {
            showResults = 1;
        }
        else
        {
            numVertices = std::stoi(argv[i]);
        }
    }

    if (numVertices <= 0)
    {
        std::cerr << "Usage: " << argv[0] << "  [--show-results] <num_vertices>" << std::endl;
        MPI_Finalize();
        return 1;
    }

    if (numProcesses > numVertices)
    {
        numProcesses = numVertices;

        if (myRank >= numProcesses)
        {
            MPI_Finalize();
            return 0;
        }
    }

    std::cerr << "Running the Floyd-Warshall algorithm for a graph with " << numVertices << " vertices." << std::endl;

    auto graph = createAndDistributeGraph(numVertices, numProcesses, myRank);
    if (graph == nullptr)
    {
        std::cerr << "Error distributing the graph for the algorithm." << std::endl;
        MPI_Finalize();
        return 2;
    }

    if (showResults)
    {
        collectAndPrintGraph(graph, numProcesses, myRank);
    }

    double startTime = MPI_Wtime();

    runFloydWarshallParallel(graph, numProcesses, myRank);

    double endTime = MPI_Wtime();

    std::cerr
        << "The time required for the Floyd-Warshall algorithm on a "
        << numVertices
        << "-node graph with "
        << numProcesses
        << " process(es): "
        << endTime - startTime
        << std::endl;

    if (showResults)
    {
        collectAndPrintGraph(graph, numProcesses, myRank);
    }

    destroyGraph(graph, numProcesses, myRank);

    MPI_Finalize();

    return 0;
}
