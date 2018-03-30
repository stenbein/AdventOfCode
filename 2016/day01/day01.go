// Advent of Code 2016 - Day 1
//
// 		The Document indicates that you should start at the given 
// coordinates (where you just landed) and face North. Then, 
// follow the provided sequence: either turn left (L) or right 
// (R) 90 degrees, then walk forward the given number of 
// blocks, ending at a new intersection.
//
// 		How many blocks away is Easter Bunny HQ?
//
// 		Approch taken was an object oriented approach to model
// a struct which tracts the position of the walker and heading
// and to increment with each instruction from the input.
//
// Issues: During part two I missintrepreted the instructions
// to only check for locations after each instruction was 
// complete rather than after each step. This was fixed by 
// wrapping the walker in a loop.
//


package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Heading int

const (
	NORTH Heading = iota
	EAST
	SOUTH
	WEST
)

type Location struct {
	x int
	y int
}

type Walker struct {
	heading Heading
	delta_x int
	delta_y int
}

func (w *Walker) rotate(direction rune) {

	switch direction {

	case 'R':

		if w.heading == WEST {
			w.heading = NORTH
		} else {
			w.heading += 1
		}

	case 'L':

		if w.heading == NORTH {
			w.heading = WEST
		} else {
			w.heading -= 1
		}

	}

}

func (w *Walker) step(distance int) {

	switch w.heading {

	case NORTH:
		w.delta_y += distance
	case EAST:
		w.delta_x += distance
	case SOUTH:
		w.delta_y -= distance
	case WEST:
		w.delta_x -= distance
	}

}

func (w *Walker) distance() int {

	return int(math.Abs(float64(w.delta_y)) + math.Abs(float64(w.delta_x)))

}

func part_one(input string) int {

	commands := strings.Split(input, ", ")

	w := Walker{NORTH, 0, 0}

	for _, command := range commands {

		rotation := rune(command[0])
		distance, _ := strconv.Atoi(command[1:])

		w.rotate(rotation)
		w.step(distance)

	}
	return w.distance()

}

func part_two(input string) int {

	commands := strings.Split(input, ", ")

	w := Walker{NORTH, 0, 0}
	locations := make(map[Location]bool)

	for _, command := range commands {

		rotation := rune(command[0])
		distance, _ := strconv.Atoi(command[1:])

		w.rotate(rotation)

		for i := 0; i < distance; i++ {

			location := Location{w.delta_x, w.delta_y}
			_, visited := locations[location]

			if visited {
				break
			} else {
				locations[location] = true
			}

			w.step(1) //step comes after the location check so we get our initial location included in the map

		}

	}
	return w.distance()

}

func main() {

	const input = "R2, L3, R2, R4, L2, L1, R2, R4, R1, L4, L5, R5, R5, R2, R2, R1, L2, L3, L2, L1, R3, L5, R187, R1, R4, L1, R5, L3, L4, R50, L4, R2, R70, L3, L2, R4, R3, R194, L3, L4, L4, L3, L4, R4, R5, L1, L5, L4, R1, L2, R4, L5, L3, R4, L5, L5, R5, R3, R5, L2, L4, R4, L1, R3, R1, L1, L2, R2, R2, L3, R3, R2, R5, R2, R5, L3, R2, L5, R1, R2, R2, L4, L5, L1, L4, R4, R3, R1, R2, L1, L2, R4, R5, L2, R3, L4, L5, L5, L4, R4, L2, R1, R1, L2, L3, L2, R2, L4, R3, R2, L1, L3, L2, L4, L4, R2, L3, L3, R2, L4, L3, R4, R3, L2, L1, L4, R4, R2, L4, L4, L5, L1, R2, L5, L2, L3, R2, L2"

	fmt.Println("Problem 1: " + strconv.Itoa(part_one(input)))
	fmt.Println("Problem 2: " + strconv.Itoa(part_two(input)))

}
