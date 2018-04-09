// Advent of Code 2016 - Day
//
// 		Every location in this area is addressed by a pair
// of non-negative integers (x,y). Each such coordinate is
// either a wall or an open space. You can't move
// diagonally. The cube maze starts at 0,0 and seems to
// extend infinitely toward positive x and y; negative
// values are invalid, as they represent a location outside
// the building. You are in a small waiting area at 1,1.
//
// 		While it seems chaotic, a nearby morale-boosting
// poster explains, the layout is actually quite logical.
// You can determine whether a given x,y coordinate will be
// a wall or an open space using a simple system:
//
// 		Find x*x + 3*x + 2*x*y + y + y*y.
//		Add the office designer's favorite number (your
// puzzle input).
//		Find the binary representation of that sum; count
// the number of bits that are 1.
// 		If the number of bits that are 1 is even, it's
// an open space.
// If the number of bits that are 1 is odd, it's a wall.
//
//		Part 1: What is the fewest number of steps
// required for you to reach 31,39?
//
//		Part 2: How many locations (distinct x,y
// coordinates, including your starting location) can you
// reach in at most 50 steps?
//
// 		Approach taken was to build out the grid recursivly
// with a modified breath first search. This worked well
// but I ended up having several problems with bugs in the
// code.
//
// 		Issues: Several mistakes with the logic, off by 1s
// mixup of != and ==, and some range check issues.
// Otherwise ok. To correct this I eventually decided to
// make tests. I should have done this from the beginning.
//

package main

import (
	"fmt"
	"strconv"
)

type tile struct {
	x int
	y int
}

type grid struct {
	start  tile
	secret int

	tiles [][]int
}

func (g *grid) paint(tiles *[]tile, dist int) {

	for _, t := range *tiles {
		g.tiles[t.y][t.x] = dist
	}
}

func (g *grid) neighbours(tiles *[]tile) *[]tile {

	var out []tile
	out = make([]tile, 0)

	for _, t := range *tiles {

		for _, nt := range [4]tile{tile{t.x, t.y - 1}, tile{t.x, t.y + 1}, tile{t.x - 1, t.y}, tile{t.x + 1, t.y}} {
			if g.isOpen(nt) && g.tiles[nt.y][nt.x] == 0 {
				out = append(out, nt)
			}
		}

	}

	return &out

}

func (g *grid) bfs(target tile, max int) {

	distance := 0
	unvisited := &[]tile{g.start}
	for {
		//paint everything in waves
		g.paint(unvisited, distance)

		//check if we reached the target,
		//taking advantage of short circuts here
		if g.inBounds(target) && g.tiles[target.y][target.x] != 0 {
			break
		}

		//find all the neighbours of our current targets
		unvisited = g.neighbours(unvisited)

		distance++
		if distance > max {
			break
		}
	}

	//I noticed when debugging that while I was getting the
	//correct answer, the start point was mislabled in distance
	//this is because the abstraction is leaking. I'm assuming
	//unvisited tiles are zeros, but my distance also starts at
	//zero, so the start gets rewritten with a distance of 2
	//for now, just going to rewrite it to zero, but I think it
	//will be a good exercise to fix this later
	g.paint(&[]tile{g.start}, 0)

}

func (g *grid) inBounds(target tile) bool {

	if target.y < len(g.tiles) && target.x < len(g.tiles[0]) {
		return true
	} else {
		return false
	}

}

func (g *grid) distance(target tile) int {

	g.bfs(target, 999) //max set to impossible number
	return g.tiles[target.y][target.x]

}

func (g *grid) grow() {

	var newCol []int
	newCol = make([]int, len(g.tiles))
	g.tiles = append(g.tiles, newCol)

	for i, _ := range g.tiles {
		g.tiles[i] = append(g.tiles[i], 0)
	}

}

//Find x*x + 3*x + 2*x*y + y + y*y.
//Add the office designer's favorite number (your puzzle input).
//Find the binary representation of that sum; count the number of bits that are 1.
// If the number of bits that are 1 is even, it's an open space.
// If the number of bits that are 1 is odd, it's a wall.
func (g *grid) isOpen(t tile) bool {

	//simplification for our search function
	if t.x < 0 || t.y < 0 {
		return false
	}
	if !g.inBounds(t) {
		g.grow()
	}

	digitsEven := true
	val := int64(t.x*t.x + 3*t.x + 2*t.x*t.y + t.y + t.y*t.y + g.secret)

	for _, digit := range strconv.FormatInt(val, 2) {
		if digit == '1' {
			digitsEven = !(digitsEven && true) && (digitsEven || true) //xOr
		}
	}

	return digitsEven

}

//I have a feeling I'll need this again
func newGrid(start tile, secret int) *grid {

	var initial [][]int
	initial = make([][]int, start.y+1)
	for i, _ := range initial {
		initial[i] = make([]int, start.x+1)
	}

	return &grid{start, secret, initial}

}

func part_one(input int) string {

	var g *grid

	g = newGrid(tile{1, 1}, input)

	return strconv.Itoa(g.distance(tile{31, 39}))

}

func part_two(input int) string {

	var g *grid

	//offset for the initial square being zero as well
	count := 1

	g = newGrid(tile{1, 1}, input)
	g.bfs(tile{52, 52}, 50) // pick a target larger than the search space

	// generate the output
	for _, row := range g.tiles {

		for _, t := range row {
			if t > 0 && t <= 50 {
				count++
			}
		}

	}

	return strconv.Itoa(count)
}

func main() {

	input := 1352 //input is again a single value instead of a set

	fmt.Println("Problem 1: " + part_one(input))
	fmt.Println("Problem 2: " + part_two(input))

}
