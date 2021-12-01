package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const windowSize = 3

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	measurements := make([]int, 0)

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		measurements = append(measurements, measurement)
	}

	incrementsPart1 := calcPart1(measurements)
	log.Printf("part1 total number of increments: %d", incrementsPart1)

	incrementsPart2 := calcPart2(measurements)
	log.Printf("part2 total number of increments: %d", incrementsPart2)
}

// calcPart1 calculates the number of increments from the input file.
func calcPart1(measurements []int) int {
	last := -1
	increments := 0
	for _, m := range measurements {
		if last != -1 && m > last {
			increments++
		}
		last = m
	}

	return increments
}

// calcPart2 calculates the number of increments from the input file in a sliding window of size 3.
func calcPart2(measurements []int) int {
	increments := 0
	window := make([]int, windowSize)
	lastSum := 0

	for idx, m := range measurements {
		sum := lastSum + m - window[0]
		window = append(window[1:], m)

		if idx < 3 {
			lastSum = sum
			continue
		}
		if lastSum < sum {
			increments++
		}
		lastSum = sum
	}

	return increments
}
