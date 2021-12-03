package main

import (
	"log"
	"strconv"

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
	var gammaBits string
	var epsilonBits string

	for i := 0; i < width; i++ {
		var bitCount int
		for _, line := range input {
			if line[i] == '1' {
				bitCount++
			} else {
				bitCount--
			}
		}

		if bitCount > 0 {
			gammaBits += "1"
			epsilonBits += "0"
		} else {
			gammaBits += "0"
			epsilonBits += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaBits, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilon, err := strconv.ParseInt(epsilonBits, 2, 64)
	if err != nil {
		return 0, err
	}

	return gamma * epsilon, nil
}
