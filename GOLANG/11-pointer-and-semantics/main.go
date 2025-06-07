package main

import "fmt"

type Widget struct {
	ID    int
	Count int
}

func ExpendOne(w Widget) Widget {
	w.Count--
	return w
}

func Expend(w *Widget) Widget {
	w.Count++
	return *w
}
func main() {
	w1 := Widget{ID: 0, Count: 1}
	Expend(&w1)
	Expend(&w1)
	Expend(&w1)
	Expend(&w1)

	fmt.Println("w: ", w1)

	w1 = ExpendOne(w1)
	w1 = ExpendOne(w1)
	w1 = ExpendOne(w1)
	w1 = ExpendOne(w1)
	w1 = ExpendOne(w1)
	w1 = ExpendOne(w1)

	fmt.Println("w1: ", w1)
	fmt.Println()
	fmt.Println()

	s := []int{1, 2, 3}

	// The value returned by range is always a copy
	for _, thing := range s {
		thing++
	} // thing is a copy
	fmt.Println("s : ", s)

	// Use the index if you need to mutate the element
	for i := range s {
		s[i]++
	}
	fmt.Println("s : ", s)
	fmt.Println()
	fmt.Println()
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	fmt.Printf("%#v\n", items)

	for _, item := range items {
		a = append(a, item[:])
	}
	fmt.Println("items: ", items)
	fmt.Println("a: ", a)

}

/*

Pointer													values
shared, not copied          Copied, not shared

Common uses of pointers:
	- some objects can't be copied safely (mutex)
	- some objects are too large to copy efficiently (consider pointers when size > 64 bytes )
	- some methods need to change (mutate) the receiver
	- when decoding protocol data into an object (JSON, etc; often in a variable argument list)
	var r response
	err:= json.Unmrshal (j, &r)
	- when using a pointer to signal a "null" object

*/

/*

- stack allocation is more efficient
- Accessing a variable directly is more efficient than following a pointer.
- Accessing a dense sequence of data is more efficient than sparse data(and array  is faster than a linked list, etc)

=>Heap allocation
Go would prefer to allocate on the stack, but sometimes can't:
	- a func returns a pointer to a local object
	- a local object is captured in a func closure
	- a pointer to a local object is sent via a channel
	- any object is assigned into an interface
	- any object whose size is variable at runtime (slices)

The use of new has nothing to do with it
use the flag -gcflags -m=2 to see the escape analysis


-- Slice safety:
anytime a func mutates a slice that's passe in we must return a copy . Maybe the func doing update to slice and resulting a memory reallocation: (append, update(thing[]thing) { return thing} example
*/
