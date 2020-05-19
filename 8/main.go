package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')

	imageData := convertToImageLayers(line, 25, 6)
	decodedImage := decodeImageLayers(imageData, 25, 6)

	fmt.Printf("Part 1: %v\n\n", checkTransmission(imageData))
	fmt.Printf("Part 2: \n")
	for _, val := range decodedImage {
		for _, pixel := range val {
			if pixel == 0 {
				fmt.Printf("⬛")
			} else if pixel == 1 {
				fmt.Printf("⬜")
			}
		}
		fmt.Printf("\n")
	}
}

func convertToImageLayers(input string, width int, height int) [][][]int {
	numLayers := len(input) / (width * height)
	layers := make([][][]int, numLayers)
	for layer := 0; layer < numLayers; layer++ {
		layers[layer] = make([][]int, height)
		for row := 0; row < height; row++ {
			layers[layer][row] = make([]int, width)
			substrStart := (layer * height * width) + (row * width)
			substrEnd := ((layer * height * width) + ((row + 1) * width))
			for i, c := range input[substrStart:substrEnd] {
				layers[layer][row][i] = int(c - '0')
			}
		}
	}
	return layers
}

// Part 1
func checkTransmission(imageData [][][]int) int {
	data := make(map[int]map[int]int)
	fewestZeroLayer, fewestZeroCount := 0, math.MaxInt32
	for layer := 0; layer < len(imageData); layer++ {
		data[layer] = make(map[int]int)
		for row := 0; row < len(imageData[layer]); row++ {
			for _, val := range imageData[layer][row] {
				data[layer][val] = data[layer][val] + 1
			}
		}
		if data[layer][0] < fewestZeroCount {
			fewestZeroLayer = layer
			fewestZeroCount = data[layer][0]
		}
	}
	return data[fewestZeroLayer][1] * data[fewestZeroLayer][2]
}

// Part 2
func decodeImageLayers(imageData [][][]int, width int, height int) [][]int {
	decodedImage := make([][]int, height)
	for x := 0; x < height; x++ {
		decodedImage[x] = make([]int, width)
		for y := 0; y < width; y++ {
			// Every cell starts as transparent
			decodedImage[x][y] = 2
		}
	}
	for layer := 0; layer < len(imageData); layer++ {
		for x := 0; x < height; x++ {
			for y := 0; y < width; y++ {
				if decodedImage[x][y] == 2 {
					decodedImage[x][y] = imageData[layer][x][y]
				}
			}
		}
	}
	return decodedImage
}
