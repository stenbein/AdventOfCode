// Advent of Code 2016 - Day
//
//		According to these documents, Easter Bunny HQ isn't
// just this building - it's a collection of buildings in the
// nearby area. They're all connected by a local monorail,
// and there's another building not far from here!
// Unfortunately, being night, the monorail is currently not
// operating.
//
// 		You remotely connect to the monorail control systems
// and discover that the boot sequence expects a password.
// The password-checking logic (your puzzle input) is easy
// to extract, but the code it uses is strange: it's
// assembunny code designed for the new computer you just
// assembled. You'll have to execute the code and get the
// password.
//
//		The assembunny code you've extracted operates on four
// registers (a, b, c, and d) that start at 0 and can hold
// any integer. However, it seems to make use of only a few
// instructions:
//
// 		cpy x y copies x (either an integer or the value of a
// register) into register y.
// 		inc x increases the value of register x by one.
// 		dec x decreases the value of register x by one.
// 		jnz x y jumps to an instruction y away (positive
// means forward; negative means backward), but only if x is
// not zero.
// 		The jnz instruction moves relative to itself: an
// offset of -1 would continue at the previous instruction,
// while an offset of 2 would skip over the next instruction.
//
//		Part 1: After executing the assembunny code in your
// puzzle input, what value is left in register a?
//
//		Part 2: If you instead initialize register c to be
// 1, what value is now left in register a?
//
// 		Approach taken was to simulate a register machine
// and build each instruction as a function acting on it.
//
// 		Issues: Off by 1 error again on the program counter.
// For part two the ASM runs slowly enough that I couldn't
// complete this remotely using the go playground. Instead
// I did it by hand like I did when I did the 2017 challanges.
// Once I got home I found out that the same approach as
// part 1 worked well enough to solve it in about 4 seconds.
//
//
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type registerMachine struct {

	//instructions interface{
	program   []string
	counter   int
	registers [4]int
}

func (r *registerMachine) getR(s string) int {

	switch s {
	case "a":
		return r.registers[0]

	case "b":
		return r.registers[1]

	case "c":
		return r.registers[2]

	case "d":
		return r.registers[3]

	default:
		return r.registers[3]

	}

}

func (r *registerMachine) setR(s string, v int) {

	switch s {
	case "a":
		r.registers[0] = v

	case "b":
		r.registers[1] = v

	case "c":
		r.registers[2] = v

	case "d":
		r.registers[3] = v

	}

}

func (r *registerMachine) jump(x string, y string) {

	jumpVal, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}

	jumpVal--                 //offset for later incremet
	v, err := strconv.Atoi(x) //if error, assume it's a register
	if err != nil {

		v := r.getR(x)
		if v != 0 {
			r.counter += jumpVal
		}

	} else {

		if v != 0 {
			r.counter += jumpVal
		}

	}

}

func (r *registerMachine) copy(x string, y string) {

	v, err := strconv.Atoi(x) //if error, assume it's a register
	if err != nil {
		v := r.getR(x)
		r.setR(y, v)
	} else {
		r.setR(y, v)
	}

}

func (r *registerMachine) decode(s string) {

	parts := strings.Split(s, " ")
	switch parts[0] {

	case "cpy":
		r.copy(parts[1], parts[2])

	case "inc":
		r.setR(parts[1], r.getR(parts[1])+1)

	case "dec":
		r.setR(parts[1], r.getR(parts[1])-1)

	case "jnz":
		r.jump(parts[1], parts[2])

	}

}

func (r *registerMachine) start(i int) {

	r.counter = i
	for r.counter < len(r.program) {

		inst := r.program[r.counter]
		r.decode(inst)

		r.counter++

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	var program []string
	program = make([]string, 0)

	for _, inst := range strings.Split(input, "\n") {
		if inst != "" {
			program = append(program, inst)
		}

	}

	r := registerMachine{}
	r.program = program
	r.start(0)

	return strconv.Itoa(r.getR("a"))

}

func part_two(input string) string {

	var program []string
	program = make([]string, 0)

	for _, inst := range strings.Split(input, "\n") {
		if inst != "" {
			program = append(program, inst)
		}

	}

	r := registerMachine{}

	r.setR("c", 1) //extra condition for part 2

	r.program = program
	r.start(0)

	return strconv.Itoa(r.getR("a"))

}

func main() {

	input, err := ioutil.ReadFile("./2016_12.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
