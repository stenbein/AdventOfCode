package main

import (
	"fmt"
	"testing"
)

/*
func (rm *registerMachine) rep() {

	fmt.Println("Counter: ", rm.c)
	fmt.Println(rm.r)
	fmt.Println(rm.p)
	fmt.Println("*******")

}
*/
//check that the init grid includes the start point
func ensureRegister(t *testing.T, rm *registerMachine, register string, expect int) {

	rm.run()
	if v := rm.get(register); v != expect {
		t.Errorf("rm.get(%s) = %d, want %d", register, v, expect)
	}

}

func ensureProgram(t *testing.T, rm *registerMachine, i int, expect string) {

	rm.run()
	if v := rm.p[i].command; v != expect {
		t.Errorf("rm.p(%d) = %s, want %s", i, v, expect)
	}

}

func TestInstructions(t *testing.T) {

	var inst1 instruction
	var inst2 instruction
	var inst3 instruction
	var inst4 instruction
	var rm registerMachine

	inst1 = instruction{"inc", []string{"a"}}
	rm = NewMachine([]instruction{inst1})
	ensureRegister(t, &rm, "a", 1)
	ensureRegister(t, &rm, "b", 0)

	inst1 = instruction{"inc", []string{"a"}}
	inst2 = instruction{"dec", []string{"a"}}
	rm = NewMachine([]instruction{inst1, inst2})
	ensureRegister(t, &rm, "a", 0)
	ensureRegister(t, &rm, "b", 0)

	inst1 = instruction{"inc", []string{"a"}}
	inst2 = instruction{"cpy", []string{"a", "b"}}
	rm = NewMachine([]instruction{inst1, inst2})
	ensureRegister(t, &rm, "a", 1)
	ensureRegister(t, &rm, "b", 1)

	inst1 = instruction{"inc", []string{"a"}}
	inst2 = instruction{"inc", []string{"a"}}
	inst3 = instruction{"jnz", []string{"b", "-1"}}
	rm = NewMachine([]instruction{inst1, inst2, inst3})
	ensureRegister(t, &rm, "a", 2)
	ensureRegister(t, &rm, "b", 0)

	inst1 = instruction{"dec", []string{"a"}}
	inst2 = instruction{"dec", []string{"a"}}
	inst3 = instruction{"inc", []string{"a"}}
	inst4 = instruction{"jnz", []string{"a", "-1"}}
	rm = NewMachine([]instruction{inst1, inst2, inst3, inst4})
	ensureRegister(t, &rm, "a", 0)
	ensureRegister(t, &rm, "b", 0)

	fmt.Println("Test complete...")

}

func TestToggle(t *testing.T) {

	var inst1 instruction
	var inst2 instruction
	//var inst3 instruction
	//var inst4 instruction
	var rm registerMachine

	inst1 = instruction{"inc", []string{"a"}}
	inst2 = instruction{"tgl", []string{"-1"}}
	rm = NewMachine([]instruction{inst1, inst2})
	ensureProgram(t, &rm, 0, "dec")

	fmt.Println("Test complete...")

}
