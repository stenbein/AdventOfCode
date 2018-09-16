// Advent of Code 2016 - Day 15
//
// 		Approach taken was to implement a state machine
// and itterate though passable states. I realized at
// end that only certain states would be possible so you
// could simply jump though states that aren't going to
// clear. TODO, revisit and modify to find constant time
// approach.
//
// 		Issues: None
//
//

package main

import (
	"fmt"
)

type Disk struct {
	start    int
	position int // might not be zero, taken care of during init
	period   int // steps to repeat

}

func (d *Disk) tick(t int) {
	d.position = (d.position + t) % d.period
}

func (d *Disk) isPassable() bool {
	return d.position == 0
}

func (d *Disk) reset() {
	d.position = d.start
}

func NewDisk(start, period int) Disk {

	return Disk{start, start, period}

}

type diskDeathMachine struct {
	ticker int
	disks  []Disk
}

func (dm *diskDeathMachine) tick(t int) {
	for i, _ := range dm.disks {
		dm.disks[i].tick(t)
	}
}

func (dm *diskDeathMachine) reset() {
	for i, _ := range dm.disks {
		dm.disks[i].reset()
	}
}

func (dm *diskDeathMachine) solve(delay int) bool {

	dm.tick(delay)
	for i, _ := range dm.disks {
		dm.tick(1)
		if !dm.disks[i].isPassable() {
			return false
		}
	}

	return true
}

/*
Disc #1 has 17 positions; at time=0, it is at position 15.
Disc #2 has 3 positions; at time=0, it is at position 2.
Disc #3 has 19 positions; at time=0, it is at position 4.
Disc #4 has 13 positions; at time=0, it is at position 2.
Disc #5 has 7 positions; at time=0, it is at position 2.
Disc #6 has 5 positions; at time=0, it is at position 0.
*/

//Disc #1 has 5 positions; at time=0, it is at position 4.
//Disc #2 has 2 positions; at time=0, it is at position 1.
func main() {

	dm := diskDeathMachine{0, []Disk{NewDisk(15, 17),
		NewDisk(2, 3),
		NewDisk(4, 19),
		NewDisk(2, 13),
		NewDisk(2, 7),
		NewDisk(0, 5),
		NewDisk(0, 11)},
	}

	delay := 1
	for {

		if !dm.solve(delay) {
			dm.reset()
			delay += 17
		} else {
			break
		}

	}

	fmt.Println(delay)

}
