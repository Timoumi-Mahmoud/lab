package main

import (
	"fmt"
)

func LinearSearch(tab []int, target int) int {
	for i := 0; i < len(tab); i++ {
		if tab[i] == target {
			return i
		}
	}

	return -1
}

func main() {

	fmt.Println("Rolling ...")
	arr := []int{1, 55, 8, 96, 5, 2, 55}

	fmt.Println("=> ", LinearSearch(arr, 8))
	fmt.Println("=> ", LinearSearch(arr, 9))
	fmt.Println("=> ", LinearSearch(arr, 96))

}
