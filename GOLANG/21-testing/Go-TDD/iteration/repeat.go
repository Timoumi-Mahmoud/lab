package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		//		repeated += character
		repeated.WriteString(character)
	}
	return repeated.String()
}

/*
Strings in go are immutable, meaning every concatenation involves copying memory to accommodate the new string.
this impacts performance, particularly during heavy string concatenation.
strings.Builder type: which minimizes memory copying, it implements a writeString method which can use to concatenate strings.

*/
