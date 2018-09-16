package day19

import (
	"fmt"
	"testing"
)

//https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go

//simple method to print out the grid
func (e *elf) rep() {

	fmt.Println("Pass")

}

//check status of presents
func ensureValue(t *testing.T, e *elf, expectValue int) {

	if v := e.presents; v != expectValue {
		t.Errorf("e.presents = %d, want %d", v, expectValue)
	}

}

//check value of int (for ids)
func ensureIndex(t *testing.T, i int, expectValue int) {

	if v := i; v != expectValue {
		t.Errorf("Index = %d, want %d", v, expectValue)
	}

}

func ensureLen(t *testing.T, es elves, expectValue int) {

	if v := len(es); v != expectValue {
		t.Errorf("Len(elves) = %d, want %d", v, expectValue)
	}

}

func expectPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered: ", r)
		}
	}()
	f()
	t.Errorf("The code did not panic")
}

func TestTake(t *testing.T) {

	e0 := elf{1, 1}
	e1 := elf{2, 1}

	e0.take(&e1)

	ensureValue(t, &e0, 2)
	ensureValue(t, &e1, 0)

	e2 := elf{3, 0}
	e2.take(&e0)

	ensureValue(t, &e0, 0)
	ensureValue(t, &e2, 2)

	e3 := elf{4, 1}
	e3.take(&e2)

	ensureValue(t, &e2, 0)
	ensureValue(t, &e3, 3)

}

func TestSearch(t *testing.T) {

	e0 := elf{1, 1}
	e1 := elf{2, 1}
	e2 := elf{3, 1}
	e3 := elf{4, 1}

	es := elves{e0, e1, e2, e3}

	ensureIndex(t, es.search(1), 0)
	ensureIndex(t, es.search(2), 1)
	ensureIndex(t, es.search(3), 2)
	ensureIndex(t, es.search(4), 3)

	f := func() {
		es.search(-1)
	}
	expectPanic(t, f)

	f = func() {
		es.search(5)
	}
	expectPanic(t, f)

}

func TestReduce(t *testing.T) {

	e0 := elf{1, 1}
	e1 := elf{2, 1}
	e2 := elf{3, 1}
	e3 := elf{4, 1}

	es := elves{e0, e1, e2, e3}

	ensureLen(t, es, 4)

	es = es.reduce()

	ensureLen(t, es, 4)

	e0 = elf{1, 1}
	e1 = elf{2, 0}
	e2 = elf{3, 1}
	e3 = elf{4, 0}

	es = elves{e0, e1, e2, e3}

	ensureLen(t, es, 4)

	es = es.reduce()

	ensureLen(t, es, 2)

	f := func() {
		es.search(2)
	}
	expectPanic(t, f)

	f = func() {
		es.search(4)
	}
	expectPanic(t, f)

}

func TestExchange(t *testing.T) {

	e0 := elf{1, 1}
	e1 := elf{2, 1}
	e2 := elf{3, 1}
	e3 := elf{4, 1}

	es := elves{e0, e1, e2, e3}

	es.exchange(0, 0)
	ensureValue(t, &es[0], 2)
	ensureValue(t, &es[2], 0)

	es.exchange(1, 0)
	ensureValue(t, &es[1], 2)
	ensureValue(t, &es[3], 0)

}

func TestAll(t *testing.T) {

	es := newElves(5)
	ensureIndex(t, process(es, 1), 2)

	es = newElves(8)
	ensureIndex(t, process(es, 1), 7)

	es = newElves(12)
	ensureIndex(t, process(es, 1), 3)

	es = newElves(3012210)
	ensureIndex(t, process(es, 1), -1)

}
