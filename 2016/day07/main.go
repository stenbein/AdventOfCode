// Advent of Code 2016 - Day 7
//
//		While snooping around the local network of EBHQ,
// you compile a list of IP addresses (they're IPv7, of
// course; IPv6 is much too limited). You'd like to figure
// out which IPs support TLS (transport-layer snooping).
//
//		Part 1:
//
// 		An IP supports TLS if it has an Autonomous Bridge
// Bypass Annotation, or ABBA. An ABBA is any four-character
// sequence which consists of a pair of two different
// characters followed by the reverse of that pair, such
// as xyyx or abba. However, the IP also must not have an
// ABBA within any hypernet sequences, which are contained
// by square brackets.
//
//		How many IPs in your puzzle input support TLS?
//
//		What is the sum of the sector IDs of the real rooms?
//
// 		Approach taken for part one is to split the address
// into components and test each for ABBA.
//
// 		Issues: Can't test the components as a set. Rather
// by set of sets divided by delimiters 1[2]3[4]5 instead
// of 1[2]1[2]1[2]1... etc, then I passed blank lines to
// the constructor, then I had a final issue with a the
// final value of the final string outside of the brackets
// being cut short. To overcome this I had to run tests on
// the input to ensure it conformed to the spec.
//

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Address struct {
	hypernet []string
	supernet []string
}

//there may be multiple sets of brackets '[]'
//assume brackets are balanced '[' before ']'
func newAddress(raw string) Address {

	var hyp []string
	var sup []string

	last := 0
	for i, c := range raw {

		switch c {
		case '[':
			sup = append(sup, raw[last:i])
			last = i + 1
		case ']':
			hyp = append(hyp, raw[last:i])
			last = i + 1
		}

	}

	//get last bit
	sup = append(sup, raw[last:len(raw)])

	return Address{hyp, sup}

}

func Any(vs []string, f func(string) bool) bool {

	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false

}

func isABBA(s string) bool {

	//check for length
	if len(s) < 4 {
		return false
	}

	for i := 0; i <= len(s)-4; i++ {

		subs := s[i : i+4]
		//exclude all same
		if subs[0] == subs[3] &&
			subs[1] == subs[2] &&
			subs[0] != subs[1] {

			return true
		}

	}

	return false //else

}

func getABAs(as []string) []string {

	results := []string{}

	for _, s := range as {

		if len(s) < 3 {
			continue
		}

		for i := 0; i <= len(s)-3; i++ {

			subs := s[i : i+3]
			//exclude all same
			if subs[0] == subs[2] &&
				subs[0] != subs[1] {

				results = append(results, subs)
			}

		}
	}

	return results

}

func (a *Address) supportsTLS() bool {

	return Any(a.supernet, isABBA) && !Any(a.hypernet, isABBA)

}

func (a *Address) supportsSSL() bool {

	ABAs := getABAs(a.supernet)
	BABs := getABAs(a.hypernet) // ABA and BAB are the same pattern

	for _, ABA := range ABAs {

		for _, BAB := range BABs {

			if ABA[0] == BAB[1] && ABA[1] == BAB[0] {
				return true
			}
		}

	}

	return false

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	count := 0

	addresses := strings.Split(input, "\n")

	for _, raw := range addresses {

		if raw != "" {
			a := newAddress(raw)

			if a.supportsTLS() {
				count++
			}
		}
	}

	return strconv.Itoa(count)

}

func part_two(input string) string {

	count := 0

	addresses := strings.Split(input, "\n")

	for _, raw := range addresses {

		if raw != "" {
			a := newAddress(raw)

			if a.supportsSSL() {
				count++
			}
		}
	}

	return strconv.Itoa(count)

}

func main() {

	input, err := ioutil.ReadFile("./2016_07.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
