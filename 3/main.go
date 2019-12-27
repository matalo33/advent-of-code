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
    manhattanDistance, shortestDistance := manhattanDistance(line1, line2)
    fmt.Printf("Manhattan Distance: %v, Shortest Distance: %v", manhattanDistance, shortestDistance)
  }
}

func manhattanDistance(wire1, wire2 string) (int, int) {
  grid := make(map[Vect]int)
  pos := Vect{}
  nearestManhattan, shortestDistance := math.MaxInt32, math.MaxInt32
  steps := 0

  for _, section := range strings.Split(wire1, ",") {
    switch section[0] {
    case 'U':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y + 1}
        steps++
        grid[pos] = steps
      }
    case 'D':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y - 1}
        steps++
        grid[pos] = steps
      }
    case 'L':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x - 1, pos.y}
        steps++
        grid[pos] = steps
      }
    case 'R':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x + 1, pos.y}
        steps++
        grid[pos] = steps
      }
    }
  }

  pos = Vect{}
  currentSteps := 0
  for _, section := range strings.Split(wire2, ",") {
    switch section[0] {
    case 'U':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y + 1}
        currentSteps++
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
          shortestDistance = min(currentSteps + grid[pos], shortestDistance)
        }
      }
    case 'D':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x, pos.y - 1}
        currentSteps++
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
          shortestDistance = min(currentSteps + grid[pos], shortestDistance)
        }
      }
    case 'L':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x - 1, pos.y}
        currentSteps++
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
          shortestDistance = min(currentSteps + grid[pos], shortestDistance)
        }
      }
    case 'R':
      for i := 0; i < strToInt(section[1:]); i++ {
        pos = Vect{pos.x + 1, pos.y}
        currentSteps++
        if grid[pos] != 0 {
          nearestManhattan = min(getDistance(pos), nearestManhattan)
          shortestDistance = min(currentSteps + grid[pos], shortestDistance)
        }
      }
    }
  }
  return nearestManhattan, shortestDistance
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