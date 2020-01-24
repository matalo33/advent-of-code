package main

import (
  "reflect"
  "testing"
)

func TestConvertStrArrayToIntArray(t *testing.T) {
  input := []string {"1", "2", "999"}
  want := []int {1, 2, 999}

  got := convertStrArrayToIntArray(input)
  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestIntcodeMachine(t *testing.T) {
  opcodeTests := []struct {
    input []int
    output []int
  }{
    {[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
    {[]int{2,3,0,3,99}, []int{2,3,0,6,99}},
    {[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}},
    {[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
  }

  for _, opcode := range opcodeTests {
    got := intcodeMachine(opcode.input)
    if !reflect.DeepEqual(got, opcode.output) {
      t.Errorf("got %v want %v", got, opcode.output)
    }
  }
}