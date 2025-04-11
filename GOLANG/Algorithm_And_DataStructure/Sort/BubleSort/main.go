package main

import (
	"fmt"
)

func BubleSort(tab []int) []int {
	for i := 0; i < len(tab)-1; i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[j] < tab[i] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}

	}

	return tab
}

func main() {

	fmt.Println("Rolling ...")
	///arr := []int{5, 4, 8, 1, 5, 2, 7, 3}
	arr := []int{77, 20, 31, 45, 55, 83, 10, 82}
	arr2 := []int{7, 3, 1}

	fmt.Println("=> before ", arr)
	fmt.Println("=> after ", BubleSort(arr))
	fmt.Println("=> before ", arr2)
	fmt.Println("=> after ", BubleSort(arr2))
}

/*
	Psuedo code:
	arr (sorted array)
	for i=0; i<len(arr)-1; i++{
		for j=i+1; j<len(arr); j++{
					if arr[j]< arr[i]
						swap(arr[i], arr[j]
	return arr

	=>Time Complexity: o(n^2)
*/
