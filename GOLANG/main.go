package main

import (
	"fmt"
	"strings"
)

type x int

func main() {

	defer fmt.Println("vim-go")
	text := "bla bla not a bla man ok fuck fuck fuck ok fuck"
	//textSlice := []string{text}
	textSlice := strings.Split(text, " ")
	fmt.Println("textSlice: ", textSlice)
	words := make(map[string]int)
	for _, v := range textSlice {
		words[v]++
		fmt.Println("v: ", v)
	}
	fmt.Println("words finally is: ", words)
	var i int
	for i := 5; i < 10; i++ {
		fmt.Println("i: ", i)
	}
	fmt.Println("i outsid the for loop: ", i)

	a := 266
	switch {
	case a < 100:
		fmt.Println("more than 100")
	case a > 100:
		fmt.Println("less than 100")
	default:
		fmt.Println("Under Zero")
	}

	var bara x

	fmt.Printf("bara: %T \n", bara)
	bb := 12

	fmt.Println("bb: ", bb)
	bara = x(bb)

	fmt.Println("bara: ", bara)

}

func init() {
	fmt.Println("Init func man ...........................")
}
