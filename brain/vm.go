package brain

import (
	"fmt"
)

type BrainVM struct {
	Memory [30000]uint8
	DP     int
	IP     int
}

func (m *BrainVM) Interpret(instructions []byte) {
	for m.IP < len(instructions) {
		switch instructions[m.IP] {
		case '+':
			m.Memory[m.DP]++
		case '-':
			m.Memory[m.DP]--

		case '>':
			m.DP++
		case '<':
			m.DP--

		case '[':
		case ']':

		case '.':
			fmt.Println(m.Memory[m.DP])
		case ',':
			// TODO

		default:
		}

		m.IP++
	}
}
