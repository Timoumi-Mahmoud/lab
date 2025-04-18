package main

import (
	"fmt"
)

func TwoSums(nums []int, target int) (r []int) {
	prev := map[int]int{} //int int

	// arr := []int{89,78,  1,2, 3 } 5
	// 6:
	fmt.Println("nums: ", nums)
	for i, v := range nums {
		diff := target - v
		fmt.Println("Diff= target- v ", diff)
		index, ok := prev[diff]
		fmt.Println("ok: ", ok)
		fmt.Println("index: ", index)
		if ok {
			return []int{index, i}
		}
		prev[v] = i

		fmt.Println("prev: ", prev)
	}

	return
}

/*
func TwoSums(nums []int, target int) (r []int) {
	for i, v := range nums {

		if v >= target {
			continue
		}
		diff := target - v

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == diff {
				r = append(r, i)
				r = append(r, j)
			}
		}
	}

	return
}
*/

/*
//second Approach

	func TwoSums(nums []int, target int) (r []int) {
		for i, v := range nums {
			if v >= target {
				continue
			}
			for j := i + 1; j < len(nums); j++ {
				if (nums[j] + nums[i]) == target {
					r = append(r, i)
					r = append(r, j)
				}
			}
		}

		return
	}
*/

func main() {

	fmt.Println("Rolling ...")
	//arr := []int{1, 5, 4, 2}
	arr := []int{55, 47, 69, 1, 2, 3}
	fmt.Println("=> reuslt:  ", TwoSums(arr, 5))

}

/*
- given an arry nums of integers and an integer target , return the indices of two numbers such that they add up to target.
*/
