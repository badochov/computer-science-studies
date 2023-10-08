package main

import (
	"github.com/badochov/MRJP/instant/src/common"
	"github.com/badochov/MRJP/instant/src/llvm/listener"
	"log"
)

func main() {
	compileFn := common.CreateCompileFunction(listener.New())
	if err := common.Main(compileFn); err != nil {
		log.Fatal(err)
	}
}
