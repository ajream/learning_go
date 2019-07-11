package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	// fmt.Print(Monday, Tuesday)
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	fmt.Print(Readable, Writable, Executable)
	// fmt.Print(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
