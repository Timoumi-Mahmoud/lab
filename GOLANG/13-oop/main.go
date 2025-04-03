package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string
	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v
	for i, x := range v {

		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
	fmt.Printf("%T %[1]v\n", s)

}

/*
OOP:
- abstraction
- encapsulation
- ploymorphism
- inheritance




Abstraction: decoupling behavior from the implementation details.(unix file system api example open , close, read, write and ioctl)

Encapsulation: hiding implementation details from misuse.
(controlling the visibli)

Polymorphism: (many shapes")

Inheritance

*/

/*

## Interfaces : specifies abstract behavior in terms of methods: it does it by listing a method sets in an interface.
## Methods : as sepcial typefo function,,
it has a receiver parameter befrore the function name parameter.
' func (is  IntSlice) String string {}  '
a method may take a pointer or value receiver but not both.






*/
