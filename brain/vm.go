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
			if m.Memory[m.DP] == 0 {
				for i := m.IP; i < len(instructions); i++ {
					if instructions[i] == ']' {
						m.IP = i + 1
						break
					}
				}
			}
		case ']':
			if m.Memory[m.DP] != 0 {
				for i := m.IP; i > 0; i-- {
					if instructions[i] == '[' {
						m.IP = i + 1
						break
					}
				}
			}

		case '.':
			fmt.Printf("%c", m.Memory[m.DP])
		case ',':
			// TODO

		default:
		}

		m.IP++
	}
}
