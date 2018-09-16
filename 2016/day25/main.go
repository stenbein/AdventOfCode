// Advent of Code 2016 - Day 25
//
// 		Approach taken was to copy day 23 and add
// the code for the out op from the day 25 instructions.
//
// 		Issues: None, worked right away and very fast.
// I was disappointed. Expecting it to take longer to
// run I had also solved the ASM code by hand.
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
	o []int          //the output to test for day 25

}

//out x transmits x (either an integer or the value of a register) as the next value for the clock signal.

func (rm *registerMachine) out(x string) {

	rm.o = append(rm.o, rm.get(x))

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
	//added for day 25
	case "out":
		rm.out(inst.arg[0])
	}

}

func (rm *registerMachine) run() {

	for rm.c < len(rm.p) && len(rm.o) < 10 {
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

	return registerMachine{0, m, p, []int{}}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkOut(o []int) bool {

	for i := 0; i < len(o)-1; i++ {
		if o[i] == o[i+1] {
			return false
		}
	}
	return true

}

func part_one(input string) string {

	program := []instruction{}
	for _, s := range strings.Split(input, "\n") {
		if s != "" {
			tokens := strings.Split(s, " ")
			program = append(program, instruction{tokens[0], tokens[1:]})
		}
	}

	i := 0
	for {

		rm := NewMachine(program)
		rm.r["a"] = i //special input for this problem

		rm.run()

		if checkOut(rm.o) {
			fmt.Println(i, rm.o)
			break
		}

		i++

	}

	return "Program Complete."

}

//no part 2 for day 25
/*func part_two(input string) string {



}*/

func main() {

	input, err := ioutil.ReadFile("./2016_25.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	//fmt.Println("Problem 2: " + part_two(string(input)))

}
