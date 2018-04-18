package main

import (
	//"fmt"
	"testing"
)

//https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go

//simple method to print out the grid
/*func (g *grid) rep() {
	for _, row := range g.tiles {
		fmt.Println(row)
	}
}*/

//check the value of the boolean slice for the doors
func ensureDoors(t *testing.T, r *room, expectValue [4]bool) {

	if v := r.doors(); v != expectValue {
		t.Errorf("r.door() = %d, want %d", v, expectValue)
	}

}

func ensureLocks(t *testing.T, r *room, expectValue [4]bool) {

	if v := r.unlocked(); v != expectValue {
		t.Errorf("r.door() = %d, want %d", v, expectValue)
	}

}

func TestDoors(t *testing.T) {

	var r *room
	r = &room{0, "a"}
	ensureDoors(t, r, [4]bool{false, true, false, true})

	r = &room{1, "a"}
	ensureDoors(t, r, [4]bool{false, true, true, true})

	r = &room{2, "a"}
	ensureDoors(t, r, [4]bool{false, true, true, true})

	r = &room{3, "a"}
	ensureDoors(t, r, [4]bool{false, true, true, false})

	r = &room{4, "a"}
	ensureDoors(t, r, [4]bool{true, true, false, true})

	r = &room{8, "a"}
	ensureDoors(t, r, [4]bool{true, true, false, true})

	r = &room{10, "a"}
	ensureDoors(t, r, [4]bool{true, true, true, true})

	r = &room{11, "a"}
	ensureDoors(t, r, [4]bool{true, true, true, false})

	r = &room{12, "a"}
	ensureDoors(t, r, [4]bool{true, false, false, true})

	r = &room{13, "a"}
	ensureDoors(t, r, [4]bool{true, false, true, true})

	r = &room{14, "a"}
	ensureDoors(t, r, [4]bool{true, false, true, true})

	r = &room{15, "a"}
	ensureDoors(t, r, [4]bool{true, false, true, false})

}

// test case ensureLocks: "ihgpwlahDD" showed issue was
// off by one error in unlocked() function, checking > 11
// hex b instead of > 10, hex a
func TestUnlocked(t *testing.T) {

	var r *room
	r = &room{0, "hijkl"}
	ensureLocks(t, r, [4]bool{false, true, false, false})

	r = &room{4, "hijklD"}
	ensureLocks(t, r, [4]bool{true, false, false, true})

	r = &room{5, "hijklDR"}
	ensureLocks(t, r, [4]bool{false, false, false, false})

	r = &room{0, "hijklDU"}
	ensureLocks(t, r, [4]bool{false, false, false, true})

	r = &room{1, "hijklDUR"}
	ensureLocks(t, r, [4]bool{false, false, false, false})

	r = &room{8, "ihgpwlahDD"}
	ensureLocks(t, r, [4]bool{false, false, false, true})
}
