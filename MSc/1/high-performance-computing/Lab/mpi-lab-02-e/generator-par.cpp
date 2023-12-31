/*
 * A template for the 2019 MPI lab at the University of Warsaw.
 * Copyright (C) 2016, Konrad Iwanicki.
 * Refactoring 2019, Łukasz Rączkowski
 */

#include <iostream>
#include <string>
#include <mpi.h>
#include "graph-base.h"
#include "graph-utils.h"


int main(int argc, char *argv[]) {
    int numVertices = 0;
    int numProcesses = 0;
    int myRank = 0;

    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &numProcesses);
    MPI_Comm_rank(MPI_COMM_WORLD, &myRank);

#ifdef USE_RANDOM_GRAPH
#ifdef USE_RANDOM_SEED
    srand(USE_RANDOM_SEED);
#endif
#endif

    if (argc == 2) {
        numVertices = std::stoi(argv[1]);
    }

    if (numVertices <= 0) {
        std::cerr << "Usage: " << argv[0] << " <num_vertices>" << std::endl;
        MPI_Finalize();
        return 1;
    }

    if (numProcesses > numVertices) {
        numProcesses = numVertices;

        if (myRank >= numProcesses) {
            MPI_Finalize();
            return 0;
        }
    }

    auto graph = createAndDistributeGraph(numVertices, numProcesses, myRank);
    if (graph == nullptr) {
        std::cerr << "Error creating the graph." << std::endl;
        MPI_Finalize();
        return 2;
    }

    collectAndPrintGraph(graph, numProcesses, myRank);
    destroyGraph(graph, numProcesses, myRank);

    MPI_Finalize();

    return 0;
}

