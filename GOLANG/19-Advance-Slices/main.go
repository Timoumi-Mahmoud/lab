package main

import "fmt"

func concatSlices() []int {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...)
	return a
}

func copySliceV1() []int {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))

	count := copy(b, a)
	fmt.Println("number of copied elements: ", count)
	return b

}

func copySliceV2() []int {
	a := []int{1, 2, 3}
	b := []int{}
	//	b := make([]int, len(a), cap(a)) // wrong
	for i := range a {
		b = append(b, a[i])
	}
	return b
}

func copySliceV3() []int {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))

	b = append([]int(nil), a...)
	return b
}

func deleteIndexV1() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	index := 1
	a = append(a[:index], a[index+1:]...)
	return a
}

func deleteNoOrder() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	index := 1
	a[index] = a[len(a)-1]
	return a[:len(a)-1]
}

func cutSlice() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	i := 1
	j := 4
	return append(a[:i], a[j+1:]...)
	//return a[i : j+1]
}

func filterInPlace() []int {
	a := []int{-1, 2, -3, -4, 5, 6}
	FilterFunc := func(x int) bool {
		return x < 0
	}

	n := 0
	for _, v := range a {
		if FilterFunc(v) {
			a[n] = v
			n++
		}
	}
	return a[:n]
}

func filterIntoNewSlice() []int {
	a := []int{-1, 2, -3, -4, 5, 6}
	FilterFunc := func(x int) bool {
		return x > 0
	}

	b := []int{}
	for _, v := range a {
		if FilterFunc(v) {
			b = append(b, v)
		}
	}
	return b
}

func insert() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	index := 2
	value := 69
	return append(a[:index], append([]int{value}, a[index:]...)...)

}

func insertSlice() []int {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	index := 2
	a = append(a[:index], append(b, a[index:]...)...)
	return a
}

func push() []int {
	a := []int{1, 2, 3}
	v := 69
	return append(a, v) // push at the End of the slice

	//	return append([]int{v}, a[0:]...)  // push at the start of the slice

}

func pop() (int, []int) {
	// remove a value from the end of the slice
	a := []int{1, 2, 3}
	return a[len(a)-1], a[:len(a)-1]

}

func shift() []int {
	// remove a value from the beginning of the slice
	a := []int{1, 2, 3}
	return a[1:len(a)]

}

func unshift() []int {
	a := []int{1, 2, 3}
	v := 69
	return append([]int{v}, a[0:]...) // add a value at the beginning of the slice

}

func main() {
	fmt.Println("vim-go")
	fmt.Println("concatSlices: ", concatSlices())
	fmt.Println("copySliceV1: ", copySliceV1())
	fmt.Println("copySliceV2: ", copySliceV2())
	fmt.Println("copySliceV3: ", copySliceV3())
	fmt.Println("deleteIndexV1: ", deleteIndexV1())
	fmt.Println("deleteNoOrder: ", deleteNoOrder())
	fmt.Println("cutSlice: ", cutSlice())
	fmt.Println("filterInPlace: ", filterInPlace())
	fmt.Println("filterIntoNewSlice: ", filterIntoNewSlice())
	fmt.Println("insert: ", insert())
	fmt.Println("insertSlice: ", insertSlice())
	fmt.Println("push: ", push())
	value, finalSlice := pop()

	fmt.Println("pop: ", value, finalSlice)

	fmt.Println("shift: ", shift())
	fmt.Println("unshift: ", unshift())

}
