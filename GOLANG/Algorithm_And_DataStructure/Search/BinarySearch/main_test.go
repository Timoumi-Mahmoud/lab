package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	subtests := []struct {
		array  []int
		target int
		result int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			2,
			1,
		},
		{
			[]int{10, 20, 31, 45, 55, 66, 77, 82},
			77,
			6,
		},
		{
			[]int{10, 20, 31, 45, 55, 66, 77, 82},
			2,
			-1,
		},
	}

	for _, st := range subtests {
		if s := BinarySearch(st.array, st.target); s != st.result {
			t.Errorf("wanted %d, got %d", st.result, s)
		}
	}

}
