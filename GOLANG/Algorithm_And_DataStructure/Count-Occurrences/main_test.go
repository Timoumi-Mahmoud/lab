package main

import (
	"reflect"
	"testing"
)

func TestCountOccurrences(t *testing.T) {

	subtests := []struct {
		stringInput string
		occurrences map[string]int
	}{
		{

			"amma",
			map[string]int{"a": 2, "m": 2},
		},
		{
			"mahmoud",
			map[string]int{"m": 2, "a": 1, "h": 1, "o": 1, "u": 1, "d": 1},
		},
	}

	for _, st := range subtests {
		if s := CountOccurrences(st.stringInput); reflect.DeepEqual(s, st.occurrences) == false {
			t.Errorf("wanted %v, got %v", st.occurrences, s)
		}
	}

}
