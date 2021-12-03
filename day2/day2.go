package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	forward direction = "forward"
	up      direction = "up"
	down    direction = "down"
)

type movement struct {
	direction direction
	distance  int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	movements := make([]movement, 0)

	for scanner.Scan() {
		l := scanner.Text()

		split := strings.Split(l, " ")
		movementDirection, movementDistance := split[0], split[1]

		distanceInt, err := strconv.Atoi(movementDistance)
		if err != nil {
			log.Fatal(err)
		}

		movements = append(movements, movement{
			direction: direction(movementDirection),
			distance:  distanceInt,
		})
	}

	position := calculateSubmarinePosition(movements)
	log.Printf("part1 position: %d", position)
}

// calculateSubmarinePosition - Calculates the position of the submarine based on the movements
func calculateSubmarinePosition(movements []movement) int {
	x := 0
	y := 0

	for _, movement := range movements {
		switch movement.direction {
		case up:
			y -= movement.distance
		case down:
			y += movement.distance
		case forward:
			x += movement.distance
		}
	}

	return x * y
}
