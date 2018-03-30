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
// Issues: Forgot to trim final value of checksum which left
// "]" as a validation character. In the second part I had
// an error with my ROTx where the modulo was in the wrong 
// place.
//
//

package main

 

import (

        "fmt"
        "strconv"
        "crypto/md5"
        "encoding/hex"

)



func check(e error) {
    if e != nil {
        panic(e)
    }
}

func part_one(input string) string {

	digits := 8
	i := 0

	for digits > 0 {

		data := []byte(input + strconv.Itoa(i))
		md5hash := md5.Sum(data)

		n := int(uint(md5hash[2]) | uint(md5hash[1])<<8 | uint(md5hash[0])<<16)

		if n < 16 {

			fmt.Println(md5hash[:3])
			digits--

		}
		i++

	}

    return "Complete"

}

func part_two(input string) string {

	digits := 15
	i := 0

	for digits > 0 {

		data := []byte(input + strconv.Itoa(i))
		md5hash := md5.Sum(data)

		n := int(uint(md5hash[2]) | uint(md5hash[1])<<8 | uint(md5hash[0])<<16)

		if n < 8 {

			fmt.Println(md5hash[:4])
			fmt.Println(hex.EncodeToString(md5hash[:4]))
			digits--

		}
		i++

	}

    return "Complete"

}

func main() {

	input := "abbhdwsy"

	fmt.Println("Problem 1: " + part_one(string(input))) //input here is single text string
	fmt.Println("Problem 2: " + part_two(string(input)))

}
