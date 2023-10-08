/**
 * stencil.cu: a simple 1d stencil on GPU and on CPU
 * 
 * Implement the basic stencil and make sure it works correctly.
 * Then, play with the code
 * - Experiment with block sizes, various RADIUSes and NUM_ELEMENTS.
 * - Measure the memory transfer time, estimate the effective memory bandwidth.
 * - Estimate FLOPS (floating point operations per second)
 * - Switch from float to double: how the performance changes?
*/

#include <time.h>
#include <stdio.h>
#include <cassert>

#include <algorithm>

#define RADIUS 3
#define NUM_ELEMENTS int(1e8)
#define THREADS_PER_BLOCK int(1024)

#define s_t double

#define BYTES (NUM_ELEMENTS * sizeof(s_t))

static void handleError(cudaError_t err, const char *file, int line)
{
  if (err != cudaSuccess)
  {
    printf("%s in %s at line %d\n", cudaGetErrorString(err), file, line);
    exit(EXIT_FAILURE);
  }
}
#define cudaCheck(err) (handleError(err, __FILE__, __LINE__))

__global__ void stencil_1d(s_t *in, s_t *out)
{
  int idx = blockIdx.x * THREADS_PER_BLOCK + threadIdx.x;
  if (idx >= NUM_ELEMENTS)
  {
    return;
  }

  out[idx] = 0;
  for (int j = max(0, idx - RADIUS); j < min(idx + RADIUS + 1, NUM_ELEMENTS); j++)
  {
    out[idx] += in[j];
  }
}

void cpu_stencil_1d(s_t *in, s_t *out)
{
  for (int i = 0; i < NUM_ELEMENTS; i++)
  {
    out[i] = 0;
    for (int j = max(0, i - RADIUS); j < min(i + RADIUS + 1, NUM_ELEMENTS); j++)
    {
      out[i] += in[j];
    }
  }
}

int main()
{
  s_t *in = (s_t *)calloc(NUM_ELEMENTS, sizeof(s_t));
  s_t *out = (s_t *)calloc(NUM_ELEMENTS, sizeof(s_t));
  s_t *outd = (s_t *)calloc(NUM_ELEMENTS, sizeof(s_t));
  s_t *devIn, *devOut;

  for (int i = 0; i < NUM_ELEMENTS; i++)
  {
    in[i] = i - NUM_ELEMENTS / 2;
  }

  cudaEvent_t start, stop;
  cudaEvent_t startTran, stopTran;
  cudaEvent_t startTranFrom, stopTranFrom;
  cudaCheck(cudaEventCreate(&start));
  cudaCheck(cudaEventCreate(&stop));
  cudaCheck(cudaEventCreate(&startTran));
  cudaCheck(cudaEventCreate(&stopTran));
  cudaCheck(cudaEventCreate(&startTranFrom));
  cudaCheck(cudaEventCreate(&stopTranFrom));
  cudaCheck(cudaEventRecord(start, 0));

  cudaCheck(cudaMalloc((void **)&devIn, NUM_ELEMENTS * sizeof(s_t)));
  cudaCheck(cudaMalloc((void **)&devOut, NUM_ELEMENTS * sizeof(s_t)));

  cudaCheck(cudaEventRecord(startTran, 0));
  cudaCheck(cudaMemcpy(devIn, in, NUM_ELEMENTS * sizeof(s_t), cudaMemcpyHostToDevice));
  cudaCheck(cudaEventRecord(stopTran, 0));
  cudaCheck(cudaEventSynchronize(stopTran));

  stencil_1d<<<NUM_ELEMENTS / THREADS_PER_BLOCK + 1, THREADS_PER_BLOCK>>>(devIn, devOut);
  cudaCheck(cudaPeekAtLastError());

  cudaCheck(cudaEventRecord(startTranFrom, 0));
  
  cudaCheck(cudaMemcpy(outd, devOut, NUM_ELEMENTS * sizeof(s_t), cudaMemcpyDeviceToHost));
  cudaCheck(cudaEventRecord(stopTranFrom, 0));
  cudaCheck(cudaEventSynchronize(stopTranFrom));

  cudaCheck(cudaEventRecord(stop, 0));
  cudaCheck(cudaEventSynchronize(stop));

  float gpuTranToElapsedTime;
  cudaEventElapsedTime(&gpuTranToElapsedTime, startTran, stopTran);
  printf("Total GPU transfer to time:  %.4f ms, %.4f GB/s\n", gpuTranToElapsedTime, BYTES / gpuTranToElapsedTime / 1e6);
  cudaCheck(cudaEventDestroy(startTran));
  cudaCheck(cudaEventDestroy(stopTran));

  float gpuTranFromElapsedTime;
  cudaEventElapsedTime(&gpuTranFromElapsedTime, startTranFrom, stopTranFrom);
  printf("Total GPU transfer from time:  %.4f ms, %.4f GB/s\n", gpuTranFromElapsedTime, BYTES / gpuTranFromElapsedTime /1e6 );
  cudaCheck(cudaEventDestroy(startTranFrom));
  cudaCheck(cudaEventDestroy(stopTranFrom));


  float gpuElapsedTime;
  cudaEventElapsedTime(&gpuElapsedTime, start, stop);
  printf("Total GPU execution time:  %.4f ms, %.4f GFLOPS\n", gpuElapsedTime, NUM_ELEMENTS * RADIUS / (gpuElapsedTime - gpuTranFromElapsedTime - gpuTranToElapsedTime) / 1e6);
  cudaCheck(cudaEventDestroy(start));
  cudaCheck(cudaEventDestroy(stop));

  cudaCheck(cudaFree(devIn));
  cudaCheck(cudaFree(devOut));

  struct timespec cpu_start, cpu_stop;
  clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &cpu_start);

  cpu_stencil_1d(in, out);

  clock_gettime(CLOCK_PROCESS_CPUTIME_ID, &cpu_stop);
  double cpuElapsedTime = (cpu_stop.tv_sec - cpu_start.tv_sec) * 1e3 + (cpu_stop.tv_nsec - cpu_start.tv_nsec) / 1e6;
  printf("CPU execution time:  %.4f ms\n", cpuElapsedTime);

  for (int i = 0; i < NUM_ELEMENTS; i++)
  {
    if (abs(outd[i] - out[i]) > 1e-6)
    {
      printf("%d %f %f\n", i, outd[i], out[i]);
      return 1;
    }
  }

  free(in);
  free(out);
  free(outd);

  return 0;
}
