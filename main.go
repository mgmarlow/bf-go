package main

import (
	"github.com/mgmarlow/brainfuck/brain"
)

func main() {
	test := []byte("++++++ [ > ++++++++++ < - ] > +++++ .")
	vm := &brain.BrainVM{}
	vm.Interpret(test)
}
