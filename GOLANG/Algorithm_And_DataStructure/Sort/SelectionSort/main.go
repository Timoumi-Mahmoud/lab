package main

import (
	"fmt"
)

func SelectionSort(tab []int) []int {
	for i := 0; i < len(tab)-1; i++ {
		iMin := i
		for j := i + 1; j < len(tab); j++ {
			if tab[j] < tab[iMin] {
				iMin = j
			}
		}
		if iMin != i {
			tab[i], tab[iMin] = tab[iMin], tab[i]
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
	fmt.Println("=> after ", SelectionSort(arr))
	fmt.Println("=> before ", arr2)
	fmt.Println("=> after ", SelectionSort(arr2))

}

/*
	Psuedo code:
	arr (sorted array)
	for i=0; i<len(arr)-1; i++{
			iMin = i
			for j=i+1; j<len(arr); j++{
					if arr[j]< arr[iMin]
						iMin=j
		if iMin != i
		swap(arr[iMin], arr[j]
	return arr

	=>Time Complexity: o(n^2)
*/
