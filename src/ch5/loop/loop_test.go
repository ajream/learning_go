package loop_test

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 10
	for i := 1; i < n; i++ {
		fmt.Print(i)
	}

	for n < 20 {
		fmt.Print(n)
		n++
	}
}
