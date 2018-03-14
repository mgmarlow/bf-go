package brain

import (
	"fmt"
	"os"
)

type BrainVM struct {
	Memory [30000]uint8
	DP     int
	IP     int
	buf    []byte
}

func (m *BrainVM) Interpret(instructions []byte) {
	m.clear()

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
				depth := 1
				for depth != 0 {
					m.IP++
					if instructions[m.IP] == '[' {
						depth++
					} else if instructions[m.IP] == ']' {
						depth--
					}
				}
			}
		case ']':
			if m.Memory[m.DP] != 0 {
				depth := 1
				for depth != 0 {
					m.IP--
					if instructions[m.IP] == ']' {
						depth++
					} else if instructions[m.IP] == '[' {
						depth--
					}
				}
			}

		case '.':
			fmt.Printf("%c", m.Memory[m.DP])
		case ',':
			m.readChar()

		default:
		}

		m.IP++
	}
}

func (m *BrainVM) clear() {
	m.Memory = [30000]uint8{}
	m.DP = 0
	m.IP = 0
	m.buf = make([]byte, 1)
}

// TODO: input from file?
func (m *BrainVM) readChar() {
	n, err := os.Stdin.Read(m.buf)
	if err != nil {
		panic(err)
	}

	if n != 1 {
		// TODO: better error handling
		panic("Please only enter one character")
	}

	m.Memory[m.DP] = uint8(m.buf[0])
}
