package main

import "testing"

func TestFuelRequired(t *testing.T) {
  fuelTests := []struct {
    mass int
    fuel int
  }{
    {12, 2},
    {14, 2},
    {1969, 654},
    {100756, 33583},
  }

  for _, ship := range fuelTests {
    got := fuelRequired(ship.mass)
    if got != ship.fuel {
      t.Errorf("got %v want %v", got, ship.fuel)
    }
  }
}

func TestFuelRequiredWithFuel(t *testing.T) {
  fuelTests := []struct {
    mass int
    fuel int
  }{
    {14, 2},
    {1969, 966},
    {100756, 50346},
  }

  for _, ship := range fuelTests {
    got := fuelRequiredWithFuel(ship.mass)
    if got != ship.fuel {
      t.Errorf("got %v want %v", got, ship.fuel)
    }
  }
}