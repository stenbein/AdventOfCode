package main

import (
	//"fmt"
	"testing"
)

//
func checkIsValid(t *testing.T, s *state, expect bool) {

	if l := s.isValid(); l != expect {
		t.Errorf("s.isValid() = %b, want %b", l, expect)
	}

}

func checkState(t *testing.T, s *state, expect [5]int) {

	if l := s.key(); l != expect {
		t.Errorf("s.isValid() = %b, want %b", l, expect)
	}

}

func TestState(t *testing.T) {

	c1 := chip("Hydrogen")
	c2 := chip("Boron")

	r1 := RTG("Hydrogen")
	r2 := RTG("Boron")

	var s *state

	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 1}, []int{1, 1}, 1}
	checkIsValid(t, s, true)

	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 1}, []int{2, 2}, 1}
	checkIsValid(t, s, true)

	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 2}, []int{3, 4}, 1}
	checkIsValid(t, s, true)

	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 3}, []int{2, 4}, 1}
	checkIsValid(t, s, true)

	//Boron chip is on different floor than it's RTG
	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 1}, []int{1, 2}, 1}
	checkIsValid(t, s, false)

	//Hydrogen chip is on different floor than it's RTG
	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 1}, []int{2, 1}, 1}
	checkIsValid(t, s, false)

	//We're not on the same floor as everything else
	s = &state{[]chip{c1, c2}, []RTG{r1, r2}, []int{1, 3}, []int{2, 4}, 5}
	checkIsValid(t, s, true)

}
