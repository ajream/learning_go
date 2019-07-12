package operator_test

import (
	"fmt"
	"log"
	"testing"
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	fmt.Println(a == b)
	log.Print(a == b)
	fmt.Println(a == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	fmt.Println(Readable, Writable, Executable)
	fmt.Println(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	fmt.Println(1&^0, 1&^1, 2&^0, 12&^0)
	// t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
