package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	strs := []string{"c", "a", "b"}

	fmt.Println(slices.IsSorted(strs))
	slices.Sort(strs)
	fmt.Println(strs)
	fmt.Println(slices.IsSorted(strs))
	fruits := []string{"peach", "banana", "kiwi"}
	/*
		lenCmp := func(a, b string) int {
			return cmp.Compare(len(a), len(b))
		}
		slices.SortFunc(fruits, lenCmp)
	*/

	slices.SortFunc(fruits, func(a, b string) int {
		return cmp.Compare(len(b), len(a))
	})
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}
	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(b.age, a.age)
	})
	fmt.Println(people)

}
