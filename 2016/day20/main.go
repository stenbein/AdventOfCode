// Advent of Code 2016 - Day 20
//
// 		Approach taken was to sort the array by the min of
// the ip range. Itterate once through to deal with overlaps
// and then count the outputs.
//
// 		Issues: I was surprised to find that the structure
// for i, r := range ipRanges is a copy and not a pointer
// to the IP range r. So updating the values dynamically
// didn't work. Instead you need to declare a pointer, or
// use ipRanges[i].
//

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const (
	MAX_IP int = 4294967295
)

type IPs struct {
	min int
	max int
}

func makeIPs(input string) []IPs {

	ranges := strings.Split(input, "\n")
	ipRanges := make([]IPs, len(ranges))
	for i, r := range ranges {
		if r != "" {
			parts := strings.Split(r, "-")
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			ipRanges[i] = IPs{min, max}
		}
	}

	sort.Sort(byMin(ipRanges))

	//adjust for overlapping ranges
	boundary := 0
	for i, _ := range ipRanges {

		if ipRanges[i].max > boundary {
			boundary = ipRanges[i].max
		} else {
			ipRanges[i].max = boundary
		}

	}

	return ipRanges

}

type byMin []IPs

func (r byMin) Len() int {
	return len(r)
}
func (r byMin) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r byMin) Less(i, j int) bool {
	return r[i].min < r[j].min
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	ipRanges := makeIPs(input)

	lastMin := 0
	for _, r := range ipRanges {

		//ranges are inclusive so check if we're better than last+1
		if r.min > lastMin {
			break
		} else {
			lastMin = r.max + 1
		}

	}

	return strconv.Itoa(lastMin)

}

func part_two(input string) string {

	ipRanges := makeIPs(input)

	total, lastMin := 0, 0
	for _, r := range ipRanges {

		if r.min > lastMin {
			total += r.min - lastMin
		}
		lastMin = r.max + 1

	}

	if lastMin < MAX_IP {
		total += (MAX_IP - lastMin) + 1
	}

	return strconv.Itoa(total)

}

func main() {

	input, err := ioutil.ReadFile("./2016_20.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
