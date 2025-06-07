package main

import (
	"fmt"
)

func IsPlandrom(input string) bool {
	l, r := 0, len(input)-1
	for l != r {
		if input[l] != input[r] {
			return false
		}
		l++
		r--
		if r == 0 {
			break
		}
	}
	return true
}

func main() {

	stringInput := "akma"
	fmt.Println("=> ", IsPlandrom(stringInput))

}

/*
- check if a string is plandrome or not

*/
