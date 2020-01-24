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

  // Restore computer to pre-crash
  opcode[1] = 12
  opcode[2] = 2
  result := intcodeMachine(opcode)
  fmt.Printf("Position 0: %v\n", result[0])

  op2 := convertStrArrayToIntArray(line)
  noun, verb := findInputForOutput(op2, 19690720)
  fmt.Printf("Noun %v, Verb %v\n", noun, verb)
}

func convertStrArrayToIntArray(input []string) []int {
  var result []int
  for _, s := range input {
    i, _ := strconv.Atoi(s)
    result = append(result, i)
  }
  return result
}

func intcodeMachine(op []int) []int {
  OpCode99:
  for p := 0; true; p += 4 {
    switch op[p] {
    case 99:
      break OpCode99
    case 1:
      op[op[p+3]] = op[op[p+1]] + op[op[p+2]]
    case 2:
      op[op[p+3]] = op[op[p+1]] * op[op[p+2]]
    default:
      log.Fatalf("OOPS %v", op[p])
    }
  }
  return op
}

func findInputForOutput(op []int, target int) (int, int) {
  for noun := 0; noun < len(op); noun++ {
    for verb := 0; verb < len(op); verb++ {
      newOp := make([]int, len(op))
      copy(newOp, op)
      newOp[1] = noun
      newOp[2] = verb
      if intcodeMachine(newOp)[0] == target {
        return noun, verb
      }
    }
  }
  return 0, 0
}