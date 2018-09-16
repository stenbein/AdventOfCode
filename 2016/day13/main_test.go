package main

import (
	"fmt"
	"testing"
)

//I realized afting writng these tests that I mixed
//Cartesian coords and tile structs. Todo, swap this out.

//simple method to print out the grid
func (g *grid) rep() {
	for _, row := range g.tiles {
		fmt.Println(row)
	}
}

//check that the init grid includes the start point
func ensureSize(t *testing.T, g *grid, expectx, expecty int) {

	if l := len(g.tiles); l != expecty {
		t.Errorf("len(g.tiles) = %d, want %d", l, expecty)
	}

	if l := len(g.tiles[0]); l != expectx {
		t.Errorf("len(g.tiles) = %d, want %d", l, expectx)
	}

}

//check the value at a particular grid point
func ensureValue(t *testing.T, g *grid, x, y, expectValue int) {

	if v := g.tiles[y][x]; v != expectValue {
		t.Errorf("g.tiles[y][x] = %d, want %d", v, expectValue)
	}

}

//check the a particular grid point is open
func ensureOpen(t *testing.T, g *grid, x, y int, expectValue bool) {

	if b := g.isOpen(tile{x, y}); b != expectValue {
		t.Errorf("g.isOpen(tile{%d,%d}) = %d, want %d", x, y, b, expectValue)
	}

}

func TestNew(t *testing.T) {

	var g *grid
	g = newGrid(tile{0, 0}, 0)
	ensureSize(t, g, 1, 1)

	g = newGrid(tile{1, 0}, 23)
	ensureSize(t, g, 2, 1)

	g = newGrid(tile{0, 1}, 0)
	ensureSize(t, g, 1, 2)

	g = newGrid(tile{1, 1}, 0)
	ensureSize(t, g, 2, 2)

	g = newGrid(tile{2, 2}, 0)
	ensureSize(t, g, 3, 3)

}

func TestGrow(t *testing.T) {

	var g *grid
	g = newGrid(tile{0, 0}, 10)
	g.grow()
	ensureSize(t, g, 2, 2)
	g.grow()
	ensureSize(t, g, 3, 3)
	g.grow()
	ensureSize(t, g, 4, 4)

	g = newGrid(tile{1, 0}, 10)
	ensureSize(t, g, 2, 1)
	g.grow()
	ensureSize(t, g, 3, 2)
	g.grow()
	ensureSize(t, g, 4, 3)

	g = newGrid(tile{0, 1}, 10)
	ensureSize(t, g, 1, 2)
	g.grow()
	ensureSize(t, g, 2, 3)
	g.grow()
	ensureSize(t, g, 3, 4)

}

func TestPaint(t *testing.T) {

	var g *grid
	g = newGrid(tile{2, 2}, 10)
	g.paint(&[]tile{tile{1, 1}, tile{0, 0}}, 1)
	g.paint(&[]tile{tile{0, 1}, tile{1, 0}}, 2)

	ensureValue(t, g, 0, 0, 1)
	ensureValue(t, g, 1, 1, 1)

	ensureValue(t, g, 0, 1, 2)
	ensureValue(t, g, 1, 0, 2)

}

/*input from day 13 sample for testing
  0123456789
0 .#.####.##
1 ..#..#...#
2 #....##...
3 ###.#.###.
4 .##..#..#.
5 ..##....#.
6 #...##.###
*/
func TestIsOpen(t *testing.T) {

	var g *grid
	g = newGrid(tile{2, 2}, 10)
	ensureOpen(t, g, 0, 0, true)
	ensureOpen(t, g, 0, 1, true)
	ensureOpen(t, g, 0, 2, false)

	ensureOpen(t, g, 1, 0, false)
	ensureOpen(t, g, 1, 1, true)
	ensureOpen(t, g, 1, 2, true)

	ensureOpen(t, g, 2, 0, true)
	ensureOpen(t, g, 2, 1, false)
	ensureOpen(t, g, 2, 2, true)

}

func TestNeighbours(t *testing.T) {

	var g *grid
	g = newGrid(tile{2, 2}, 10)

	n := g.neighbours(&[]tile{tile{0, 0}})
	if len(*n) != 1 {
		fmt.Println(n)
		fmt.Println("Wrong number of neighbours.")
		t.Fail()
	}
	if (*n)[0].x != 0 && (*n)[0].y != 1 {
		fmt.Println(n)
		fmt.Println("Wrong neighbour returned.")
		t.Fail()
	}

	n = g.neighbours(&[]tile{tile{1, 1}})
	if len(*n) != 2 {
		fmt.Println(n)
		fmt.Println("Wrong number of neighbours.")
		t.Fail()
	}
	if !((*n)[0].x == 0 && (*n)[0].y == 1) && !((*n)[1].x == 0 && (*n)[1].y == 1) {
		fmt.Println(n)
		fmt.Println("Wrong neighbour returned, 1.")
		t.Fail()
	}
	if !((*n)[0].x == 1 && (*n)[0].y == 2) && !((*n)[1].x == 1 && (*n)[1].y == 2) {
		fmt.Println(n)
		fmt.Println("Wrong neighbour returned, 2.")
		t.Fail()
	}
}
