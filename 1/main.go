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

  for scanner.Scan() {
    i, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatal(err)
    }
    fuelRequired := fuelRequired(i)
    fmt.Println(i, fuelRequired)
    totalFuel += fuelRequired
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Total Fuel: %v", totalFuel)
}

func fuelRequired(mass int) int {
  return (mass/3)-2
}