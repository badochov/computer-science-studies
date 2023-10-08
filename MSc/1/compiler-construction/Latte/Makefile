.PHONY:
	build
	test
	all
	clean
	antlr

all:
	make build

build:
	make build_runtime
	go build -o latc_llvm src/main.go

build_runtime:
	llvm-as -o lib/runtime.bc lib/runtime.ll

antlr:
	antlr4 -Dlanguage=Go -visitor -o src/parser Latte.g4

tar:
	-rm hb417666.tar.gz
	go mod vendor
	tar -czvf hb417666.tar.gz Makefile go.mod go.sum vendor README.MD src lib

clean_test:
	find tools/rs_tester -name "*.ll" -delete
	find tools/rs_tester -name "*.bc" -delete
	rm -rf tools/rs_tester/lib
	rm tools/rs_tester/latc_llvm

test:
	-make clean_test
	make build && mv ./latc_llvm tools/rs_tester
	make build_runtime && cp -r lib/ tools/rs_tester
	cd tools/rs_tester && cargo test

test_unofficial:
	-rm -rf tools/projectdir
	mkdir tools/projectdir
	make tar
	mv hb417666.tar.gz tools/projectdir
	cd tools/lattester && cabal run lattester ../projectdir ../mrjp-tests

test_front:
	make build
	mv latc tools/front_tester
	cd tools/front_tester && ./test.sh latc
	rm tools/front_tester/latc


clean:
	rm latc latc_llvm hb417666.tar.gz