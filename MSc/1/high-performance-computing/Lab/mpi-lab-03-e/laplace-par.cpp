/*
 * A template for the 2019 MPI lab at the University of Warsaw.
 * Copyright (C) 2016, Konrad Iwanicki.
 * Refactoring 2019, Łukasz Rączkowski
 */

#include <iostream>
#include <iomanip>
#include <cstring>
#include <sys/time.h>
#include <mpi.h>
#include "laplace-common.h"

#define OPTION_VERBOSE "--verbose"

static void printUsage(char const *progName)
{
    std::cerr << "Usage:" << std::endl
              << "    " << progName << " [--verbose] <N>" << std::endl
              << "Where:" << std::endl
              << "   <N>         The number of points in each dimension (at least 4)." << std::endl
              << "   " << OPTION_VERBOSE << "   Prints the input and output systems." << std::endl;
}

static InputOptions parseInput(int argc, char *argv[], int numProcesses)
{
    int numPointsPerDimension = 0;
    bool verbose = false;
    int errorCode = 0;

    if (argc < 2)
    {
        std::cerr << "ERROR: Too few arguments!" << std::endl;
        printUsage(argv[0]);
        errorCode = 1;
        MPI_Finalize();
    }
    else if (argc > 3)
    {
        std::cerr << "ERROR: Too many arguments!" << std::endl;
        printUsage(argv[0]);
        errorCode = 2;
        MPI_Finalize();
    }
    else
    {
        int argIdx = 1;

        if (argc == 3)
        {
            if (strncmp(argv[argIdx], OPTION_VERBOSE, strlen(OPTION_VERBOSE)) != 0)
            {
                std::cerr << "ERROR: Unexpected option '" << argv[argIdx] << "'!" << std::endl;
                printUsage(argv[0]);
                errorCode = 3;
                MPI_Finalize();
            }
            verbose = true;
            ++argIdx;
        }

        numPointsPerDimension = std::strtol(argv[argIdx], nullptr, 10);

        if ((numPointsPerDimension < 4) || (numProcesses > numPointsPerDimension / 2))
        {
            /* If we had a smaller grid, we could use the sequential version. */
            std::cerr << "ERROR: The number of points, '"
                      << argv[argIdx]
                      << "', should be an iteger greater than or equal to 4; and at least 2 points per process!"
                      << std::endl;
            printUsage(argv[0]);
            MPI_Finalize();
            errorCode = 4;
        }
    }

    return {numPointsPerDimension, verbose, errorCode};
}

#define SEND_UP 0
#define SEND_DOWN 1
#define RECV_UP 2
#define RECV_DOWN 3

enum
{
    SEND_UP1 = 0,
    SEND_UP2,
    SEND_DOWN1,
    SEND_DOWN2,
    RECV_UP1,
    RECV_UP2,
    RECV_DOWN1,
    RECV_DOWN2,
    _REQ_SIZE
};


#define debug printf("%d, %s:%d\n", myRank, __FILE__, __LINE__);

static std::tuple<int, double> performAlgorithm(
    int myRank, int numProcesses, GridFragment *frag, double omega, double epsilon)
{
    bool isFirst = myRank == 0;
    bool isLast = myRank == numProcesses - 1;

    int startRowIncl = frag->firstRowIdxIncl + (isFirst ? 1 : 0);
    int endRowExcl = frag->lastRowIdxExcl - (isLast ? 1 : 0);

    double maxDiff = 0;
    int numIterations = 0;

    MPI_Request requests[_REQ_SIZE];

    /* TODO: change the following code fragment */
    /* Implement asynchronous communication of neighboring elements */
    /* and computation of the grid */
    /* the following code just recomputes the appropriate grid fragment */
    /* but does not communicate the partial results */
    size_t sendSize = (frag->gridDimension + 1) / 2;
    do
    {
        printf("ITERATION %d\n", myRank);
        maxDiff = 0.0;
        if (!isFirst)
        {
debug
            MPI_Isend(
                &frag->data[0][1],
                sendSize,
                MPI_DOUBLE,
                myRank - 1,
                SEND_UP1,
                MPI_COMM_WORLD,
                &requests[SEND_UP1]);
debug
            MPI_Isend(
                &frag->data[1][1],
                sendSize,
                MPI_DOUBLE,
                myRank - 1,
                SEND_UP2,
                MPI_COMM_WORLD,
                &requests[SEND_UP2]);
debug

            MPI_Irecv(
                &frag->data[0][0],
                sendSize,
                MPI_DOUBLE,
                myRank - 1,
                SEND_DOWN1,
                MPI_COMM_WORLD,
                &requests[RECV_UP1]);
debug
            MPI_Irecv(
                &frag->data[1][0],
                sendSize,
                MPI_DOUBLE,
                myRank - 1,
                SEND_DOWN2,
                MPI_COMM_WORLD,
                &requests[RECV_UP2]);
debug
        }
        if (!isLast)
        {
            auto lastRowIdx = frag->lastRowIdxExcl - frag->firstRowIdxIncl + 1;
            printf("LAST ROW %d %d %d %d\n", myRank, lastRowIdx, sendSize, frag->data[1][6][1]);
debug 
            MPI_Isend(
                &frag->data[0][lastRowIdx - 1],
                sendSize,
                MPI_DOUBLE,
                myRank + 1,
                SEND_DOWN1,
                MPI_COMM_WORLD,
                &requests[SEND_DOWN1]);
debug
            MPI_Isend(
                &frag->data[1][lastRowIdx - 1],
                sendSize,
                MPI_DOUBLE,
                myRank + 1,
                SEND_DOWN2,
                MPI_COMM_WORLD,
                &requests[SEND_DOWN2]);
debug

            MPI_Irecv(
                &frag->data[0][lastRowIdx],
                sendSize,
                MPI_DOUBLE,
                myRank + 1,
                SEND_UP1,
                MPI_COMM_WORLD,
                &requests[RECV_DOWN1]);
debug
            MPI_Irecv(
                &frag->data[1][lastRowIdx],
                sendSize,
                MPI_DOUBLE,
                myRank + 1,
                SEND_UP2,
                MPI_COMM_WORLD,
                &requests[RECV_DOWN2]);
debug
        }

        for (int color = 0; color < 2; ++color)
        {
debug
            for (int rowIdx = startRowIncl; rowIdx < endRowExcl; ++rowIdx)
            {
debug
                if (rowIdx == startRowIncl && !isFirst) // First row needed
                {
                    debug
                    MPI_Request reqs[] = {requests[RECV_UP1], requests[RECV_UP2]};
                    debug

                    MPI_Waitall(2, reqs, MPI_STATUS_IGNORE);debug
debug
                }
                if (rowIdx == endRowExcl - 1 && !isLast) // Last row needed
                {
                    debug
                    MPI_Request reqs[] = {requests[RECV_DOWN1], requests[RECV_DOWN2]};
                    debug

                    MPI_Waitall(2, reqs, MPI_STATUS_IGNORE);debug
                }
debug
                for (int colIdx = 1 + (rowIdx % 2 == color ? 1 : 0);
                     colIdx < frag->gridDimension - 1;
                     colIdx += 2)
                {
debug
printf("INFO1 %d %d %d %d %d\n", myRank, isFirst, isLast, rowIdx - 1, (rowIdx - 1) - (frag)->firstRowIdxIncl + 1);
printf("INFO2 %d %d %d %d %d\n", myRank, isFirst, isLast, rowIdx + 1, (rowIdx + 1) - (frag)->firstRowIdxIncl + 1);
printf("INFO3 %d %d %d %d\n",((rowIdx + 1) + (colIdx)) % 2, (rowIdx + 1) - (frag)->firstRowIdxIncl + 1,  (colIdx) / 2, color);
printf("INFO4\n", GP(frag, rowIdx + 1, colIdx));
volatile int a = GP(frag, rowIdx + 1, colIdx);debug
debug
a = GP(frag, rowIdx - 1, colIdx);
debug
a = GP(frag, rowIdx, colIdx - 1);
debug
a = GP(frag, rowIdx, colIdx + 1);
debug
                    double tmp =
                        (GP(frag, rowIdx - 1, colIdx) +
                         GP(frag, rowIdx + 1, colIdx) +
                         GP(frag, rowIdx, colIdx - 1) +
                         GP(frag, rowIdx, colIdx + 1)) /
                        4.0;
debug
                    double diff = GP(frag, rowIdx, colIdx);
debug
                    GP(frag, rowIdx, colIdx) = (1.0 - omega) * diff + omega * tmp;
debug
                    diff = fabs(diff - GP(frag, rowIdx, colIdx));

debug
                    if (diff > maxDiff)
                    {
debug
                        maxDiff = diff;
                    }
                }
            }
        }
debug
        ++numIterations;

debug
        MPI_Request waitSendReqs[4];
debug
        int waitSendCount = 0;
debug
        if (myRank != 0)
        {
debug
            waitSendCount += 2;
debug
            waitSendReqs[0] = requests[SEND_UP1];
debug
            waitSendReqs[1] = requests[SEND_UP2];
debug
        }
        if (myRank != numProcesses - 1)
        {
debug
            waitSendReqs[waitSendCount] = requests[SEND_UP1];
debug
            waitSendReqs[waitSendCount + 1] = requests[SEND_UP2];
debug
            waitSendCount += 2;
debug
        }
        if (waitSendCount != 0)
        {
debug
            MPI_Waitall(waitSendCount, waitSendReqs, MPI_STATUSES_IGNORE);
        }
debug
    } while (maxDiff > epsilon);
debug
    /* no code changes beyond this point should be needed */
    return std::make_tuple(numIterations, maxDiff);
}

int main(int argc, char *argv[])
{
    int numProcesses;
    int myRank;
    struct timeval startTime
    {
    };
    struct timeval endTime
    {
    };

    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &numProcesses);
    MPI_Comm_rank(MPI_COMM_WORLD, &myRank);

    auto inputOptions = parseInput(argc, argv, numProcesses);
    if (inputOptions.getErrorCode() != 0)
    {
        return inputOptions.getErrorCode();
    }

    auto numPointsPerDimension = inputOptions.getNumPointsPerDimension();
    auto isVerbose = inputOptions.isVerbose();

    double omega = Utils::getRelaxationFactor(numPointsPerDimension);
    double epsilon = Utils::getToleranceValue(numPointsPerDimension);

    auto gridFragment = new GridFragment(numPointsPerDimension, numProcesses, myRank);
    gridFragment->initialize();

    if (gettimeofday(&startTime, nullptr))
    {
        gridFragment->free();
        std::cerr << "ERROR: Gettimeofday failed!" << std::endl;
        MPI_Finalize();
        return 6;
    }

    /* Start of computations. */
    auto result = performAlgorithm(myRank, numProcesses, gridFragment, omega, epsilon);
    /* End of computations. */

    if (gettimeofday(&endTime, nullptr))
    {
        gridFragment->free();
        std::cerr << "ERROR: Gettimeofday failed!" << std::endl;
        MPI_Finalize();
        return 7;
    }
    double duration =
        ((double)endTime.tv_sec + ((double)endTime.tv_usec / 1000000.0)) -
        ((double)startTime.tv_sec + ((double)startTime.tv_usec / 1000000.0));
    std::cerr << "Statistics: duration(s)="
              << std::fixed
              << std::setprecision(10)
              << duration << " #iters="
              << std::get<0>(result)
              << " diff="
              << std::get<1>(result)
              << " epsilon="
              << epsilon
              << std::endl;
    if (isVerbose)
    {
        gridFragment->printEntireGrid(myRank, numProcesses);
    }
    gridFragment->free();
    MPI_Finalize();
    return 0;
}
