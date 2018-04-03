// Advent of Code 2016 - Day 9
//
//		Wandering around a secure area, you come across a
// datalink port to a new part of the network. After
// briefly scanning it for interesting files, you find one
// file in particular that catches your attention. It's
// compressed with an experimental format, but fortunately,
// the documentation for the format is nearby.
//
// 		The format compresses a sequence of characters.
// Whitespace is ignored. To indicate that some sequence
// should be repeated, a marker is added to the file, like
// (10x2). To decompress this marker, take the subsequent
// 10 characters and repeat them 2 times. Then, continue
// reading the file after the repeated data. The marker
// itself is not included in the decompressed output.
//
// 		If parentheses or other characters appear within
// the data referenced by a marker, that's okay - treat it
// like normal data, not a marker, and then resume looking
// for markers after the decompressed section.
//
//      Part 1:
//
// 		What is the decompressed length of the file (your
// puzzle input)? Don't count whitespace.
//
//      Part 2:
//
//      Apparently, the file actually uses version two of
// the format.
//
//      In version two, the only difference is that markers
// within decompressed data are decompressed. This, the
// documentation explains, provides much more substantial
// compression capabilities, allowing many-gigabyte files to
// be stored in only a few kilobytes.
//
//      Unfortunately, the computer you brought probably
// doesn't have enough memory to actually decompress the
// file; you'll have to come up with another way to get its
// decompressed length.
//
//      What is the decompressed length of the file using
// this improved format?
//
// 		Approach taken in part 1 was to simply parse the
// input and rebuild it as we go. At the end we take the
// len() of the final product. This worked ok though I had
// some issues getting the indexes correct on my first pass.
// Part 2 was deliberately set up to ensure difficulty with
// that approach. I believe the final output would be around
// 10 GB in size if you used the first method. Instead I just
// rewrote the function to be recursive. This was
// unsurprisingly easier to manage than the approach taken in
// the first method.
//
// 		Issues: Off by 1 errors.
//
//

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func decode(s string) (int, int) {

	parts := strings.Split(s, "x")
	length, _ := strconv.Atoi(parts[0])
	count, _ := strconv.Atoi(parts[1])

	return length, count

}

func decompress(s string) string {

	next := 0
	last := 0
	output := ""
	//for i, c := range s {
	i := 0
	for i < len(s) {

		c := s[i]

		switch c {

		case '(':

			output += s[last:i]
			next = i

		case ')':

			length, count := decode(s[next+1 : i])
			output += strings.Repeat(s[i+1:i+1+length], count)
			last = i + 1 + length //offset to trim the '('
			i += length

		}
		i++

	}

	if last < len(s) {
		output += s[last:len(s)]
	}

	return output

}

func decompressEx(s string) int {

	next := 0
	last := 0
	output := 0

	i := 0
	for i < len(s) {

		c := s[i]

		switch c {

		case '(':

			next = i
			output += next - last //count all the characters not recursed

		case ')':

			length, count := decode(s[next+1 : i])
			//multiply the substring by the length of the recursed element
			output += count * decompressEx(s[i+1:i+length+1])
			last = i + 1 + length //offset to trim the '('
			i += length
		}
		i++

	}

	output += i - last

	return output

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	l := 0

	sequences := strings.Split(input, "\n")
	for _, seq := range sequences {

		if seq != "" {
			l = len(decompress(seq))
		}
	}

	return strconv.Itoa(l)

}

func part_two(input string) string {

	l := 0

	sequences := strings.Split(input, "\n")
	for _, seq := range sequences {

		if seq != "" {
			l = decompressEx(seq)
		}
	}

	return strconv.Itoa(l)

}

func main() {

	input, err := ioutil.ReadFile("./2016_09.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
