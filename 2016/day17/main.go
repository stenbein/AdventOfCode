// Advent of Code 2016 - Day 17
//
// 		Approach taken was
//
// 		Issues:
//
//

package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

/*Only the first four characters of the hash are used;
they represent, respectively, the doors up, down, left,
and right from your current position. Any b, c, d, e,
or f means that the corresponding door is open; any
other character (any number or a) means that the
corresponding door is closed and locked.

#########
#S| | | #
#-#-#-#-#
# | | | #
#-#-#-#-#
# | | | #
#-#-#-#-#
# | | |
####### V

*/

type room struct {
	id   int
	path string
}

//up, down, left, and right
func (r *room) doors() [4]bool {

	return [4]bool{r.id > 3, r.id < 12, (r.id % 4) != 0, ((r.id + 1) % 4) != 0}

}

func (r *room) unlocked() [4]bool {

	data := []byte(r.path)
	bytes := md5.Sum(data)

	doors := r.doors()
	unlocked := [4]bool{}

	//up
	unlocked[0] = doors[0] && ((uint(bytes[0]>>4) & uint(15)) > 10)
	//down
	unlocked[1] = doors[1] && ((uint(bytes[0]) & uint(15)) > 10)
	//left
	unlocked[2] = doors[2] && ((uint(bytes[1]>>4) & uint(15)) > 10)
	//right
	unlocked[3] = doors[3] && ((uint(bytes[1]) & uint(15)) > 10)

	return unlocked

}

func (r *room) nextRooms() [4]room {

	rooms := [4]room{}
	paths := r.unlocked()

	if paths[0] == true {
		rooms[0] = room{r.id - 4, r.path + "U"}
	}
	if paths[1] == true {
		rooms[1] = room{r.id + 4, r.path + "D"}
	}
	if paths[2] == true {
		rooms[2] = room{r.id - 1, r.path + "L"}
	}
	if paths[3] == true {
		rooms[3] = room{r.id + 1, r.path + "R"}
	}

	return rooms
}

func part_one(input string) string {

	var current room
	var queue chan room

	queue = make(chan room, 500)
	queue <- room{0, input}

	for current = range queue {

		//check if that path leads us to the goal
		if current.id == 15 {
			//if so, break and output path
			break
		} else {

			for _, p := range current.nextRooms() {
				if p.path != "" {
					queue <- p
				}
			}

		}

	}

	return current.path

}

func part_two(input string) string {

	longest := ""

	var queue chan room

	queue = make(chan room, 500)
	queue <- room{0, input}

loop:
	for {

		select {

		case current := <-queue:

			//check if that path leads us to the goal
			if current.id == 15 {
				//check against current longest
				if len(current.path) > len(longest) {
					longest = current.path
				}
			} else {

				for _, p := range current.nextRooms() {
					if p.path != "" {
						queue <- p
					}
				}

			}

		case <-time.After(1 * time.Second):
			break loop
		}

	}

	return strconv.Itoa(len(longest) - len(input))

}

func main() {

	input := "njfxhljp"

	fmt.Println("Problem 1: " + part_one(input))
	fmt.Println("Problem 2: " + part_two(input))

}
