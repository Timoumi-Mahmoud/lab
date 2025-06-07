package main

import (
	"testing"
)

func TestIsPlandrom(t *testing.T) {

	subtests := []struct {
		stringInput string
		isplandrome bool
	}{
		{

			"amma",
			true,
		},
		{
			"Mahmoud",
			false,
		},
	}

	for _, st := range subtests {

		result := IsPlandrom(st.stringInput)
		if result != st.isplandrome {
			t.Errorf("wanted %v, got %v", st.isplandrome, result)
		}
	}

}
