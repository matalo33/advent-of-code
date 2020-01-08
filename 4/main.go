package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()
  reader := bufio.NewReader(file)
  line, _ := reader.ReadString('\n')
  possibleRange := strings.Split(line, "-")

  countPossiblePasswords := countPossiblePasswords(possibleRange[0], possibleRange[1])
  fmt.Printf("Possible passwords: %v\n", countPossiblePasswords)
}

func countPossiblePasswords(strLower, strUpper string) int {
  min, _ := strconv.Atoi(strLower)
  max, _ := strconv.Atoi(strUpper)
  matches := 0

  fmt.Printf("Number of considerations: %v\n", max-min)

  for i := min; i < max+1; i++ {
    p := make([]int, 6)
    s := strings.Split(strconv.Itoa(i), "")
    for d := 0; d <= 5; d++ {
      p[d], _ = strconv.Atoi(s[d])
    }

    adjacencyTest := false
    incrementingTest := true
    for c := 0; c <= 4; c++ {
      if p[c] == p[c+1] {
        adjacencyTest = true
      }
    }

    for c := 0; c <= 4; c++ {
      if p[c] > p[c+1] {
        incrementingTest = false
      }
    }

    fmt.Printf("%v a: %v i: %v\n", i, adjacencyTest, incrementingTest)
    if (adjacencyTest) && (incrementingTest) {
      matches++
    }
  }

  return matches
}