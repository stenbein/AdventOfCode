// Advent of Code 2016 - Day 2
//
// 		The document goes on to explain that each button to be
// pressed can be found by starting on the previous button and
// moving to adjacent buttons on the keypad: U moves up, D
// moves down, L moves left, and R moves right. Each line of
// instructions corresponds to one button, starting at the
// previous button (or, for the first line, the "5" button);
// press whatever button you're on at the end of each line.
// If a move doesn't lead to a button, ignore it.
//
// 		Your puzzle input is the instructions from the document
// you found at the front desk. What is the bathroom code?
//
// 		Approach take was same as day 1, model the keypad as
// an actual object and imagine pressing the buttons as methods.
//
// Issues: The switch to part 2 intruduced a change from int to
// character value as a possible output. To do this I copied
// the entire struct and methods but changed the output to runes.
// Once I got the correct answer I realized I could have just
// made an array of ints representing the new characters and
// mapped it in the output.
//
// Modified the code to merge the two keypad structs into a
// a single type.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Keypad struct {
	row int
	col int

	pad [][]rune
}

func newKeypad() Keypad {

	k := Keypad{1, 1, make([][]rune, 3)}

	k.pad[0] = []rune{'1', '2', '3'}
	k.pad[1] = []rune{'4', '5', '6'}
	k.pad[2] = []rune{'7', '8', '9'}

	return k
}

func newAdvancedKeypad() Keypad {

	k := Keypad{1, 1, make([][]rune, 5)}

	k.pad[0] = []rune{'0', '0', '1', '0', '0'}
	k.pad[1] = []rune{'0', '2', '3', '4', '0'}
	k.pad[2] = []rune{'5', '6', '7', '8', '9'}
	k.pad[3] = []rune{'0', 'A', 'B', 'C', '0'}
	k.pad[4] = []rune{'0', '0', 'D', '0', '0'}

	return k
}

func (k *Keypad) getDigit() rune {

	return k.pad[k.row][k.col]

}

func (k *Keypad) translocate(command rune) {

	holdrow := k.row
	holdcol := k.col
	size := len(k.pad) - 1

	switch command {

	case 'U':
		if k.row > 0 {
			k.row -= 1
		}
	case 'D':
		if k.row < size {
			k.row += 1
		}
	case 'L':
		if k.col > 0 {
			k.col -= 1
		}
	case 'R':
		if k.col < size {
			k.col += 1
		}

	}

	//undo change if change is illegial character
	if k.getDigit() == '0' {
		k.row = holdrow
		k.col = holdcol
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	kp := newKeypad()
	output := ""

	instructions := strings.Split(input, "\n")
	for _, commands := range instructions {

		if commands != "" {
			for _, command := range commands {
				kp.translocate(command)
			}

			output += string(kp.getDigit())
		}

	}

	return output

}

func part_two(input string) string {

	kp := newAdvancedKeypad()
	output := ""

	instructions := strings.Split(input, "\n")
	for _, commands := range instructions {

		if commands != "" {
			for _, command := range commands {
				kp.translocate(command)
			}

			output += string(kp.getDigit())
		}

	}

	return output

}

func main() {

	input, err := ioutil.ReadFile("./2016_02.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
