package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type amp struct {
	memory []int
	pc     int
	offset int
	debug  bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(file))
	line, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	opcode := convertStrArrayToIntArray(line)
	fmt.Printf("BOOST Keycode: %v", checkIntcodeMachine(opcode, 1, true))
}

func convertStrArrayToIntArray(input []string) []int {
	var result []int
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func checkIntcodeMachine(opcode []int, input int, debug bool) []int {
	memory := make([]int, math.MaxInt16)
	copy(memory, opcode)
	amp := amp{
		memory,
		0,
		0,
		debug,
	}

	result, _ := intcodeMachine(&amp, input)
	return result
}

func intcodeMachine(amp *amp, input int) (output []int, finished bool) {

	// Retrieve a parameter abiding by the current parameter mode
	getParam := func(pos int, rw string) int {
		// amp.pc is always the opcode
		// pos is the forward offset from amp.pc from which to get the value
		parameter := amp.memory[amp.pc+pos]

		// Now that we have parameter using pos, pos is now used to determine the mode of this parameter
		// pos is used as 1-indexed and the opcode is always 2 chars long
		// Add 1 to the pos and we have the position of this parameters mode
		pos++

		// Stringify the opcode and reverse it
		paramCode := reverse(strconv.Itoa(amp.memory[amp.pc]))

		// If not defined, parameter mode is 0
		mode := "0"

		// If defined, get the parameter mode
		if len(paramCode) > pos {
			mode = string(paramCode[pos])
		}

		// "Parameters that an instruction writes to will never be in immediate mode"
		if rw == "w" && mode == "1" {
			mode = "0"
		}

		switch mode {
		// POSITION MODE: The parameter is a position
		// Return the value stored in memory at the position
		case "0":
			if amp.debug {
				fmt.Printf("param: %v, rw: %v, mode: %v\n", pos-1, rw, mode)
			}
			return amp.memory[parameter]

		// IMMEDIATE MODE: The parameter is the value
		// Return the parameter
		case "1":
			if amp.debug {
				fmt.Printf("param: %v, rw: %v, mode: %v\n", pos-1, rw, mode)
			}
			return parameter

		// RELATIVE MODE: As position mode, but offset from the relative base
		// Return the value stored in memory offset by the relative offset
		case "2":
			if amp.debug {
				fmt.Printf("param: %v, rw: %v, mode: %v\n", pos-1, rw, mode)
			}
			return amp.memory[amp.offset+parameter]

		default:
			panic("Received an invalid parameter mode")
		}
	}

	//Infinite loop. Opcode 99 will exit
	for {

		// Two Rightmost (ones, tens) values in opcode make up the instruction
		opcode := amp.memory[amp.pc] % 100

		switch opcode {

		case 1: // ADD
			a, b, c := getParam(1, "r"), getParam(2, "r"), getParam(3, "w")
			if amp.debug {
				fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
				fmt.Printf("Opcode: ADD : %v\n\n", amp.memory[amp.pc:amp.pc+4])
			}
			amp.memory[c] = a + b
			amp.pc += 4

		case 2: // MULTIPLY
			a, b, c := getParam(1, "r"), getParam(2, "r"), getParam(3, "w")
			if amp.debug {
				fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
				fmt.Printf("Opcode: MULTIPLY : %v\n\n", amp.memory[amp.pc:amp.pc+4])
			}
			amp.memory[c] = a * b
			amp.pc += 4

		case 3: // INPUT
			a := getParam(1, "w")
			if amp.debug {
				fmt.Printf("a: %v\n", a)
				fmt.Printf("Opcode: INPUT : %v\n\n", amp.memory[amp.pc:amp.pc+2])
			}
			amp.memory[a] = input
			amp.pc += 2

		case 4: // OUTPUT
			a := getParam(1, "r")
			if amp.debug {
				fmt.Printf("a: %v\n", a)
				fmt.Printf("Opcode: OUTPUT : %v\n\n", amp.memory[amp.pc:amp.pc+2])
			}
			output = append(output, a)
			amp.pc += 2

		case 5: // JUMP IF TRUE
			a, b := getParam(1, "r"), getParam(2, "w")
			if amp.debug {
				fmt.Printf("a: %v, b: %v\n", a, b)
				fmt.Printf("Opcode: JTRUE : %v\n\n", amp.memory[amp.pc:amp.pc+3])
			}
			if a != 0 {
				amp.pc = b
			} else {
				amp.pc += 3
			}

		case 6: // JUMP IF FALSE
			a, b := getParam(1, "r"), getParam(2, "w")
			if amp.debug {
				fmt.Printf("a: %v, b: %v\n", a, b)
				fmt.Printf("Opcode: JFALSE : %v\n\n", amp.memory[amp.pc:amp.pc+3])
			}
			if a == 0 {
				amp.pc = b
			} else {
				amp.pc += 3
			}

		case 7: // LESS THAN
			a, b, c := getParam(1, "r"), getParam(2, "r"), getParam(3, "w")
			if amp.debug {
				fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
				fmt.Printf("Opcode: LESS : %v\n\n", amp.memory[amp.pc:amp.pc+4])
			}
			if a < b {
				amp.memory[c] = 1
			} else {
				amp.memory[c] = 0
			}
			amp.pc += 4

		case 8: // EQUAL
			a, b, c := getParam(1, "r"), getParam(2, "r"), getParam(3, "w")
			if amp.debug {
				fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
				fmt.Printf("Opcode: EQUAL : %v\n\n", amp.memory[amp.pc:amp.pc+4])
			}
			if a == b {
				amp.memory[c] = 1
			} else {
				amp.memory[c] = 0
			}
			amp.pc += 4

		case 9: // ADJUST OFFSET
			a := getParam(1, "r")
			if amp.debug {
				fmt.Printf("a: %v\n", a)
				fmt.Printf("Opcode: OFFSET : %v\n", amp.memory[amp.pc:amp.pc+2])
				fmt.Printf("Offset becomes %v\n\n", amp.offset+a)
			}
			amp.offset += a
			amp.pc += 2

		case 99:
			return output, true

		default:
			if amp.debug {
				fmt.Printf("OOPS: %v\n", opcode)
			}
			log.Fatalf("Tried to switch on an invalid opcode: %v\n", amp.memory[amp.pc])
		}
	}
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
