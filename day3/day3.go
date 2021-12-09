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

	incrementsPart1, err := part1(inputLines, 12)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("part1 (gamma rate * epsilon rate): %d", incrementsPart1)

	part2, err := part2(inputLines, 12)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("part2 (life support rating): %d", part2)
}

// calculateGammaAndEpsilonRatePart1 calculates the gamma and epsilon rate based on binary data with the most common gamma bit and the least common epsilon bit
func part1(input []string, width int) (int64, error) {
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

// part2 calculates the life support rating by multiplying the oxygen generator rating by the co2 scrubber rating.
func part2(input []string, width int) (int, error) {
	return search(input, oxygenFilter) * search(input, co2filter), nil
}

func search(input []string, filter func(rune, rune) rune) int {
	remaining := make([]string, len(input))
	copy(remaining, input)

	offset := 0

	for len(remaining) > 1 {
		left := make([]string, 0, len(remaining))
		target := filter(mostLeast(remaining, offset))

		for _, line := range remaining {
			if rune(line[offset]) == target {
				left = append(left, line)
			}
		}

		remaining = left
		offset++
	}

	v, _ := strconv.ParseInt(remaining[0], 2, 32)
	return int(v)
}

func oxygenFilter(most, _ rune) rune {
	return most
}

func co2filter(_, least rune) rune {
	return least
}

func mostLeast(input []string, idx int) (rune, rune) {
	zeroes := 0
	ones := 0

	for _, line := range input {
		if line[idx] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes > ones {
		return '0', '1'
	}

	return '1', '0'
}
