package main

import (
	"reflect"
	"testing"
)

func TestBubleSort(t *testing.T) {

	subtests := []struct {
		unsortedArray []int
		sortedArray   []int
	}{
		{

			[]int{5, 4, 8, 1, 6, 2, 7, 3},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{77, 20, 31, 45, 55, 83, 10, 82},
			[]int{10, 20, 31, 45, 55, 77, 82, 83},
		},
	}

	for _, st := range subtests {
		if s := BubleSort(st.unsortedArray); reflect.DeepEqual(s, st.sortedArray) == false {
			t.Errorf("wanted %v, got %v", st.sortedArray, s)
		}
	}

}
