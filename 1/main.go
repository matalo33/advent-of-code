package main

import (
  "bufio"
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
  defer file.Close()

  scanner := bufio.NewScanner(file)
  totalFuel := 0
  totalFuelWithFuel := 0

  for scanner.Scan() {
    i, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatal(err)
    }
    totalFuel += fuelRequired(i)
    totalFuelWithFuel += fuelRequiredWithFuel(i)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Total Fuel: %v\n", totalFuel)
  fmt.Printf("Total Fuel with added fuel: %v\n", totalFuelWithFuel)
}

func fuelRequired(mass int) int {
  result := (mass/3)-2
  if result < 0 {
    return 0
  } else {
    return result
  }
}

func fuelRequiredWithFuel(mass int) int {
  totalFuel := 0
  currentMass := mass
  for {
    newFuel := fuelRequired(currentMass)
    if newFuel == 0 {
      break
    }
    totalFuel += newFuel
    currentMass = newFuel
  }
  return totalFuel
}