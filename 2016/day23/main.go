// Advent of Code 2016 - Day 23
//
// 		Approach taken was to build out a register
// machine and run the program.
//
// 		Issues: Missed the section which said if
// the toggle produces an invalid program, to ignore it.
// And missed the bit about toggling only valid ranges,
// which didn't immediately produce a bug. Also the
// second part runs pretty slowly. Over 10 minutes on my
// machine using this register machine. TODO, check what
// the program is doing and see if there's a faster way
// to do it.
//

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	id    string
	size  int
	used  int
	avail int
}

type instruction struct {
	command string
	arg     []string
}

type registerMachine struct {
	c int            //counter
	r map[string]int //registers
	p []instruction  //the program

}

func (rm *registerMachine) jnz(x, y string) {

	var test int
	if i, err := strconv.Atoi(x); err != nil {
		test = rm.get(x)
	} else {
		test = i
	}

	j := rm.get(y) //always expected to be valid int
	if test != 0 {
		rm.c += j - 1 //offset for the register increment
	}
}

func (rm *registerMachine) inc(x string) {
	rm.r[x]++
}

func (rm *registerMachine) dec(x string) {
	rm.r[x]--
}

func (rm *registerMachine) copy(x, y string) {
	if _, err := strconv.Atoi(y); err != nil {
		rm.r[y] = rm.get(x)
	}
}

func (rm *registerMachine) toggle(x string) {

	i := rm.get(x) //always expected to be valid int
	if rm.c+i < len(rm.p) {
		inst := rm.p[rm.c+i]
		switch len(inst.arg) {
		case 1:
			if inst.command == "inc" {
				inst.command = "dec"
			} else {
				inst.command = "inc"
			}
			rm.p[rm.c+i] = inst
		case 2:
			if inst.command == "jnz" {
				inst.command = "cpy"
			} else {
				inst.command = "jnz"
			}
			rm.p[rm.c+i] = inst
		}
	}
}

//returns either the value of the register x if x is a string
//or the value of x if x is an int
func (rm *registerMachine) get(x string) int {

	if i, err := strconv.Atoi(x); err == nil {
		return i
	} else {
		return rm.r[x]
	}

}

func (rm *registerMachine) step(i int) {

	inst := rm.p[i]

	switch inst.command {
	case "cpy":
		rm.copy(inst.arg[0], inst.arg[1])
	case "inc":
		rm.inc(inst.arg[0])
	case "dec":
		rm.dec(inst.arg[0])
	case "jnz":
		rm.jnz(inst.arg[0], inst.arg[1])
	case "tgl":
		rm.toggle(inst.arg[0])

	}

}

func (rm *registerMachine) run() {

	for rm.c < len(rm.p) {
		rm.step(rm.c)
		rm.c++
	}

}

func NewMachine(p []instruction) registerMachine {

	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 0
	m["c"] = 0
	m["d"] = 0
	m["e"] = 0

	return registerMachine{0, m, p}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	program := []instruction{}
	for _, s := range strings.Split(input, "\n") {
		if s != "" {
			tokens := strings.Split(s, " ")
			program = append(program, instruction{tokens[0], tokens[1:]})
		}
	}

	rm := NewMachine(program)
	rm.r["a"] = 7 //special input for this problem

	rm.run()

	return strconv.Itoa(rm.get("a"))

}

func part_two(input string) string {

	program := []instruction{}
	for _, s := range strings.Split(input, "\n") {
		if s != "" {
			tokens := strings.Split(s, " ")
			program = append(program, instruction{tokens[0], tokens[1:]})
		}
	}

	rm := NewMachine(program)
	rm.r["a"] = 12 //special input for this problem

	rm.run()

	return strconv.Itoa(rm.get("a"))

}

func main() {

	input, err := ioutil.ReadFile("./2016_23.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
