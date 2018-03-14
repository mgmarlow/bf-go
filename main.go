package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mgmarlow/brainfuck/brain"
)

func main() {
	bfFile, err := ioutil.ReadFile("./mandelbrot.bf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading input file: %v", err)
	}

	vm := &brain.BrainVM{}
	vm.Interpret(bfFile)
}
