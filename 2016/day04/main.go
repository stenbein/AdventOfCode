// Advent of Code 2016 - Day 4
//
//		Each room consists of an encrypted name (lowercase
// letters separated by dashes) followed by a dash, a sector
// ID, and a checksum in square brackets.
//
//		A room is real (not a decoy) if the checksum is the
// five most common letters in the encrypted name, in order,
// with ties broken by alphabetization.
//
//		What is the sum of the sector IDs of the real rooms?
//
// 		Approach taken was to itterate through the given data
// generating
//
// 		Issues: Forgot to trim final value of checksum which left
// "]" as a validation character. In the second part I had
// an error with my ROTx where the modulo was in the wrong
// place.
//
//

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type RoomCode struct {
	code     string
	id       string
	checksum string
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {

	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value

}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (c *RoomCode) isValid() bool {

	//take a map of the frequencies
	//todo, try this without a map after 2016 is done
	counts := make(map[string]int)

	for _, char := range c.code {
		counts[string(char)] += 1
	}

	//sort the map as a slice by values
	countList := make(PairList, len(counts))
	i := 0
	for k, val := range counts {
		countList[i] = Pair{k, val}
		i++
	}
	sort.Sort(sort.Reverse(countList))

	//check the result list against the checksum
	for i, char := range c.checksum {
		if countList[i].Key != string(char) {
			return false
		}
	}

	return true

}

func (c *RoomCode) decode() string {

	id, _ := strconv.Atoi(c.id)
	shift := id

	rotx := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+rune(shift))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(shift))%26
		}
		return r
	}

	return strings.Map(rotx, c.code)

}

func newCode(raw string) RoomCode {

	parts := strings.Split(raw, "-")

	code := parts[:len(parts)-1]
	ID := strings.Split(parts[len(parts)-1], "[")[0]
	checksum := strings.Split(parts[len(parts)-1], "[")[1]

	//trim off the final character of the checksum
	checksum = checksum[:len(checksum)-1]

	return RoomCode{strings.Join(code, ""), ID, checksum}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	total := 0

	codes := strings.Split(input, "\n")
	for _, code := range codes {

		if code != "" {

			c := newCode(code)
			if c.isValid() {
				id, _ := strconv.Atoi(c.id)
				total += id
			}

		}

	}

	return strconv.Itoa(total)

}

func part_two(input string) string {

	valid := make([]RoomCode, 0)

	raws := strings.Split(input, "\n")
	for _, code := range raws {

		if code != "" {

			c := newCode(code)

			//todo - same with the map above
			//try this without append
			valid = append(valid, c)

		}

	}

	//because I split on the hypens earlier
	//this doesn't have spaces. I peaked at the
	//format by taking a dump during testing
	for _, c := range valid {

		if c.decode() == "northpoleobjectstorage" {
			return c.id
		}
	}

	return "Not found"

}

func main() {

	input, err := ioutil.ReadFile("./2016_04.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
