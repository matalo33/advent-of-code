package main

import (
  "bufio"
  "fmt"
  "log"
  "math"
  "os"
  "strconv"
  "strings"
)

type Vect struct {
  x, y int
}

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  reader := bufio.NewScanner(file)

  for reader.Scan() {
    line1 := reader.Text()
    reader.Scan() // 2 lines per scenario
    line2 := reader.Text()
    fmt.Printf("Manhattan Distance: %v", manhattanDistance(line1, line2))
  }
}

func manhattanDistance(wire1, wire2 string) int {
  grid := make(map[Vect]int)
  pos := Vect{}
  nearestManhattan := math.MaxInt32

  for _, section := range strings.Split(wire1, ",") {
    switch section[0] {
    case 'U':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y + 1}
        grid[pos] = 1
      }
    case 'D':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y - 1}
        grid[pos] = 1
      }
    case 'L':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x - 1, pos.y}
        grid[pos] = 1
      }
    case 'R':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x + 1, pos.y}
        grid[pos] = 1
      }
    }
  }

  pos = Vect{}
  for _, section := range strings.Split(wire2, ",") {
    switch section[0] {
    case 'U':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y + 1}
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
        }
      }
    case 'D':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y - 1}
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
        }
      }
    case 'L':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x - 1, pos.y}
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
        }
      }
    case 'R':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x + 1, pos.y}
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
        }
      }
    }
  }
  return nearestManhattan
}

func strToInt (str string) int {
  result, _ := strconv.Atoi(str)
  return result
}

func getDistance(vect Vect) int {
  return abs(vect.x) + abs(vect.y)
}

func abs (i int) int {
  if i < 0 {
    return -i
  }
  return i
}

func min (a, b int) int {
  if a < b {
    return a
  }
  return b
}