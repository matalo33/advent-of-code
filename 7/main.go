package main

import (
"bufio"
"encoding/csv"
"fmt"
"log"
"os"
"strconv"
)

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
  signal := calculateThrustSignal(opcode)

  fmt.Printf("Thruster signal: %v", signal)

}

func convertStrArrayToIntArray(input []string) []int {
  var result []int
  for _, s := range input {
    i, _ := strconv.Atoi(s)
    result = append(result, i)
  }
  return result
}

func calculateThrustSignal(opcode []int) (thrustSignal int) {
  signals := permutations([]int{0,1,2,3,4})
  ampOutput, highAmpOutput := 0, 0
  for _, signal := range signals {
    ampOutput = 0
    for amp := 0; amp <=4; amp++ {
      //fmt.Printf("Running on code %v, amp %v, phase %v, input %v\n", signal, amp, signal[amp], ampOutput)
      ampOutput = intcodeMachine(opcode, []int{signal[amp], ampOutput})
    }
    if ampOutput > highAmpOutput {
      highAmpOutput = ampOutput
    }
  }
  return highAmpOutput
}

func intcodeMachine(op, input []int) (output int) {
  pc := 0
  memory := make([]int, len(op))
  copy(memory, op)

  for {
    opcode := memory[pc] % 100

    getParam := func(pos int) int {
      parameter := memory[pc+pos]
      paramCode := reverse(strconv.Itoa(memory[pc]))
      pos += 1 // opcode is 2 chars long, pos is 1 indexed

      mode := "0"
      if len(paramCode) > pos {
        mode = string(paramCode[pos])
      }

      switch mode {
      case "0":
        return memory[parameter]
      case "1":
        return parameter
      default:
        panic("foo")
      }
    }

    switch opcode {
    case 1: // ADD
      //fmt.Printf("CODE 1: %v, 1: %v, 2: %v, 3: %v\n", memory[pc], memory[pc+1], memory[pc+2], memory[pc+3])
      a, b := getParam(1), getParam(2)
      c := memory[pc+3]
      memory[c] = a + b
      pc += 4

    case 2: // MULTIPLY
      //fmt.Printf("CODE 2: %v, 1: %v, 2: %v, 3: %v\n", memory[pc], memory[pc+1], memory[pc+2], memory[pc+3])
      a, b := getParam(1), getParam(2)
      c := memory[pc+3]
      memory[c] = a * b
      pc += 4

    case 3: // INPUT
      a := memory[pc+1]
      memory[a] = input[0]
      // Read it first time then pop 2nd input into 1st position
      input[0] = input[1]
      pc += 2

    case 4: // OUTPUT
      a := getParam(1)
      //output = append(output, a)
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
      c := memory[pc+3]
      if a < b {
        memory[c] = 1
      } else {
        memory[c] = 0
      }
      pc += 4

    case 8: // EQUAL
      a, b := getParam(1), getParam(2)
      c := memory[pc+3]
      if a == b {
        memory[c] = 1
      } else {
        memory[c] = 0
      }
      pc += 4

    case 99:
      return output

    default:
      log.Fatalf("OOPS %v", memory[pc])
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
func permutations(arr []int)[][]int{
  var helper func([]int, int)
  res := [][]int{}

  helper = func(arr []int, n int){
    if n == 1{
      tmp := make([]int, len(arr))
      copy(tmp, arr)
      res = append(res, tmp)
    } else {
      for i := 0; i < n; i++{
        helper(arr, n - 1)
        if n % 2 == 1{
          tmp := arr[i]
          arr[i] = arr[n - 1]
          arr[n - 1] = tmp
        } else {
          tmp := arr[0]
          arr[0] = arr[n - 1]
          arr[n - 1] = tmp
        }
      }
    }
  }
  helper(arr, len(arr))
  return res
}

