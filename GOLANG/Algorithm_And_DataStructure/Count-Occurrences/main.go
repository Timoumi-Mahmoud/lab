package main

import (
	"fmt"
)

func CountOccurrences(input string) map[string]int {
	m := map[string]int{}
	for _, v := range input {
		m[string(v)]++
	}
	return m
}

func main() {

	stringInput := "mahmoud timoumi"
	fmt.Println("=> ", CountOccurrences(stringInput))
	m := map[string]int{}
	m["hh"] = 55
	fmt.Println(m)

}

/*
- count the occurrences of each letter in a string

*/
