package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{5, 6, 7, 8}

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{5, 6, 7, 8}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for i, e := range arr2 {
		t.Log(i, e)
	}

	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [4]int{1, 2, 3, 4}
	arr_sec := arr[1:]

	t.Log(arr_sec)
}
