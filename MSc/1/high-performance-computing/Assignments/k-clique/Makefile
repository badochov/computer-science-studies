NVCC ?= /usr/local/cuda/bin/nvcc
CUDA_FLAGS ?= -arch=sm_70 -gencode=arch=compute_70,code=sm_70 -gencode=arch=compute_70,code=compute_70

all:
	make kcliques


kcliques: kcliques.cu
	${NVCC} -o kcliques -O3 ${CUDA_FLAGS} kcliques.cu

clean:
	-rm kcliques hb417666.zip
	-rm -rf hb417666

zip:
	mkdir -p hb417666
	cp kcliques.cu Makefile Report.md hb417666
	zip hb417666.zip -r hb417666
	rm -rf hb417666