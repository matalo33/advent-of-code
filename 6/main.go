package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

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
  orbitalTransfers := orbitalTransfer(orbits)

  fmt.Printf("Direct and indirect orbits: %v\n", totalOrbits)
  fmt.Printf("Orbital Transfers moving YOU to SAN: %v", orbitalTransfers)
}

func loadData(inputLines []string) map[string]string {
  orbits := make(map[string]string)
  for _, line := range inputLines {
    orbits[strings.Split(line, ")")[1]] = strings.Split(line, ")")[0]
  }
  return orbits
}

func countOrbits(orbits map[string]string) int {
  totalOrbits := 0
  for orbit := range orbits {
    for {
      star, keyFound := orbits[orbit]

      if !keyFound {
       break
      }
      // This is another way of doing this check
      // The 2nd output from a map lookup is a boolean about whether the key was found or not
      //if orbit == "COM" {
      //  break
      //}

      orbit = star
      totalOrbits++
    }
  }
  return totalOrbits
}

func orbitalTransfer(orbits map[string]string) int {

  // Make a path from YOU to COM
  youToCom := make(map[string]int)
  orbit, distance := orbits["YOU"], 0
  for {
    youToCom[orbit] = distance
    star, keyFound := orbits[orbit]

    if !keyFound {
      break
    }

    orbit = star
    distance++
  }

  // Make a path from SAN to COM until crossing paths with YOU to COM
  orbit, distance = orbits["SAN"], 0
  for {
    // If this path intersects with youToCom
    if youDistance, ok := youToCom[orbit]; ok {
      return distance + youDistance
    }
    // Else continue walking sanToCom
    star := orbits[orbit]
    orbit = star
    distance++
  }
}