package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {

	subtests := []struct {
		array  []int
		result bool
	}{
		{

			[]int{1, 2, 4, 5},
			false,
		},
		{

			[]int{1, 2, 1, 6},
			true,
		},
		{

			[]int{11, 2, 474, 5, 58, 447, 474},
			true,
		},
	}

	for _, st := range subtests {

		s := containsDup(st.array)
		if s != st.result {
			t.Errorf("wanted %v, got %v", st.array, s)
		}
	}

}
