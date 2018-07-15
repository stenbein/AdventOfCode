// Advent of Code 2016 - Day 14
//
// 		Approach taken was to keep track of only the triples
// to attempt to conserve memory. This worked ok but I'm sure
// there are more ways to improve on this.
//
// 		Issues: This one was worse than day
// 11 in the ammount of time it took to complete. The issue
// was an off by 1 error that caused the check if hasTriple or
// hasFives to return false in some cases. Very hard debug,
// ended up rebuilding it twice before I caught the problem.
// Could have been solved by using a complete test case
// coverage on the minor functions in the approach.
//
//

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type hash struct {
	index int
	value string

	isTriple bool
	isFiver  bool

	tripVal  string
	fivesVal string
}

type hashList struct {
	size     int
	lastHash int
	salt     string
	list     []hash
}

//todo replace with regex
func hasTriple(s string) (string, bool) {

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return string(s[i]), true
		}
	}

	return "", false

}

//todo replace with regex
func hasFives(s string) (string, bool) {

	for i := 0; i < len(s)-4; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] &&
			s[i] == s[i+3] && s[i] == s[i+4] {
			return string(s[i]), true
		}
	}

	return "", false

}

func (hl *hashList) next() {

	i := hl.lastHash + 1
	for {

		data := []byte(hl.salt + strconv.Itoa(i))
		md5hash := md5.Sum(data)
		strHash := hex.EncodeToString(md5hash[:])

		if tri, ok := hasTriple(strHash); ok == true {

			five, fok := hasFives(strHash)
			h := hash{i, strHash, true, fok, tri, five}
			hl.list = append(hl.list, h)
			hl.lastHash = i
			break
		}
		i++

	}
	hl.size++
}

func (hl *hashList) nextXL() {

	var strHash string

	i := hl.lastHash + 1
	for {

		data := []byte(hl.salt + strconv.Itoa(i))
		md5hash := md5.Sum(data)
		strHash = hex.EncodeToString(md5hash[:])

		for j := 0; j < 2016; j++ {
			data := []byte(strHash)
			md5hash := md5.Sum(data)
			strHash = hex.EncodeToString(md5hash[:])

		}

		if tri, ok := hasTriple(strHash); ok == true {

			five, fok := hasFives(strHash)
			h := hash{i, strHash, true, fok, tri, five}
			hl.list = append(hl.list, h)
			hl.lastHash = i
			break
		}
		i++

	}
	hl.size++
}

//finds the next key from the hash at index i
func (hl *hashList) nextKey(lastIndex int) int {

	//recover the last index
	i := 0
	for j, h := range hl.list {
		if h.index > lastIndex {
			i = j
			break
		}
	}

	for hl.size <= i {
		hl.next()
	}

loop:
	for {

		if hl.list[i].isTriple {
			//catch up the list
			for hl.list[i].index+1000 > hl.list[hl.size-1].index {
				hl.next()
			}
			j := i + 1
			for hl.list[i].index+1000 > hl.list[j].index {
				if hl.list[j].isFiver && (hl.list[i].tripVal[0] == hl.list[j].fivesVal[0]) {
					break loop
				}
				j++
			}

			i++
		}

	}
	return hl.list[i].index
}

//finds the next key from the hash at index i
func (hl *hashList) nextKeyXL(lastIndex int) int {

	//recover the last index
	i := 0
	for j, h := range hl.list {
		if h.index > lastIndex {
			i = j
			break
		}
	}

	for hl.size <= i {
		hl.nextXL()
	}

loop:
	for {

		if hl.list[i].isTriple {
			//catch up the list
			for hl.list[i].index+1000 > hl.list[hl.size-1].index {
				hl.nextXL()
			}
			j := i + 1
			for hl.list[i].index+1000 > hl.list[j].index {
				if hl.list[j].isFiver && (hl.list[i].tripVal[0] == hl.list[j].fivesVal[0]) {
					break loop
				}
				j++
			}

			i++
		}

	}
	return hl.list[i].index
}

func part_one(input string) string {

	hl := hashList{}
	hl.salt = input
	hl.next()

	index := 0
	for i := 1; i < 64; i++ { //

		index = hl.nextKey(index)

	}

	return strconv.Itoa(index)

}

func part_two(input string) string {

	hl := hashList{}
	hl.lastHash = -1
	hl.salt = input
	hl.nextXL()

	index := 0
	for i := 1; i <= 64; i++ { //

		index = hl.nextKeyXL(index)

	}

	return strconv.Itoa(index)

}

func main() {

	salt := "zpqevtbw"

	fmt.Println("Problem 1: " + part_one(salt))

	//start := time.Now() //timestamp

	fmt.Println("Problem 2: " + part_two(salt))

	//Elapse time
	//elapsed := time.Since(start)
	//fmt.Printf("Execution took %s\n", elapsed)

}
