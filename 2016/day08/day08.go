// Advent of Code 2016 - Day 8
//
//		You come across a door implementing what you can only
// assume is an implementation of two-factor authentication
// after a long game of requirements telephone.
//
// 		To get past the door, you first swipe a keycard (no
// problem; there was one on a nearby desk). Then, it displays
// a code on a little screen, and you type that code on a
// keypad. Then, presumably, the door unlocks.
//
// 		Unfortunately, the screen has been smashed. After a
// few minutes, you've taken everything apart and figured
// out how it works. Now you just have to work out what the
// screen would have displayed.
//
// 		The magnetic strip on the card you swiped encodes
// a series of instructions for the screen; these
// instructions are your puzzle input. The screen is 50
// pixels wide and 6 pixels tall, all of which start off,
// and is capable of three somewhat peculiar operations:
//
// 		rect AxB turns on all of the pixels in a rectangle at
// the top-left of the screen which is A wide and B tall.
// 		rotate row y=A by B shifts all of the pixels in row A
// (0 is the top row) right by B pixels. Pixels that would
// fall off the right end appear at the left end of the row.
// 		rotate column x=A by B shifts all of the pixels in
// column A (0 is the left column) down by B pixels. Pixels
// that would fall off the bottom appear at the top of the
// column.
// 
// 		Part 1: There seems to be an intermediate check of 
// the voltage used by the display: after you swipe your card,
// if the screen did work, how many pixels should be lit?
//
//		Part 2: You notice that the screen is only capable 
// of displaying capital letters; in the font it uses, each 
// letter is 5 pixels wide and 6 tall.
//
//		After you swipe your card, what code is the screen 
// trying to display?
//
// 		Approach taken was to model the screen as an array of
// arrays and capture the rotations in functions.
//
// 		Issues: Worked fine, but I forgot to include the index
// in the modulo for the rotation functions 'i + '
//
//

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type screen [6][50]bool

func (s *screen) decode(inst string) {

	parts := strings.Split(inst, " ")
	switch parts[0] {
	case "rect":

		w, _ := strconv.Atoi(strings.Split(parts[1], "x")[0])
		l, _ := strconv.Atoi(strings.Split(parts[1], "x")[1])
		s.rect(w, l)

	case "rotate":

		y, _ := strconv.Atoi(parts[2][2:])
		b, _ := strconv.Atoi(parts[4])

		switch parts[1] {
		case "row":
			s.rotateRow(y, b)
		case "column":
			s.rotateCol(y, b)
		}
	}

}

func (s *screen) rect(w int, h int) {

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			s[j][i] = true
		}
	}

}

func (s *screen) rotateRow(r int, b int) {

	size := len(s[0]) //50 for now
	shift := b % size
	buff := make([]bool, size)

	for i := 0; i < size; i++ {
		buff[(i+shift)%size] = s[r][i]
	}
	for i := 0; i < size; i++ {
		s[r][i] = buff[i]
	}
}

func (s *screen) rotateCol(c int, b int) {

	size := len(s) //6 for now
	shift := b % size
	buff := make([]bool, size)

	for i := 0; i < size; i++ {
		buff[(i+shift)%size] = s[i][c]
	}
	for i := 0; i < size; i++ {
		s[i][c] = buff[i]
	}
}

func (s *screen) lit() int {

	count := 0
	for _, row := range s {

		for _, col := range row {
			if col {
				count++
			}
		}

	}

	return count

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	var s screen

	lines := strings.Split(input, "\n")
	for _, inst := range lines {

		if inst != "" {
			s.decode(inst)
		}

	}

	return strconv.Itoa(s.lit())

}

func part_two(input string) string {

	var s screen

	lines := strings.Split(input, "\n")
	for _, inst := range lines {

		//todo, check for alternative
		if inst != "" {
			s.decode(inst)
		}

	}

	for _, row := range s {
		for _, col := range row {

			if col == true {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("")
	}

	return ""

}

func main() {

	input, err := ioutil.ReadFile("./2016_08.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
