package main

import "fmt"

type node struct {
	v int
	t *node
}

func insert(i int, h *node) *node {
	t := &node{i, nil}
	if h != nil {
		h.t = t
	}
	return t

}

func main() {
	fmt.Println("link list")
	var h, t *node
	h = insert(0, h)
	t = insert(1, h)
	for i := 2; i < 10; i++ {
		t = insert(i, t)
	}
	sumList(h)

}
func sumList(n *node) {
	for h := n; h != nil; h = h.t {
		fmt.Println(h.v)
	}
}
