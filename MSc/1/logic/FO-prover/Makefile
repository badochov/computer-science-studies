.PHONY: all

all: FO-prover

FO-prover: FO-prover.cabal *.hs
	cabal install --overwrite-policy=always --installdir . -O2
clean:
	-rm FO-prover
	-rm -rf dist-newstyle
	-rm hb417666.zip

zip:
	make clean
	zip hb417666.zip -r .