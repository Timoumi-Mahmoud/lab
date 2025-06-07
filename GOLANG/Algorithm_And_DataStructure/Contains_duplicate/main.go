package main

import (
	"fmt"
)

//	func containsDup(array []int) bool {
//		temp_map := map[int]bool{}
//		for _, v := range array {
//			temp_map[v] = true
//		}
//		if len(temp_map) != len(array) {
//			return true
//		}
//
//		return false
//	}
func containsDup(array []int) bool {
	temp_map := map[int]bool{}
	for _, v := range array {
		temp_map[v] = true
	}
	return len(temp_map) != len(array)

}

func main() {

	fmt.Println("Rolling ...")
	inputArray := []int{1, 2, 3, 4, 1, 5}
	fmt.Println("=> after ", containsDup(inputArray))

}

/*
- Given an array of int , return true if it contains duplicate else return false.
*/
