#include "cuda.h"
#include "common/errors.h"
#include "common/cpu_bitmap.h"

#define DIM 1024
#define rnd(x) (x * rand() / RAND_MAX)
#define INF 2e10f
#define SPHERES 100

texture<float, 1> t_red, t_green, t_blue, t_radius, t_x, t_y, t_z;

__device__ float hit(float red, float green, float blue, float radius, float x, float y, float z, float bitmapX, float bitmapY, float *colorFalloff)
{
	float distX = bitmapX - x;
	float distY = bitmapY - y;

	if (distX * distX + distY * distY < radius * radius)
	{
		float distZ = sqrtf(radius * radius - distX * distX - distY * distY);
		*colorFalloff = distZ / sqrtf(radius * radius);
		return distZ + z;
	}

	return -INF;
}

__global__ void kernel(unsigned char *bitmap)
{
	int x = threadIdx.x + blockIdx.x * blockDim.x;
	int y = threadIdx.y + blockIdx.y * blockDim.y;
	int offset = x + y * blockDim.x * gridDim.x;

	// tex

	float bitmapX = (x - DIM / 2);
	float bitmapY = (y - DIM / 2);

	float red = 0, green = 0, blue = 0;
	float maxDepth = -INF;

	for (int i = 0; i < SPHERES; i++)
	{
		float colorFalloff;
		float depth = hit(tex1D(t_red, i), tex1D(t_green, i), tex1D(t_blue, i), tex1D(t_radius, i), tex1D(t_x, i), tex1D(t_y, i), tex1D(t_z, i), bitmapX, bitmapY, &colorFalloff);

		if (depth > maxDepth)
		{
			red = tex1D(t_red, i) * colorFalloff;
			green = tex1D(t_green, i) * colorFalloff;
			blue = tex1D(t_blue, i) * colorFalloff;
			maxDepth = depth;
		}
	}

	bitmap[offset * 4 + 0] = (int)(red * 255);
	bitmap[offset * 4 + 1] = (int)(green * 255);
	bitmap[offset * 4 + 2] = (int)(blue * 255);
	bitmap[offset * 4 + 3] = 255;
}

struct DataBlock
{
	unsigned char *hostBitmap;
};

int main(void)
{
	DataBlock data;
	cudaEvent_t start, stop;
	HANDLE_ERROR(cudaEventCreate(&start));
	HANDLE_ERROR(cudaEventCreate(&stop));
	HANDLE_ERROR(cudaEventRecord(start, 0));

	CPUBitmap bitmap(DIM, DIM, &data);
	unsigned char *devBitmap;

	HANDLE_ERROR(cudaMalloc((void **)&devBitmap, bitmap.image_size()));

	float host_red[SPHERES];
	float host_blue[SPHERES];
	float host_green[SPHERES];
	float host_radius[SPHERES];
	float host_x[SPHERES];
	float host_y[SPHERES];
	float host_z[SPHERES];

	float *d_red, *d_green, *d_blue, *d_radius, *d_x, *d_y, *d_z;
	cudaChannelFormatDesc desc = cudaCreateChannelDesc<float>();
	cudaMalloc((void **)&d_red, SPHERES * sizeof(float));
	cudaMalloc((void **)&d_green, SPHERES * sizeof(float));
	cudaMalloc((void **)&d_blue, SPHERES * sizeof(float));
	cudaMalloc((void **)&d_radius, SPHERES * sizeof(float));
	cudaMalloc((void **)&d_x, SPHERES * sizeof(float));
	cudaMalloc((void **)&d_y, SPHERES * sizeof(float));
	cudaMalloc((void **)&d_z, SPHERES * sizeof(float));

	cudaBindTexture(0, t_red, d_red, desc, SPHERES * sizeof(float));
	cudaBindTexture(0, t_green, d_green, desc, SPHERES * sizeof(float));
	cudaBindTexture(0, t_blue, d_blue, desc, SPHERES * sizeof(float));
	cudaBindTexture(0, t_radius, d_radius, desc, SPHERES * sizeof(float));
	cudaBindTexture(0, t_x, d_x, desc, SPHERES * sizeof(float));
	cudaBindTexture(0, t_y, d_y, desc, SPHERES * sizeof(float));
	cudaBindTexture(0, t_z, d_z, desc, SPHERES * sizeof(float));

	for (int i = 0; i < SPHERES; i++)
	{
		host_red[i] = rnd(1.0f);
		host_green[i] = rnd(1.0f);
		host_blue[i] = rnd(1.0f);
		host_x[i] = rnd(1000.0f) - 500;
		host_y[i] = rnd(1000.0f) - 500;
		host_z[i] = rnd(1000.0f) - 500;
		host_radius[i] = rnd(100.0f) + 20;
	}

	cudaMemcpy(d_red, host_red, SPHERES * sizeof(float), cudaMemcpyHostToDevice);
	cudaMemcpy(d_green, host_green, SPHERES * sizeof(float), cudaMemcpyHostToDevice);
	cudaMemcpy(d_blue, host_blue, SPHERES * sizeof(float), cudaMemcpyHostToDevice);
	cudaMemcpy(d_radius, host_radius, SPHERES * sizeof(float), cudaMemcpyHostToDevice);
	cudaMemcpy(d_x, host_x, SPHERES * sizeof(float), cudaMemcpyHostToDevice);
	cudaMemcpy(d_y, host_y, SPHERES * sizeof(float), cudaMemcpyHostToDevice);
	cudaMemcpy(d_z, host_z, SPHERES * sizeof(float), cudaMemcpyHostToDevice);

	dim3 grids(DIM / 16, DIM / 16);
	dim3 threads(16, 16);
	kernel<<<grids, threads>>>(devBitmap);

	HANDLE_ERROR(cudaEventRecord(stop, 0));
	HANDLE_ERROR(cudaEventSynchronize(stop));

	float elapsedTime;
	HANDLE_ERROR(cudaEventElapsedTime(&elapsedTime, start, stop));
	printf("Time to generate: %3.1f ms\n", elapsedTime);

	HANDLE_ERROR(cudaEventDestroy(start));
	HANDLE_ERROR(cudaEventDestroy(stop));

	HANDLE_ERROR(cudaMemcpy(bitmap.get_ptr(), devBitmap, bitmap.image_size(), cudaMemcpyDeviceToHost));

	bitmap.dump_ppm("image-texture.ppm");

	cudaUnbindTexture(t_red);
	cudaUnbindTexture(t_green);
	cudaUnbindTexture(t_blue);
	cudaUnbindTexture(t_x);
	cudaUnbindTexture(t_y);
	cudaUnbindTexture(t_z);
	cudaUnbindTexture(t_radius);

	cudaFree(devBitmap);
}

// Global
