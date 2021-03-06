package main

import "testing"

var testData = []string{
	"COM)B",
	"B)C",
	"C)D",
	"D)E",
	"E)F",
	"B)G",
	"G)H",
	"D)I",
	"E)J",
	"J)K",
	"K)L",
}

var orbitalTransferTestData = []string{
	"COM)B",
	"B)C",
	"C)D",
	"D)E",
	"E)F",
	"B)G",
	"G)H",
	"D)I",
	"E)J",
	"J)K",
	"K)L",
	"K)YOU",
	"I)SAN",
}

func TestLoadData(t *testing.T) {
	orbits := loadData(testData)

	if len(orbits) != 11 {
		t.Errorf("got %v want %v", len(orbits), 11)
	}
}

func TestCalculateOrbits(t *testing.T) {
	orbits := loadData(testData)

	if countOrbits(orbits) != 42 {
		t.Errorf("got %v want %v", countOrbits(orbits), 42)
	}
}

func TestOrbitalTransfer(t *testing.T) {
	orbits := loadData(orbitalTransferTestData)

	if orbitalTransfer(orbits) != 4 {
		t.Errorf("Got %v want %v", orbitalTransfer(orbits), 4)
	}
}
