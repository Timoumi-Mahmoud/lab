package main

import (
	"fmt"
	"slices"
)

func BinarySearch(tab []int, target int) int {
	l := 0
	r := len(tab) - 1

	for l <= r {

		m := (l + r) / 2
		if tab[m] < target {
			l = m + 1
		} else if tab[m] > target {
			r = m - 1
		} else {
			return m
		}

	}
	return -1
}

func main() {

	fmt.Println("Rolling ...")
	arr := []int{1, 55, 8, 96, 5, 2, 55}

	slices.Sort(arr)

	fmt.Println("=> tes ", BinarySearch(arr, 8))
	fmt.Println("=> ", BinarySearch(arr, 1))
	fmt.Println("=> ", BinarySearch(arr, 96))

}

/*
Psuedo code:
arr (sorted array), target (searched number)

left = 0
right = arr.length -1
while l<= r do :
	middle = (left+right) /2
	if arr[middle]== target then return middle
	esle if arr[middle] < target then  left = middle +1
	esle if arr[middle] > target then  right = middle -1

return -1


=>Time Complexity: O(log(n))
but the array have to be sorted initialy in case of Golang the function slices.Sort hava a time complexity of O(n*log(n))  see doc (https://go.dev/src/sort/sort.go)
so the overall complexity will be: n*log(n)  *  log(n)

*/
