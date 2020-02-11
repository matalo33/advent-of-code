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

  // Part 1
  result := intcodeMachine(opcode, []int{1})
  fmt.Printf("Part 1 output: %v\n", result)
}

func convertStrArrayToIntArray(input []string) []int {
  var result []int
  for _, s := range input {
    i, _ := strconv.Atoi(s)
    result = append(result, i)
  }
  return result
}

func intcodeMachine(op, input []int) (output []int) {
  pc := 0
  memory := make([]int, len(op))
  copy(memory, op)

  for {
    fmt.Printf("PC: %v\n", pc)
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
        fmt.Printf("paramcode: %v, pos: %v, position value: %v\n", paramCode, pos, memory[parameter])
        return memory[parameter]
      case "1":
        fmt.Printf("paramcode: %v, pos: %v, immediate value: %v\n", paramCode, pos, parameter)
        return parameter
      default:
        fmt.Printf("paramcode: %v, pos: %v, PANIC mode: %v\n", paramCode, pos, mode)
        panic("foo")
      }
    }

    switch opcode {
    case 1:
      fmt.Printf("CODE 1: %v, 1: %v, 2: %v, 3: %v\n", memory[pc], memory[pc+1], memory[pc+2], memory[pc+3])
      a, b := getParam(1), getParam(2)
      c := memory[pc+3]
      memory[c] = a + b
      fmt.Printf("Storing %v + %v (%v) at memory %v\n", a, b, a+b, c)
      pc += 4
    case 2:
      fmt.Printf("CODE 2: %v, 1: %v, 2: %v, 3: %v\n", memory[pc], memory[pc+1], memory[pc+2], memory[pc+3])
      a, b := getParam(1), getParam(2)
      c := memory[pc+3]
      memory[c] = a * b
      fmt.Printf("Storing %v + %v (%v) at memory %v\n", a, b, a+b, c)
      pc += 4
    case 3:
      a := memory[pc+1]
      memory[a] = input[0]
      pc += 2
    case 4:
      a := getParam(1)
      output = append(output, a)
      pc += 2
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