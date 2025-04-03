package main

import "fmt"

func say(i bool) bool {
	return !i
}

func do(b []int) int {
	b[0] = 99
	fmt.Printf("b@ %p\n", b)
	return b[1]
}

func doMap(m1 *map[int]int) {
	(*m1)[3] = 99
	*m1 = make(map[int]int)

	(*m1)[4] = 100
	fmt.Println("m1", m1)
}

func main() {
	m := map[int]int{1: 2, 3: 4, 5: 6, 78: 9}
	f := true

	fmt.Println("Say: ", say(f))
	fmt.Println("F after the execution: ", f)
	a := []int{1, 2, 3}

	fmt.Printf("a@ %p\n", a)
	fmt.Println("do : ", do(a))
	fmt.Println("a after the execution of do: ", a)

	fmt.Println("Before m is: ", m)
	doMap(&m)

	fmt.Println("After m is: ", m)

	val := 100
	defer fmt.Println(val)
	val = 999

	fmt.Println(val)
}
