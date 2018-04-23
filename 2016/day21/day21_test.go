package main

import (
	"fmt"
	"testing"
)

//check the value of the boolean slice for the doors
func ensureDecode(t *testing.T, instruction, testString, expectValue string) {

	if v := decode(instruction, testString); v != expectValue {
		fmt.Println(instruction)
		t.Errorf("decode() = %s, want %s", v, expectValue)
	}

}

func TestDecode(t *testing.T) {

	ensureDecode(t, "swap position 0 with position 4", "abcde", "ebcda")

	ensureDecode(t, "swap position 0 with position 1", "ab", "ba")
	ensureDecode(t, "swap position 1 with position 0", "ab", "ba")
	ensureDecode(t, "swap position 0 with position 2", "acb", "bca")

	ensureDecode(t, "swap letter d with letter b", "ebcda", "edcba")

	ensureDecode(t, "swap letter a with letter b", "ab", "ba")
	ensureDecode(t, "swap letter a with letter b", "abc", "bac")
	ensureDecode(t, "swap letter a with letter b", "acb", "bca")

	ensureDecode(t, "rotate left 1 step", "abcde", "bcdea")
	ensureDecode(t, "rotate right 8 step", "ab", "ab")
	ensureDecode(t, "rotate left 8 step", "ab", "ab")
	ensureDecode(t, "rotate right 8 step", "abc", "bca")
	ensureDecode(t, "rotate left 8 step", "abc", "cab")

	ensureDecode(t, "rotate right 1 step", "abcde", "eabcd")
	ensureDecode(t, "rotate right 4 step", "abcde", "bcdea")

	ensureDecode(t, "move position 1 to position 4", "bcdea", "bdeac")
	ensureDecode(t, "move position 1 to position 3", "bcdea", "bdeca")

	ensureDecode(t, "move position 3 to position 0", "bdeac", "abdec")
	ensureDecode(t, "move position 4 to position 0", "bdeac", "cbdea")

	ensureDecode(t, "rotate based on position of letter b", "abdec", "ecabd")
	ensureDecode(t, "rotate based on position of letter a", "abdec", "cabde")

	ensureDecode(t, "rotate based on position of letter d", "ecabd", "decab")

	//here was my error with part 1, off by 1 caused by incrementing the index
	//before the test was completed to see if the index was >= 4
	ensureDecode(t, "rotate based on position of letter d", "ecadb", "cadbe")

}
