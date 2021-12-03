package main

import (
	"log"

	"github.com/joakimhew/aoc2021/utils"
)

func main() {
	inputLines, err := utils.ReadInput("input")
	if err != nil {
		log.Fatal(err)
	}

	incrementsPart1, err := calcPart1(inputLines, 12)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("part1 (gamma rate * epsilon rate): %d", incrementsPart1)
}

// calculateGammaAndEpsilonRatePart1 calculates the gamma and epsilon rate based on binary data with the most common gamma bit and the least common epsilon bit
func calcPart1(input []string, width int) (int64, error) {
	var gamma int64
	var epsilon int64

	for i := 0; i < width; i++ {
		var bitCount int
		for _, line := range input {
			if line[i] == '1' {
				bitCount++
			} else {
				bitCount--
			}
		}

		gamma <<= 1
		epsilon <<= 1

		if bitCount > 0 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	return gamma * epsilon, nil
}
