// Advent of Code 2016 - Day 3
//
//		The design document gives the side lengths of each 
// triangle it describes, but... 5 10 25? Some of these aren't
// triangles. You can't help but mark the impossible ones.

// 		In a valid triangle, the sum of any two sides must be 
// larger than the remaining side. For example, the "triangle"
//  given above is impossible, because 5 + 10 is not larger 
// than 25.

//		In your puzzle input, how many of the listed triangles 
// are possible?
//
// 		Approach taken
// 
// Issues: strconv.Atoi() returns both a value and an error, in
// this context the error is ignorable but it means that we can't
// inline the converstion into the triagle struct. Go isn't 
// "clever" like python. Best get it out of your head.
//
//
//
//
//
//
//
//
//

package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

type Triangle struct {

	sideA 	int
	sideB 	int
	sideC 	int

}

func (t *Triangle) isValid() bool {

	return (t.sideA + t.sideB > t.sideC) && (t.sideA + t.sideC > t.sideB) && (t.sideB + t.sideC > t.sideA)

}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func part_one(input string) string {

    count := 0

    triangles := strings.Split(input, "\n")
    for _, triangle := range triangles {

    	if triangle != "" {
	    	sides := strings.Fields(triangle)

	    	a, _ := strconv.Atoi(sides[0])
	    	b, _ := strconv.Atoi(sides[1])
	    	c, _ := strconv.Atoi(sides[2])

	    	t := Triangle{a,b,c}
	    	if t.isValid() {
	    		count++
	    	}
    	}

    }

    return strconv.Itoa(count)

}

func part_two(input string) string {

    count := 0

    rows := strings.Split(input, "\n")

    for i := 0; i < len(rows); i += 3{

    	if rows[i] != "" {

    		lengthsA := strings.Fields(rows[i])
    		lengthsB := strings.Fields(rows[i+1])
    		lengthsC := strings.Fields(rows[i+2])

    		for j := 0; j < 3; j++ {

    			a, _ := strconv.Atoi(lengthsA[j])
				b, _ := strconv.Atoi(lengthsB[j])
				c, _ := strconv.Atoi(lengthsC[j])

				t := Triangle{a,b,c}
				if t.isValid() {
					count++
				}

    		}

    	}

    }

    return strconv.Itoa(count)

}

func main() {

	input, err := ioutil.ReadFile("/home/mark/AdventOfCode/2016/Inputs/2016_03.txt")
    check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
