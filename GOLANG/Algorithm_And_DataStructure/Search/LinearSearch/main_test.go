package main

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	testCase := struct {
		f string
	}{
		"mdkfjdk",
	}

	fmt.Println(testCase)
	subtests := []struct {
		array  []int
		target int
		result int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			5,
			4,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			555,
			-1,
		},
	}

	for _, st := range subtests {
		if s := LinearSearch(st.array, st.target); s != st.result {
			t.Errorf("wanted %d, got %d", st.result, s)
		}
	}

}
