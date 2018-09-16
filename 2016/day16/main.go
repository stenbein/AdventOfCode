// Advent of Code 2016 - Day 16
//
//		You're done scanning this part of the network,
// but you've left traces of your presence. You need
// to overwrite some disks with random-looking data to
// cover your tracks and update the local security
// system with a new checksum for those disks.
//
//		For the data to not be suspicious, it needs to
// have certain properties; purely random data will be
// detected as tampering. To generate appropriate
// random data, you'll need to use a modified dragon
// curve.
//
//		Start with an appropriate initial state (your
// puzzle input). Then, so long as you don't have
// enough data yet to fill the disk, repeat the
// following steps:
//
//		Call the data you have at this point "a".
//		Make a copy of "a"; call this copy "b".
//		Reverse the order of the characters in "b".
//		In "b", replace all instances of 0 with 1
//			and all 1s with 0.
// 		The resulting data is "a", then a single
//			0, then "b".
//
// 		Approach: Naive approch here worked ok. But
// ran over night without finished on part two.
// After checking the run time of each function,
// the reduce step was taking too long due to due
// memory reallocation on each string concat. Swapped
// out the string + string loop for a pre defined
// slice of runes to populate and it finished in 2
// seconds. I reworked it to use the bytes library
// after that.
//
// 		Issues: The speed of the string concat is
// something I haven't had to deal with during
// the 2016 advent. A quick google search brings up
// many builtin string buffer options. TODO:
// explore the options a bit more with benchmarks.
//

package main

import (
	"bytes"
	"fmt"
)

func SwapBinary(s string) string {

	runes := []rune(s)
	for i, r := range s {
		switch r {
		case '0':
			runes[i] = '1'
		case '1':
			runes[i] = '0'
		}
	}
	return string(runes)

}

func Reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

//Call the data you have at this point "a".
//Make a copy of "a"; call this copy "b".
//Reverse the order of the characters in "b".
//In "b", replace all instances of 0 with 1 and all 1s with 0.
//The resulting data is "a", then a single 0, then "b".
func fill(l int, pattern string) string {

	for {
		if len(pattern) < l {

			pattern = pattern + "0" + SwapBinary(Reverse(pattern))

		} else {
			break
		}

	}
	return pattern[:l] //cut to size

}

func reduce(s string) string {

	//output := make([]rune, len(s)/2)
	var b bytes.Buffer

	for i, j := 0, 0; i < len(s); i, j = i+2, j+1 {

		if s[i] == s[i+1] {
			b.WriteString("1")
		} else {
			b.WriteString("0")
		}
	}
	return b.String()

}

func checksum(s string) string {

	if len(s)%2 == 0 {
		s = checksum(reduce(s))
	}
	return s

}

func part_one(input string, size int) string {

	return checksum(fill(size, input))

}

func part_two(input string, size int) string {

	return checksum(fill(size, input))

}

func main() {

	raw := "00101000101111010"

	fmt.Println("Problem 1: " + part_one(raw, 272))
	fmt.Println("Problem 2: " + part_two(raw, 35651584))

}
