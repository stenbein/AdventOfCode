// Advent of Code 2016 - Day 11
//
// 		The idea here is to generate a list of possible transitions
// to new states for each state and to record at each transition 
// how much time it took to get there. My biggest issue was figuring
// how to generate the state keys.
//
// 		Issues: No major issues. But left this one behind while I 
// did some other things and fell off the wagon.
//
//

package main

import (
	"fmt"
	"strconv"
)

type chip string
type RTG string

//index of the chip or RTG is the id
//in sync with the floor values
type state struct {
	chips        []chip
	RTGs         []RTG
	chipFloor    []int
	RTGFloor     []int
	currentFloor int
}

//check that no chip present without it's own RTG on a floor with other RTGs
func (s *state) isValid() bool {

	//checkElevator := false

	for i, _ := range s.chips {

		//check that you're on a floor with at least one chip or RTG
		/*if s.chipFloor[i] == s.currentFloor ||
		  s.RTGFloor[i] == s.currentFloor {
		      checkElevator = true
		  }*/

		//chip and RTG not on same floor
		if s.chipFloor[i] != s.RTGFloor[i] {

			//chip is unshielded, check to see if any other RTG is on this floor
			for j, _ := range s.RTGs {
				if s.chipFloor[i] == s.RTGFloor[j] {
					return false
				}
			}

		}

	}
	return true
}

//represent pairs of chips as len(chip) to avoid collision
func (s *state) key() [5]int {

	//space for the count on each floor plus the current floor
	floors := [5]int{}
	//fmt.Println("-----")
	for i, _ := range s.chips {
		//pair is together
		//fmt.Println(s.chipFloor[i])
		//fmt.Println(s.RTGFloor[i])
		if s.chipFloor[i] == s.RTGFloor[i] {
			floors[s.chipFloor[i]-1] += len(s.chips)
		} else {

			floors[s.chipFloor[i]-1] += 1
			floors[s.RTGFloor[i]-1] += 1

		}

	}

	//current floor
	floors[4] = s.currentFloor

	return floors

}

func (s state) copyState() state {

	//chips           []chip
	//RTGs            []RTG
	//chipFloor       []int
	//RTGFloor        []int
	//currentFloor    int

	l := len(s.chips)

	chips := make([]chip, l)
	RTGs := make([]RTG, l)
	chipFloor := make([]int, l)
	RTGFloor := make([]int, l)

	copy(chips, s.chips)
	copy(RTGs, s.RTGs)
	copy(chipFloor, s.chipFloor)
	copy(RTGFloor, s.RTGFloor)

	return state{chips, RTGs, chipFloor, RTGFloor, s.currentFloor}

}

//this code smells, todo, rework this idea to fold the logic into
//the state/floor structs
func (s *state) siblingChips(first, second int) []state {

	var cs1, cs2 state

	cs1 = s.copyState()

	switch s.currentFloor {
	case 1:
		cs1.chipFloor[first] = 2
		cs1.chipFloor[second] = 2
		cs1.currentFloor = 2
		return []state{cs1}
	case 4:
		cs1.chipFloor[first] = 3
		cs1.chipFloor[second] = 3
		cs1.currentFloor = 3
		return []state{cs1}
	default:
		cs2 = s.copyState()
		cs1.chipFloor[first] = s.currentFloor - 1
		cs1.chipFloor[second] = s.currentFloor - 1
		cs1.currentFloor = s.currentFloor - 1
		cs2.chipFloor[first] = s.currentFloor + 1
		cs2.chipFloor[second] = s.currentFloor + 1
		cs2.currentFloor = s.currentFloor + 1
		return []state{cs1, cs2}
	}

}

func (s *state) siblingRTGs(first, second int) []state {

	var cs1, cs2 state

	cs1 = s.copyState()

	switch s.currentFloor {
	case 1:
		cs1.RTGFloor[first] = 2
		cs1.RTGFloor[second] = 2
		cs1.currentFloor = 2
		return []state{cs1}
	case 4:
		cs1.RTGFloor[first] = 3
		cs1.RTGFloor[second] = 3
		cs1.currentFloor = 3
		return []state{cs1}
	default:
		cs2 = s.copyState()
		cs1.RTGFloor[first] = s.currentFloor - 1
		cs1.RTGFloor[second] = s.currentFloor - 1
		cs1.currentFloor = s.currentFloor - 1
		cs2.RTGFloor[first] = s.currentFloor + 1
		cs2.RTGFloor[second] = s.currentFloor + 1
		cs2.currentFloor = s.currentFloor + 1
		return []state{cs1, cs2}
	}

}

func (s *state) siblings(chip, RTG int) []state {

	var cs1, cs2 state

	cs1 = s.copyState()

	switch s.currentFloor {
	case 1:
		cs1.chipFloor[chip] = 2
		cs1.RTGFloor[RTG] = 2
		cs1.currentFloor = 2
		return []state{cs1}
	case 4:
		cs1.chipFloor[chip] = 3
		cs1.RTGFloor[RTG] = 3
		cs1.currentFloor = 3
		return []state{cs1}
	default:
		cs2 = s.copyState()
		cs1.chipFloor[chip] = s.currentFloor - 1
		cs1.RTGFloor[RTG] = s.currentFloor - 1
		cs1.currentFloor = s.currentFloor - 1
		cs2.chipFloor[chip] = s.currentFloor + 1
		cs2.RTGFloor[RTG] = s.currentFloor + 1
		cs2.currentFloor = s.currentFloor + 1
		return []state{cs1, cs2}
	}

}

//for a given state, return the possible next state steps
//valid steps will be all possible valid actions for the chips
//and RTGs on the current floor
func (s *state) validSteps() []state {

	states := []state{}
	validStates := []state{}

	//get ids of all chips and RTGs on this floor
	l := len(s.chips)
	for i := 0; i < l; i++ {
		if s.chipFloor[i] == s.currentFloor {
			for j := i; j < l; j++ {
				if s.chipFloor[j] == s.currentFloor {
					//get siblings for chips only
					states = append(states, s.siblingChips(i, j)...)
				}
			}
			for j := 1; j < l; j++ {
				if s.RTGFloor[j] == s.currentFloor {
					//get siblings for chips and RTGs
					states = append(states, s.siblings(i, j)...)
				}
			}
		}
		if s.RTGFloor[i] == s.currentFloor {
			for j := i; j < l; j++ {
				if s.RTGFloor[j] == s.currentFloor {
					//siblings for RTGs only
					states = append(states, s.siblingRTGs(i, j)...)
				}
			}
		}
	}

	for i, _ := range states {
		if states[i].isValid() {
			//fmt.Println(states[i])
			validStates = append(validStates, states[i])
		}
	}

	return validStates

}

//given a state, an a map of states already visited. Check to see if
//the key for that state is already in the map. If so, check to see if
//we got there faster this way. Else set the current state to the
//current step count. This is functionally a depth first search of the
//state space which terminates when the minimal step count for each
//state is reached.
func itter(s state, visited map[[5]int]int, steps int) map[[5]int]int {

	k := s.key()
	_, ok := visited[k]

	if !ok {

		//we haven't seen this state before
		//record it and itterate
		visited[k] = steps
		for _, next := range s.validSteps() {
			itter(next, visited, steps+1)
		}

	} else {

		//else, we've seen the state before, if we did 'better'
		//by fewer steps, update the steps to state and itterate
		//else nip in the bud and return
		if visited[k] > steps {

			visited[k] = steps
			for _, next := range s.validSteps() {
				itter(next, visited, steps+1)
			}

		}

	}

	return visited

}

func process(s state, goal state) int {

	states := itter(s, make(map[[5]int]int), 0)

	return states[goal.key()]

}

/*The first floor contains a polonium generator, a thulium generator, a
    thulium-compatible microchip, a promethium generator, a ruthenium
    generator, a ruthenium-compatible microchip, a cobalt generator,
    and a cobalt-compatible microchip.
The second floor contains a polonium-compatible microchip and a
    promethium-compatible microchip.
The third floor contains nothing relevant.
The fourth floor contains nothing relevant.
*/
//part 2 add elerium and dilithium
func final_1() state {

	c := []chip{"polonium", "thulium", "promethium", "ruthenium", "cobalt"}
	r := []RTG{"polonium", "thulium", "promethium", "ruthenium", "cobalt"}
	cFloor := []int{4, 4, 4, 4, 4}
	rFloor := []int{4, 4, 4, 4, 4}
	endFloor := 4

	return state{c, r, cFloor, rFloor, endFloor}

}

func final_2() state {

	c := []chip{"polonium", "thulium", "promethium", "ruthenium", "cobalt", "elerium", "dilithium"}
	r := []RTG{"polonium", "thulium", "promethium", "ruthenium", "cobalt", "elerium", "dilithium"}
	cFloor := []int{4, 4, 4, 4, 4, 4, 4}
	rFloor := []int{4, 4, 4, 4, 4, 4, 4}
	endFloor := 4

	return state{c, r, cFloor, rFloor, endFloor}

}

func part_one() string {

	c := []chip{"polonium", "thulium", "promethium", "ruthenium", "cobalt"}
	r := []RTG{"polonium", "thulium", "promethium", "ruthenium", "cobalt"}
	cFloor := []int{2, 1, 2, 1, 1}
	rFloor := []int{1, 1, 1, 1, 1}

	initial := state{c, r, cFloor, rFloor, 1}

	dist := process(initial, final_1())

	return strconv.Itoa(dist)

}

func part_two() string {

	c := []chip{"polonium", "thulium", "promethium", "ruthenium", "cobalt", "elerium", "dilithium"}
	r := []RTG{"polonium", "thulium", "promethium", "ruthenium", "cobalt", "elerium", "dilithium"}
	cFloor := []int{2, 1, 2, 1, 1, 1, 1}
	rFloor := []int{1, 1, 1, 1, 1, 1, 1}

	initial := state{c, r, cFloor, rFloor, 1}

	dist := process(initial, final_2())

	return strconv.Itoa(dist)

}

func main() {

	//input is a set state
	fmt.Println("Problem 1: " + part_one())
	fmt.Println("Problem 2: " + part_two())

}
