package main

import (
	"fmt"
    "strconv"
    "io/ioutil"
	"testing"
)

//check that the init grid includes the start point
func ensureFind(t *testing.T, g *grid, id, expectx, expecty int) {

	if x, y := g.findID(id); x != expectx || y != expecty {
		t.Errorf("g.findID(%d) = %d, %d; want %d, %d", id, x, y, expectx, expecty)
	}

}


func TestFindID(t *testing.T) {

	input, _ := ioutil.ReadFile("./2016_24_t.txt")

	var g grid
	g = makeGrid(string(input))

	ensureFind(t, &g, 0, 1, 1)
	ensureFind(t, &g, 1, 3, 1)
	ensureFind(t, &g, 2, 9, 1)
	ensureFind(t, &g, 3, 9, 3)
	ensureFind(t, &g, 4, 1, 3)

	input, _ = ioutil.ReadFile("./2016_24.txt")

	g = makeGrid(string(input))

	ensureFind(t, &g, 0, 149, 11)
	ensureFind(t, &g, 1, 7, 5)
	ensureFind(t, &g, 2, 149, 3)
	ensureFind(t, &g, 6, 27, 27)
	ensureFind(t, &g, 7, 171, 17)

}
