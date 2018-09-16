// Advent of Code 2016 - Day 5
//
//		The eight-character password for the door is
// generated one character at a time by finding the MD5 hash
// of some Door ID (your puzzle input) and an increasing
// integer index (starting with 0).
//
// 		Part 1:
//
//		A hash indicates the next character in the password
// if its hexadecimal representation starts with five zeroes.
// If it does, the sixth character in the hash is the next
// character of the password.
//
//		Given the actual Door ID, what is the password?
//
// 		Part 2:
//
// 		Instead of simply filling in the password from left
// to right, the hash now also indicates the position within
// the password to fill. You still look for hashes that
// begin with five zeroes; however, now, the sixth character
// represents the position (0-7), and the seventh character
// is the character to put in that position.
//
// 		Approach taken was to itterate through the given data
// generating hashes as we go. Hashes are converted to ints,
// checked against the target size and printed if they match.
// I manually typed the password from that.
//
// 		Modified the outputs to produce strings like the other
// days. Cleaned up the code a bit.

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	digits := 8
	i := 0
	out := ""

	for digits > 0 {

		data := []byte(input + strconv.Itoa(i))
		hash := md5.Sum(data)

		n := int(uint(hash[2]) | uint(hash[1])<<8 | uint(hash[0])<<16)

		if n < 16 {

			s := hex.EncodeToString(hash[:3])
			out += string(s[5])
			digits--

		}
		i++

	}

	return out

}

func part_two(input string) string {

	digits := 8
	i := 0

	var out [8]string

	for digits > 0 {

		data := []byte(input + strconv.Itoa(i))
		hash := md5.Sum(data)

		n := int(uint(hash[2]) | uint(hash[1])<<8 | uint(hash[0])<<16)

		if n < 8 {

			if out[n] == "" {
				s := hex.EncodeToString(hash[:4])
				out[n] = string(s[6])
				digits--
			}

		}
		i++

	}

	return strings.Join(out[:], "")

}

func main() {

	input := "abbhdwsy" //input here is single text string

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
