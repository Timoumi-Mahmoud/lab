package main

import (
	"fmt"
)

type StringStack struct {
	data []string
}

func (s *StringStack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *StringStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		s.data = s.data[:l-1]
		return t
	}
	//log.Fatal("pop from empty stack man")
	panic("pop from empty stack man")

}

func main() {
	fmt.Println()
	s := StringStack{[]string{"hh", "ff"}}
	fmt.Println("s: ", s)
	s.Push("fffffffffff")
	fmt.Println("s: ", s)

	fmt.Println("", s.Pop())
	fmt.Println("", s.Pop())
	fmt.Println("", s.Pop())

	fmt.Println("after pop: ", s)
}
