.PHONY:
	build_jvm
	build_jvm
	all
	clean
	antlr

all:
	make build_jvm
	make build_llvm

build_jvm:
	go build -o insc_jvm src/jvm/main.go

build_llvm:
	go build -o insc_llvm src/llvm/main.go

antlr:
	antlr4 -Dlanguage=Go -o src/parser Instant.g4

tar:
	-rm hb417666.tar.gz
	tar -czvf hb417666.tar.gz Makefile go.mod go.sum vendor lib README.MD src

test:
	-rm -rf tools/projectdir
	mkdir tools/projectdir
	make tar
	mv hb417666.tar.gz tools/projectdir
	cd tools/instester && stack run ../projectdir ../examples


clean:
	rm insc_llvm insc_jvm hb417666.tar.gz