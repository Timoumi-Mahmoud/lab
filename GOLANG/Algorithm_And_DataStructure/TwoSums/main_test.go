package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {

	subtests := []struct {
		arr    []int
		target int
		output []int
	}{
		{
			[]int{1, 5, 78, 6, 9, 7, 1},
			2,
			[]int{0, 6},
		},
	}

	for _, st := range subtests {

		result := TwoSums(st.arr, st.target)
		if reflect.DeepEqual(result, st.output) == false {
			t.Errorf("wanted %v, got %v", st.output, result)
		}
	}

}
