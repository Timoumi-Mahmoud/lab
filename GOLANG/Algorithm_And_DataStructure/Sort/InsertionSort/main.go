package main

import (
				"fmt"
				"reflect"
	   )

func InsertionSort(tab []int) []int {
		for i := 1; i < len(tab); i++ {
position :=i 
				  temp_value:=tab[i]
				  for position >0 && tab[position-1] > temp_value {
						  tab[position] = tab[position-1]
								  position = position-1
				  }
		  tab[position] = temp_value
		}

		return tab
}


func main() {

		fmt.Println("Rolling ...")
				arr := []int{5, 4, 8, 1, 5, 2, 7, 3}

		fmt.Println("=> before ", arr)
				fmt.Println("=> after ", InsertionSort(arr))
				a := []int{1, 3, 2}
b := []int{1, 2, 3}

equal := reflect.DeepEqual(a, b)
			   fmt.Println(equal)

}

/*
   Psuedo code:
   arr (sorted array)
   for 1 range len(tab){
		postion=i 
		temp_value=tab[i]
		while position> 0 && tab[postion-1] > temp_value{
				swap(tab[position], tab[position- 1]
				position = position-1

		tab[position]=temp_value





   =>Time Complexity: o(n^2)
 */
