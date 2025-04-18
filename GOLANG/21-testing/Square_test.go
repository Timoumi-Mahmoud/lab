package main

import "testing"

func TestSquare(t *testing.T) {
	tcases := []struct {
		description string
		input       int
		expected    int
	}{
		{
			description: "Correctly squares a number 4",
			input:       4,
			expected:    16,
		},
		{

			description: "Correctly squares a number 3",
			input:       3,
			expected:    9,
		},
		{

			description: "Correctly squares a number 4",
			input:       10,
			expected:    100,
		},
		{

			description: "Correctly squares a number 10",
			input:       100,
			expected:    10000,
		},
		{

			description: "Correctly squares a negative number  -10",
			input:       100,
			expected:    10000,
		},
	}
	result := Square(4)

	for _, ts := range tcases {
		t.Run(ts.description, func(t *testing.T) {
			if result = Square(ts.input); result != ts.expected {
				t.Errorf("Want %d , got %d \n", ts.expected, result)
			}
		})
	}

	/*	for _, st := range tcases {
			result = Square(st.input)
			if result != st.expected {
				t.Errorf("Want %d , got %d \n", st.expected, result)
			}
		}
	*/

}
