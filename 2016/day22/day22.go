// Advent of Code 2016 - Day 22
//
// 		Approach taken was
//
// 		Issues:
//
//


package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

type grid [][]node

type node struct {

	id		string

	x		int
	y		int

	size	int
	used 	int
	avail 	int

}

var grids grid

//Your goal is to gain access to the data which begins in the node 
//with y=0 and the highest x (that is, the node in the top-right corner)
func (g *grid) rep() {

	for _, row := range *g {
		for _, n := range row {
			fmt.Print(n.rep())
		}
		fmt.Println()
	}

}

func (n *node) rep() string {

	if n.y == 0 && n.x == 36 {
		return "G"
	} else if n.used == 0 {
		return "_"
	} else if n.size > 300 {
		return "#"
	} else {
		return "."
	}

}
/*
type byAvail []node

func (n byAvail) Len() int {
	return len(n)
}
func (n byAvail) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n byAvail) Less(i, j int) bool {
	return n[i].avail < n[j].avail
}*/

//Node A is not empty (its Used is not zero).
//Nodes A and B are not the same node.
//The data on node A (its Used) would fit on node B (its Avail).
func isPair(a node, b node) bool {

	return a.used != 0 && a.id != b.id && a.used <= b.avail

}

//Example format
//Filesystem              Size  Used  Avail  Use%
///dev/grid/node-x0-y0     91T   66T    25T   72%
func newNode(def string) node {

	tokens := strings.Fields(def) //fields not split since we have padding

	id := tokens[0]
	size, _ := strconv.Atoi(tokens[1][:len(tokens[1])-1])
	used, _ := strconv.Atoi(tokens[2][:len(tokens[2])-1])
	avail, _ := strconv.Atoi(tokens[3][:len(tokens[3])-1])

	loc := tokens[0][16:] //len("/dev/grid/node-x")
	tokens = strings.Split(loc, "-")
	
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1][1:]) //trim "y"
	return node{id, x, y, size, used, avail}

}

func process(input string) []node {

	defs := strings.Split(input, "\n")[2:] //first two lines are meta data
	nodes := make([]node, len(defs))

	for i, def := range defs {
		if def != "" {
			nodes[i] = newNode(def)
		}
	}

	return nodes

}

func load(nodes []node) {

	var max_x, max_y int

	//find the grid boundaries
	for _, n := range nodes {
		if n.x > max_x {
			max_x = n.x
		}
		if n.y > max_y {
			max_y = n.y
		}

	}
	fmt.Println(max_x)
	fmt.Println(max_y)

	grids = make([][]node, max_y + 1) //+1 values are inclusive
	for i := range grids {
		grids[i] = make([]node, max_x + 1)
	}

	for _, n := range nodes {
		grids[n.y][n.x] = n

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	nodes := process(input)

	count := 0
	for i, _ := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			if isPair(nodes[i], nodes[j]) {count++}
			if isPair(nodes[j], nodes[i]) {count++}
		}
	}
	return strconv.Itoa(count)

}

//Prints the grid, which is then solved by hand counting.
//Todo: replace with automated process.
func part_two(input string) string {

	nodes := process(input)
	load(nodes)
	grids.rep()
	
	return "complete"
}

func main() {

	input, err := ioutil.ReadFile("./2016_22.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
