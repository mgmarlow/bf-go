package main

import "fmt"

var (
	memory   = [30000]int{}
	mpointer = 0
)

func main() {

}

func parseByte(bfChar byte) {
	switch bfChar {
	case '+':
		memory[mpointer]++
	case '-':
		memory[mpointer]--

	case '>':
		mpointer++
	case '<':
		mpointer--

	case '[':
	case ']':

	case '.':
		fmt.Println(rune(memory[mpointer]))
	case ',':

	default:
	}
}
