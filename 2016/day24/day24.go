// Advent of Code 2016 - Day 24
//
// 		Approach taken was to generate a set of min
// distances and a set of path permutations and itterate
// through the path set. To get the shortest path distances
// I created a second data structure to paint across the 
// maze with the value of the distances at each step. 
// Once the painting was complete you can move to step two
// and keep track of which path set results in the best 
// overall distance score.
//
// 		Issues: Major issue with memory leak from appending 
// the same slice was not corrected and I eventually made a 
// workaround. Everything else worked well.
//
// TODO investigate this. It's some kind of behavior where
// the prior slice is not overwritten by the new slice.
//

package main

import (

    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
    "os"

)

var depth int

type square struct {

	passable 	bool
	distance 	int
	ID 			int

}

//represents each square in the maze
//used to find the shortest paths
type grid struct {

	squares		[][]square
	
}

//matrix represeting distance between paths
//paths[x][x] == 0, paths[x][y] == paths[y][x]
type paths struct {
	
	distances [][]int
	IDs []int

}

//simple method to print out the grid
func (g *grid) rep() {

	for _, row := range g.squares {
		out := ""
		for i, _ := range row {
			if !row[i].passable {
				out += "#"
			} else if row[i].ID != -1 {
				out += "+"
			} else if row[i].distance == -1 {
				out += "."
			} else {
				out += "7" //strconv.Itoa(row[i].distance)
			}

		}
		fmt.Println(out)
	}

}

func (g *grid) findID(id int) (x, y int) {

	for y, _ := range g.squares {
		for x, _ := range g.squares[y] {
			if g.squares[y][x].ID == id {
				return x, y
			}
		}
	}
	panic("ID not found in grid search...")

}

func (g *grid) reset() {

	for i, _ := range g.squares {
		for j, _ := range g.squares[i] {
			g.squares[i][j].distance = -1
		}
	}

}

//of the values just painted, check the surroundings and 
//return the coords of the items which can next be painted
//note this approach is slow, but we only need to call it a few times
func (g *grid) getNext(tarDist int) ([]int, []int) {

	xout := []int{}
	yout := []int{}

	min := 0
	ymax := len(g.squares) - 1
	xmax := len(g.squares[min])	- 1

	//for i, _ := range xs {
	for i, _ := range g.squares {

		for j, _ := range g.squares[i] {

			if g.squares[i][j].distance == tarDist {

				//one above
				if i > min {
					if g.squares[i - 1][j].passable && g.squares[i - 1][j].distance == -1 {
						yout = append(yout, i - 1)
						xout = append(xout, j)
					}
				}
				
				//one below
				if i < ymax {
					if g.squares[i + 1][j].passable && g.squares[i + 1][j].distance == -1 {
						yout = append(yout, i + 1)
						xout = append(xout, j)
					}
				}
				
				//one left
				if j > min {
					if g.squares[i][j - 1].passable && g.squares[i][j - 1].distance == -1 {
						yout = append(yout, i)
						xout = append(xout, j - 1)
					}
				}
				
				//one right
				if j < xmax {
					if g.squares[i][j + 1].passable && g.squares[i][j + 1].distance == -1 {
						yout = append(yout, i)
						xout = append(xout, j + 1)
					}
				}

			}
		}
	
	}

	return xout, yout

}

//paint the passed coords
func (g *grid) paint(xs, ys []int, dis int) {

	for i, _ := range xs {
		
		if g.squares[ys[i]][xs[i]].distance == -1 {
			g.squares[ys[i]][xs[i]].distance = dis		
		}
	
	}

}


func (g *grid) flood(x, y int) {

	dist := 0
	//count := 0
	xs := []int{x}
	ys := []int{y}
	
	for len(xs) > 0 {
		
		g.paint(xs, ys, dist)
		xs, ys = g.getNext(dist)
		dist++
/*
		if count > depth {
			g.rep()
		}
		if count > depth + 2 {panic("stop")}
		count++*/

	
	}
	
}


func (g *grid) dist(id int) []int {

	out := make([]int, 8) //todo 

	g.reset()
	
	//find location of id and flood from there
	x, y := g.findID(id)
	g.flood(x, y)
	
	for y, _ := range g.squares {
		for x, _ := range g.squares[y] {
			if g.squares[y][x].ID != -1 {
				out[g.squares[y][x].ID] = g.squares[y][x].distance
			}
		}
	}

	return out

}

//convert the wrar string maze into the maze grid structs
func makeGrid(raw string) grid {

	var out grid
	rows := strings.Split(raw, "\n")
	
	out.squares = make([][]square, len(rows))

	for i, r := range rows {
		
		out.squares[i] = make([]square, len(rows[0]))
		for j, c := range r {
			
			switch c {
			case '#':
				out.squares[i][j] = square{false, -1, -1}
			case '.':
				out.squares[i][j] = square{true, -1, -1}
			default:
				s, _ := strconv.Atoi(string(c))
				out.squares[i][j] = square{true, -1, s}
			}
		}
	}

	return out

}

//get the collection of shortest distances between each point in the maze
func makePaths(raw string) paths {

	var out [][]int
	
	IDs := []int{0,1,2,3,4,5,6,7}
	g := makeGrid(raw)
	
	out = make([][]int, len(IDs))

	for _, id := range IDs {
		dis := g.dist(id)
		out[id] = dis
	}
	
	return paths{out, IDs}
	
}

//borrowed a permutation function again
//the below is really neat. Builds a function 
//which closes over the result list
func permutations(arr []int) [][]int {
    
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }

    helper(arr, len(arr))

    return res

}


//input is a representation of a maze
//points of interest are integers
func getIDs(s string) []int {

    out := []int{}
    for _, r := range s {
        if r >= '0' && r <= '9' {
            out = append(out, int(r - '0'))
        }
    }
    return out
}

func distance_one(p *paths, path []int) int {

    i := 0
    dist := p.distances[0][path[0]] //zero to the first part of the path
    for i < len(path)-1 {
        dist += p.distances[path[i]][path[i+1]]
        i++
    }

    return dist

}

func distance_two(p *paths, path []int) int {

    i := 0
    dist := p.distances[0][path[0]] //zero to the first part of the path
    for i < len(path)-1 {
        dist += p.distances[path[i]][path[i+1]]
        i++
    }
    dist += p.distances[path[i]][0] //return to zero from end of path

    return dist

}

//start with 0, check all permutations of legal paths
//save the shortest and return the distance and path
func (p *paths) solve(fdistance func(*paths, []int) int) (int, []int) {

    var path, shortest    []int //path is a list of ids to visit in an order
    var dist, mindist     int

    //init mindist
    mindist = 9999999

    //visit each set of paths in order
    //for _, path = range permutations(p.IDs) {
    for _, path = range permutations([]int{1,2,3,4,5,6,7}) {

        dist = fdistance(p, path)
		
        if dist < mindist {
            shortest = path
            mindist = dist
        }

    }

    return mindist, shortest

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func part_one(input string) string {

	p := makePaths(input) //the shortest paths for the maze from each point of interest

	dist, path := p.solve(distance_one)
	
	fmt.Println(path) //what could it be!?
	
	return strconv.Itoa(dist)
	
}

func part_two(input string) string {

	p := makePaths(input) //the shortest paths for the maze from each point of interest

	dist, path := p.solve(distance_two)
	
	fmt.Println(path) //what could it be!?
	
	return strconv.Itoa(dist)

}

func main() {

	sdepth := os.Args[1:][0]
	depth, _ = strconv.Atoi(sdepth)

	input, err := ioutil.ReadFile("./2016_24.txt")
	
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
