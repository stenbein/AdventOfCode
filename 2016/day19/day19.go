// Advent of Code 2016 - Day 19
//
// 		Approach taken was first to model the elves as
// as a linked list. This proved too costly during the
// second part where the input was 3 million. Modified
// the code to itterate a slice and update it as we 
// progressed.
//
// 		Issues: Lots. First approach worked without 
// issue for part 1, but part 2 was too slow and failed
// to run with the list traversial I had written. 
// Choosing between working with the existing approach 
// and rewriting I decided to rewrite to handle the 
// exchanges better. The result is not what I consider
// a good bit of code and the approach doesn't work for
// part 1 and 2.
//
// TODO: refactor and bring back part 1.
//
//

package day19

import (
	"fmt"
	//"strconv"
	"errors"
)


var elfCount int

type elf struct {

	id 			int
	presents 	int

}

type elves []elf

func (e *elf) take(o *elf) {

	e.presents += o.presents
	(*o).presents = 0 //zero it out for boundary testing

}

//binary search on the elves to find the index if the
//id passed as param or panics if it can't find it
//elves are expected to be in order numerically by id
func (es elves) search(id int) int {

	l := 0
	h := len(es)

	for l != h {

		t := (l + h) / 2
		if es[t].id == id {
			return t
		} else if es[t].id > id { //guessed too high
			h = t
		} else { //guessed too low
			l = t + 1
		}

	}

	panic("Can't find id in array, you goofed.")

}

//loops the array starting with the id of the elf given
//returns the id of the elf which last was update
func (es elves) exchange(index int, offset int) (error){

	//steal from elf 1/2 across the array
	l := len(es)
	t := (index + offset + (l / 2)) % l

	if es[t].presents != 0 {
		es[index].take(&es[t])
		return nil
	} else {
		return errors.New("Elf has no presents")
	}

}

func (es elves) reduce() elves {

	count := 0
	for _, e := range es {
		if e.presents > 0 {count++}
	}

	out := make(elves, count)

	i := 0
	for _, e := range es {
		if e.presents > 0 {
			out[i] = e
			i++
		}
	}
	return out

}

func process(es elves, id int) int {

	l := len(es)
	offset := 0
	offsettoggle := false

	index := es.search(id)
	for len(es) > 1 {

		for {

			//fmt.Println(es)
			
			if es[index % l].presents != 0 { //was already processed
			
				err := es.exchange(index, offset)
				if err == nil {
					id = es[index % l].id //store id of last successful exchange
					index = (index+1) % l //incremend index and wrap
					if offsettoggle {
						offset++
						//fmt.Println("Offset: " + strconv.Itoa(offset))
					}
					offsettoggle = !offsettoggle
					
				
				} else {
					//fmt.Println("Exchange failed for id: " + strconv.Itoa(id))
					es = es.reduce()
					l = len(es) //update len to new size after reduce
					index = es.search(id) //recover index of the current id
					offset = 0
					break
				}

			} else {
				//fmt.Println("Id already processed: " + strconv.Itoa(id))
				es = es.reduce()
				l = len(es) //update len to new size after reduce
				index = (es.search(id) + 1) % l //recover index of the last success
				offset = 0
				break
			}
		}

	}

	return es[0].id

}

func newElves(size int) elves {

	es := make(elves, size)
	for i, _ := range es {
		es[i].id = i+1
		es[i].presents = 1
	}
	return es
}




/*func part_one(input int) string {

	es := newElves(input)
	return strconv.Itoa(process(es, 1))

}*/

func part_two(input string) string {

	es := newElves(input)
	return strconv.Itoa(process(es, 1))

}

func main() {

	input := 3012210

	//fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(input))

}














