package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Amp struct {
	initialised bool
	signal      int
	memory      []int
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
	signal := calculateSeriesThrustSignal(opcode)
	//feedbackSignal := calculateFeedbackThrustSignal(opcode)

	fmt.Printf("Thruster signal: %v\n", signal)
	//fmt.Printf("Feedback mode thruster signal: %v\n", feedbackSignal)
}

func convertStrArrayToIntArray(input []string) []int {
	var result []int
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func calculateSeriesThrustSignal(opcode []int) (thrustSignal int) {
	signals := permutations([]int{0, 1, 2, 3, 4})
	return calculateThrustSignal(opcode, signals, false)
}

func calculateFeedbackThrustSignal(opcode []int) (thrustSignal int) {
	signals := permutations([]int{5, 6, 7, 8, 9})
	return calculateThrustSignal(opcode, signals, true)
}

func calculateThrustSignal(opcode []int, signals [][]int, feedbackMode bool) (thrustSignal int) {
	ampOutput, highAmpOutput := 0, 0

	for _, signal := range signals { // For each permutation of signal
		fmt.Printf("Run ")
		amps := make([]Amp, 5) // Create a new collection for amps
		for amp := 0; amp < 5; amp++ {
			memory := make([]int, len(opcode))
			copy(memory, opcode) // Crate memory for new amp
			// Add new amp to collection
			amps[amp] = Amp{
				false,
				signal[amp],
				memory,
			}
		}

		currentAmp := 0
		ampOutput = 0 // Reset last amp output
		for {         // Enter infinite loop
			ampFinished := false
			ampOutput, ampFinished = intcodeMachine(amps[currentAmp], ampOutput) // Run amp
			if ampOutput > highAmpOutput {
				highAmpOutput = ampOutput
			}
			fmt.Printf("currentAmp: %v, finished: %v, ampoutput: %v\n", currentAmp, ampFinished, ampOutput)
			if (currentAmp == 4) && (ampFinished) {
				break
			}
			currentAmp = (currentAmp + 1) % 5 // Increment by 1, capped at 5
		}
	}
	return highAmpOutput
}

func intcodeMachine(amp Amp, input int) (output int, finished bool) {
	pc, inputPc := 0, 0

	for {
		opcode := amp.memory[pc] % 100

		getParam := func(pos int) int {
			parameter := amp.memory[pc+pos]
			paramCode := reverse(strconv.Itoa(amp.memory[pc]))
			pos++ // opcode is 2 chars long, pos is 1 indexed

			mode := "0"
			if len(paramCode) > pos {
				mode = string(paramCode[pos])
			}

			switch mode {
			case "0":
				return amp.memory[parameter]
			case "1":
				return parameter
			default:
				panic("foo")
			}
		}

		switch opcode {
		case 1: // ADD
			a, b := getParam(1), getParam(2)
			c := amp.memory[pc+3]
			amp.memory[c] = a + b
			pc += 4

		case 2: // MULTIPLY
			a, b := getParam(1), getParam(2)
			c := amp.memory[pc+3]
			amp.memory[c] = a * b
			pc += 4

		case 3: // INPUT
			if !amp.initialised {
				a := amp.memory[pc+1]
				amp.memory[a] = amp.signal
				pc += 2
				amp.initialised = true
			} else {
				a := amp.memory[pc+1]
				amp.memory[a] = input
				inputPc++
				pc += 2
			}

		case 4: // OUTPUT
			a := getParam(1)
			output = a
			pc += 2

		case 5: // JUMP IF TRUE
			a, b := getParam(1), getParam(2)
			if a != 0 {
				pc = b
			} else {
				pc += 3
			}

		case 6: // JUMP IF FALSE
			a, b := getParam(1), getParam(2)
			if a == 0 {
				pc = b
			} else {
				pc += 3
			}

		case 7: // LESS THAN
			a, b := getParam(1), getParam(2)
			c := amp.memory[pc+3]
			if a < b {
				amp.memory[c] = 1
			} else {
				amp.memory[c] = 0
			}
			pc += 4

		case 8: // EQUAL
			a, b := getParam(1), getParam(2)
			c := amp.memory[pc+3]
			if a == b {
				amp.memory[c] = 1
			} else {
				amp.memory[c] = 0
			}
			pc += 4

		case 99:
			return output, true

		default:
			log.Fatalf("OOPS %v", amp.memory[pc])
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

// From https://stackoverflow.com/a/30226442
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
