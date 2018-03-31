// Advent of Code 2016 - Day 6
//
//		Something is jamming your communications with Santa.
// Fortunately, your signal is only partially jammed, and
// protocol in situations like this is to switch to a simple
// repetition code to get the message through.
//
// 		In this model, the same message is sent repeatedly.
// You've recorded the repeating message signal (your puzzle
// input), but the data seems quite corrupted - almost too
//badly to recover. Almost.
//
// 		Part 1:
//
// 		All you need to do is figure out which character is
// most frequent for each position. Given the recording in
// your puzzle input, what is the error-corrected version
// of the message being sent?
//
// 		Part 2:
//
//		In this modified code, the sender instead transmits
// what looks like random data, but for each character, the
// character they actually want to send is slightly less
// likely than the others. Even after signal-jamming noise,
// you can look at the letter distributions in each column
// and choose the least common letter to reconstruct the
// original message.
//
// 		Approach taken was to consider each character as a
// rune and generate a list of counts. I wanted to avoid
// using maps because they've been bothering me lately as
// a style choice. Instead here we have slices generated
// during the run to contain the runes and counts.
//
// 		Issues: off by one error on the slice generation
// allowed illegial input into the mostCommon function.
// And I made a mistake assuming I was being passed a
// string instead of []rune.
//

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//input is entirely lower case letters
//use this to create an array to count the
//values of each rune. Input must have
//valid runes a-z inclusive. error
//was caused by input of ''
func mostCommon(runes []rune) string {

	counts := [26]rune{}
	for _, r := range runes {
		counts[r-'a']++
	}

	max := 0
	for i, count := range counts {
		if count > counts[max] {
			max = i
		}
	}

	return string(rune('a' + max))

}

func leastCommon(runes []rune) string {

	counts := [26]rune{}
	for _, r := range runes {
		counts[r-'a']++
	}

	min := 0
	for i, count := range counts {
		if count < counts[min] {
			min = i
		}
	}

	return string(rune('a' + min))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	lines := strings.Split(input, "\n")

	columns := make([][]rune, len(lines[0]))
	for i, _ := range columns {
		columns[i] = make([]rune, len(lines)-1) //else slice has rune of ''
	}

	for i, line := range lines {

		for j, char := range line {
			columns[j][i] = rune(char)
		}

	}

	output := ""
	for _, column := range columns {
		output += mostCommon(column[:])
	}

	return output

}

func part_two(input string) string {

	lines := strings.Split(input, "\n")

	columns := make([][]rune, len(lines[0]))
	for i, _ := range columns {
		columns[i] = make([]rune, len(lines)-1) //else slice has rune of ''
	}

	for i, line := range lines {

		for j, char := range line {
			columns[j][i] = rune(char)
		}

	}

	output := ""
	for _, column := range columns {
		output += leastCommon(column[:])
	}

	return output

}

func main() {

	input, err := ioutil.ReadFile("./2016_06.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
