package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Orbit struct {
  star, planet string
}

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var inputLines []string

  for scanner.Scan() {
    inputLines = append(inputLines, scanner.Text())
  }

  orbits := loadData(inputLines)
  totalOrbits := countOrbits(orbits)

  fmt.Printf("Direct and indirect orbits: %v\n", totalOrbits)
}

func loadData(inputLines []string) []Orbit {
  var orbits []Orbit
  for _, line := range inputLines {
    orbits = append(orbits, Orbit{
      strings.Split(line, ")")[0],
      strings.Split(line, ")")[1],
    })
  }
  return orbits
}

func countOrbits(orbits []Orbit) int {

  return 0
}