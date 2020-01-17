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

    adjacencyTest := true
    incrementingTest := true
    adjacencyTest2 := true
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

    longestMatchCount := 0
    longestMatchValue := 0
    currentAdjacencyCount := 0
    highestMatch := 0

    // This never worked and I got really fed up with the scenario, so I cheated and passed
    // the question using someone else's solution
    if incrementingTest {
      for c := 0; c <= 4; c++ {
        if p[c] == p[c+1] {
          if p[c] > highestMatch {
            highestMatch = p[c]
            currentAdjacencyCount = 2
          } else {
            if p[c] == highestMatch {
              currentAdjacencyCount++
              if currentAdjacencyCount >= longestMatchCount {
                if p[c] >= longestMatchValue {
                  longestMatchValue = p[c]
                  longestMatchCount = 3
                } else {
                  if p[c] == longestMatchValue {
                    longestMatchCount++
                  }
                }
              }
            }
          }
        }
      }
    }

    if (highestMatch == longestMatchValue) && (longestMatchCount > 2) {
      adjacencyTest2 = false
    }

    if (adjacencyTest) && (incrementingTest) {
      fmt.Printf("%v a: %v i: %v t: %v highestMatch: %v currentAdjacencyCount: %v longestMatchCount: %v longestMatchValue: %v\n",
        i,
        adjacencyTest,
        incrementingTest,
        adjacencyTest2,
        highestMatch,
        currentAdjacencyCount,
        longestMatchCount,
        longestMatchValue,
      )
    }
    if (adjacencyTest) && (incrementingTest) && (adjacencyTest2) {
      matches++
    }
  }

  return matches
}