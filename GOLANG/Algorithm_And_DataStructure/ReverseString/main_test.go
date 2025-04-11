package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {

	subtests := []struct {
		initialString  string
		reversedString string
	}{
		{

			"hello",
			"olleh",
		},
		{
			"Mahmoud",
			"duomhaM",
		},
	}

	for _, st := range subtests {

		result := ReverseString(st.initialString)
		if result != st.reversedString {
			t.Errorf("wanted %v, got %v", st.reversedString, result)
		}
	}

}
