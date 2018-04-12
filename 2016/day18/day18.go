// Advent of Code 2016 - Day 18
//
// 		Approach taken was to treat the row of tiles as
// booleans and to pass the upper row as slices to determine
// the next row down. Since the pattern was so simple I didn't
// see a need to 
//
// 		Issues: None, just off by ones and init errors
//
//

package main

import (
	"fmt"
	"strconv"
)

type row []bool //true is trapped, false is not trapped

type rowBuilder func(row) row //template incase we need to swap patterns

type room struct {
	rb rowBuilder
	r  []row
}

func (r *room) reset() {
	r.r = []row{r.r[0]} //save input
}

func (r *room) next() {
	r.r = append(r.r, r.rb(r.r[len(r.r)-1]))
}

func (r *room) numTrapped() int {

	count := 0
	for _, row := range r.r {
		for _, b := range row {
			if b == false {
				count++
			}
		}
	}
	return count
}

//    Its left and center tiles are traps, but its right tile is not.
//    Its center and right tiles are traps, but its left tile is not.
//    Only its left tile is a trap.
//    Only its right tile is a trap.
func mapRow(r row) bool {

	var result bool
	switch r[0] {
	case true:
		result = !r[2]
	case false:
		result = r[2]
	}

	return result
}

func part_one(input string, count int) string {

	r := []row{row{}}
	for _, c := range input {
		if c == '.' {
			r[0] = append(r[0], false)
		} else {
			r[0] = append(r[0], true)
		}
	}

	rb := func(r row) row {

		nr := row{}
		//first row

		nr = append(nr, mapRow(row{false, r[0], r[1]}))
		for i := 0; i < len(r)-2; i++ {
			nr = append(nr, mapRow(row{r[i], r[i+1], r[i+2]}))

		}
		//last row
		nr = append(nr, mapRow(row{r[len(r)-2], r[len(r)-1], false}))
		return nr
	}

	room := room{rb, r}

	for i := 0; i < count-1; i++ { // -1 because we start with the first row init already
		room.next()
	}

	return strconv.Itoa(room.numTrapped())

}

func part_two(input string, count int) string {
	return part_one(input, count) //part two is just a change of params
}

func main() {

	input := ".^^^.^.^^^.^.......^^.^^^^.^^^^..^^^^^.^.^^^..^^.^.^^..^.^..^^...^.^^.^^^...^^.^.^^^..^^^^.....^...."

	fmt.Println("Problem 1: " + part_one(input, 40))
	fmt.Println("Problem 2: " + part_two(input, 400000))

}
