package main

import (
	"errors"
	"fmt"
)

var (
	ErrorNotFound = errors.New("resource not found")
)

type error interface {
	Error() string
}
type Fizgig struct{}

func (f Fizgig) Error() string {
	return "Your fizgig is bent"
}

func abc() {
	panic("OMG")

}
func main() {
	f := Fizgig{}

	fmt.Println(f.Error())
	fmt.Printf("%T\n", f.Error())

	fmt.Println(ErrorNotFound)
	fmt.Printf("%T\n", ErrorNotFound)

	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover: ", p)
		}
	}()

	s := make([]string, 0)
	s = append(s, "hello")
	fmt.Println("s: ", s)
	var b string
	fmt.Println("b: ", len(b))

	abc()
}

/*
- Error object is an interface
1. Normal errors: result from input or external conditions.example file not found, network close, memory space...
2. Abnormal errors bugs: result from invalid program logic example a nil pointer => go handle it with panic
Fail Hard, Fail Fast because:
		- if my server crashes, it will get immediate attention
		- logs are often noisy
		- proactive log searches are rare
		- in distributed system, crash failures are the safest type to handle: it's better than to be zombie or babble or corrupt the DB
		Byzantine failures

* When Should I use panic:
- I can't walk a data structure that i built
- off-by-one bug encoding bytes
=> should be used when our assumptions of my own programming design are wrong.(B-tree data structure example)

------------------------------------------------------------------
Error edge cases:
Desing my abstraction so that most (or all) operation are safe:
- reading from a nil map
- appending to a nil slice
- deleting a non-existing item from a map
- taking the length of an unitialized string
-----------------------------------------------------------------
Proactively prevent problems:
- Every piece of data in my software shouls start life in a valid state.
- Every trasfomation should leave it in a valid state
* break large programs into small peices that I can undrestand.
* hide inforamtion to reduce the change of corruption
* avoid cleaver code and side effects
* avoid unsafe operations
* assert my invariants
* never ignore errors
* test, test, test
- Never, ever accept input from a user (or environment) without validation.
*/
