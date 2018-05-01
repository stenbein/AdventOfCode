// Advent of Code 2016 - Day 10
//
// 		Approach taken was to keep track of the bots in a
// closure and print the id of the bot which handled the
// chips we were interested in to a global variable. That
// feels wrong, maybe I can redo it later.
//
// 		Issues: Tried this one a few ways. First attempt
// was what ultimately worked but I got distracted
// assuming that I needed to have the bots running
// concurrently. I had a set of code which used channels
// to pass values between bots and it got the first answer
// correct but I had issues with the second one and left
// it at that. So far day 10 has given me the most issues.
// By this point I have day 1-21 complete save 10, 11, and
// 14. 10 was a bit of a sore spot, though looking back
// nothing appears to be wrong with it. Just a mental
// block I suppose.
//

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type chip int

type Receiver interface {
	receive(chip)
}

type bot struct {
	id        int
	inventory [2]chip
}

type bin struct {
	id        int
	inventory []chip
}

var hiReceiver map[int]Receiver
var lowReceiver map[int]Receiver
var botID int // store the id of the bot for problem 1

//this is here to sort the instructions so we have value last
type valuesLast []string

func (v valuesLast) Len() int {
	return len(v)
}
func (v valuesLast) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v valuesLast) Less(i, j int) bool {
	return v[i] < v[j]
}

func (b *bot) receive(c chip) {

	if b.inventory[0] == 0 {
		b.inventory[0] = c
	} else {
		b.inventory[1] = c
	}
	b.process()

}

func (bi *bin) receive(c chip) {
	bi.inventory = append(bi.inventory, c)
}

func (b *bot) check() {

	if (b.inventory[0] == 61 && b.inventory[1] == 17) ||
		(b.inventory[0] == 17 && b.inventory[1] == 61) {
		botID = b.id
	}

}

func (b *bot) process() {

	var hi chip
	var low chip

	if b.inventory[0] != 0 && b.inventory[1] != 0 {

		b.check() //determines value of part 1

		if b.inventory[0] > b.inventory[1] {
			hi = b.inventory[0]
			low = b.inventory[1]
		} else {
			hi = b.inventory[1]
			low = b.inventory[0]
		}

		//give hi to hi target
		hiReceiver[b.id].receive(hi)

		//give low to low target
		lowReceiver[b.id].receive(low)

		//zero out inventory
		b.inventory = [2]chip{0, 0}
	}

}

func bots() func(int) *bot {

	store := make(map[int]*bot)

	return func(i int) *bot {

		b, ok := store[i]
		if ok == false {
			b = &bot{i, [2]chip{0, 0}}
			store[i] = b
			return b
		}
		return b
	}

}

func bins() func(int) *bin {

	store := make(map[int]*bin)

	return func(i int) *bin {

		bi, ok := store[i]
		if ok == false {
			bi = &bin{i, []chip{}}
			store[i] = bi
			return bi
		}
		return bi
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part_one(input string) string {

	getBot := bots()
	getBin := bins()
	hiReceiver = make(map[int]Receiver)
	lowReceiver = make(map[int]Receiver)

	instructions := strings.Split(input, "\n")
	sort.Sort(valuesLast(instructions))

loop:
	for _, instruction := range instructions {

		if instruction == "" {
			continue loop
		}

		tokens := strings.Split(instruction, " ")
		switch tokens[0] {
		case "value":
			v, _ := strconv.Atoi(tokens[1])
			id, _ := strconv.Atoi(tokens[5])
			if tokens[4] == "bot" {
				getBot(id).receive(chip(v))
			} else {
				getBin(id).receive(chip(v))
			}

		default:

			id, _ := strconv.Atoi(tokens[1])
			hiId, _ := strconv.Atoi(tokens[11])
			lowId, _ := strconv.Atoi(tokens[6])

			if tokens[10] == "bot" {
				hiReceiver[id] = getBot(hiId)
			} else {
				hiReceiver[id] = getBin(hiId)
			}
			if tokens[5] == "bot" {
				lowReceiver[id] = getBot(lowId)
			} else {
				lowReceiver[id] = getBin(lowId)
			}
		}

	}

	return strconv.Itoa(botID) //the use of a global here doesn't feel right...

}

func part_two(input string) string {

	getBot := bots()
	getBin := bins()
	hiReceiver = make(map[int]Receiver)
	lowReceiver = make(map[int]Receiver)

	instructions := strings.Split(input, "\n")
	sort.Sort(valuesLast(instructions))

loop:
	for _, instruction := range instructions {

		if instruction == "" {
			continue loop
		}

		tokens := strings.Split(instruction, " ")
		switch tokens[0] {
		case "value":
			v, _ := strconv.Atoi(tokens[1])
			id, _ := strconv.Atoi(tokens[5])
			if tokens[4] == "bot" {
				getBot(id).receive(chip(v))
			} else {
				getBin(id).receive(chip(v))
			}

		default:

			id, _ := strconv.Atoi(tokens[1])
			hiId, _ := strconv.Atoi(tokens[11])
			lowId, _ := strconv.Atoi(tokens[6])

			if tokens[10] == "bot" {
				hiReceiver[id] = getBot(hiId)
			} else {
				hiReceiver[id] = getBin(hiId)
			}
			if tokens[5] == "bot" {
				lowReceiver[id] = getBot(lowId)
			} else {
				lowReceiver[id] = getBin(lowId)
			}
		}

	}

	output := int(getBin(0).inventory[0] * getBin(1).inventory[0] * getBin(2).inventory[0])

	return strconv.Itoa(output)

}

func main() {

	input, err := ioutil.ReadFile("./2016_10.txt")
	check(err)

	fmt.Println("Problem 1: " + part_one(string(input)))
	fmt.Println("Problem 2: " + part_two(string(input)))

}
