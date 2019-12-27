package main

import "testing"

func TestManhattanDistance(t *testing.T) {
  manhattanTests := []struct {
    wires []string
    manhattanDistance int
    shortestDistance int
  }{
    {[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159, 610},
    {[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135, 410},
  }

  for _, test := range manhattanTests {
    manhattanDistance, shortestDistance := manhattanDistance(test.wires[0], test.wires[1])
    if manhattanDistance != test.manhattanDistance {
      t.Errorf("got %v want %v", manhattanDistance, test.manhattanDistance)
    }
    if shortestDistance != test.shortestDistance {
      t.Errorf("got %v want %v", shortestDistance, test.shortestDistance)
    }
  }
}