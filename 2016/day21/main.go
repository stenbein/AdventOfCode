// Advent of Code 2016 - Day 21
//
// 		Approach taken was to perform each transformation
// as it's own function. And then process the input iteratively.
// This worked ok, but it took a bit of testing to narrow down
// what my issue was. Part two was mean, having build all that
// my first instinct was not to build inverses of all the functions.
// Instead I took the permutation of the string input and ran it
// against the decoder until I hit a match. Ran in about 2 seconds,
// which is obviously not great but fine for these purposes.
//
// 		Once this is all said and done this is a TODO, take a look
// at how the others solved this day. I have a feeling the
// instructions are collapsable but I'm not sure of the math.
//
// 		Issues: Off by 1 error in the rotate base on position code.
// I also though I had an issue with my code for the move function.
// But it turned out to be user error defining the tests.
//

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	PLAIN_TEXT  string = "abcdefgh"
	TARGET_TEXT string = "fbgdceah"
)

//swap position X with position Y means that the letters at indexes X and Y (counting from 0) should be swapped.
func swapI(s string, x, y int) string {

	r := []rune(s)
	r[x], r[y] = r[y], r[x]

	return string(r)

}

//swap letter X with letter Y means that the letters X and Y should be swapped (regardless of where they appear in the string).
func swapR(s string, x, y rune) string {

	xi := strings.IndexRune(s, x)
	yi := strings.IndexRune(s, y)

	//if the above searches return nothing, value will be -1 and will cause panic
	return swapI(s, xi, yi)

}

//rotate left/right X steps means that the whole string should be rotated; for example, one right rotation would turn abcd into dabc.
func rotate(s string, steps int, left bool) string {

	l := len(s)
	out := make([]rune, l)

	steps %= l
	if left {
		steps = l - steps
	}

	for i, r := range s {
		out[(i+steps)%l] = r
	}

	return string(out)

}

/*rotate based on position of letter X means that the whole string should be rotated to the right based on the index of letter X (counting from 0) as determined before this instruction does any rotations. Once the index is determined, rotate the string to the right one time, plus a number of times equal to that index, plus one additional time if the index was at least 4.*/
func rotateR(s string, x rune) string {

	xi := strings.IndexRune(s, x)

	if xi > 3 {
		xi += 1
	}
	xi++

	return rotate(s, xi, false)

}

//reverse positions X through Y means that the span of letters at indexes X through Y (including the letters at X and Y) should be reversed in order.
func reverse(s string, x, y int) string {

	out := []rune(s)

	for x < y {
		out[x], out[y] = out[y], out[x]
		x++
		y--
	}
	return string(out)

}

//move position X to position Y means that the letter which is at index X should be removed from the string, then inserted such that it ends up at index Y.
func move(s string, x, y int) string {

	out := make([]rune, len(s))
	t := rune(s[x])

	s = s[:x] + s[x+1:] //remove x

	i, j := 0, 0
	for i < len(out) {

		if i == y {
			out[i] = t
		} else {
			out[i] = rune(s[j])
			j++
		}
		i++

	}

	return string(out)

}

func decode(instruction, s string) string {

	var i int
	var j int

	tokens := strings.Split(instruction, " ")

	switch tokens[0] {
	case "swap":

		if tokens[1] == "position" {
			i, _ = strconv.Atoi(tokens[2])
			j, _ = strconv.Atoi(tokens[5])
			s = swapI(s, i, j)
		} else {
			s = swapR(s, rune(tokens[2][0]), rune(tokens[5][0]))
		}

	case "reverse":
		i, _ = strconv.Atoi(tokens[2])
		j, _ = strconv.Atoi(tokens[4])
		s = reverse(s, i, j)

	case "rotate":

		if tokens[1] == "based" {
			s = rotateR(s, rune(tokens[6][0]))
		} else {
			i, _ = strconv.Atoi(tokens[2])
			s = rotate(s, i, tokens[1] == "left")
		}

	case "move":
		i, _ = strconv.Atoi(tokens[2])
		j, _ = strconv.Atoi(tokens[5])
		s = move(s, i, j)

	default:

		panic("You goofed. No instruction for: " + tokens[0])

	}

	return s

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func runDecoder(input, initial string) string {

	for _, instruction := range strings.Split(input, "\n") {

		if instruction != "" {

			initial = decode(instruction, initial)

		}

	}
	return initial

}

//join and permutations sourced from some quick googles
//http://www.golangprograms.com/golang-program-to-print-all-permutations-of-a-given-string.html
func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func part_one(input string) string {

	return runDecoder(input, PLAIN_TEXT)

}

func part_two(input string) string {

	out := ""
	for _, s := range permutations(PLAIN_TEXT) {

		if runDecoder(input, s) == TARGET_TEXT {
			out = s
			break
		}

	}
	return out
}

func main() {

	instructions, err := ioutil.ReadFile("./2016_21.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(instructions)))
	fmt.Println("Problem 2: " + part_two(string(instructions)))

}
