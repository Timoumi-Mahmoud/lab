package main

import "testing"

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

func mkList(n int) *node {
	var h, t *node

	h = insert(0, h) //head
	t = insert(1, h)
	for i := 2; i < n; i++ {
		t = insert(i, t)
	}
	return h
}
func sumList(h *node) (i int) {
	for n := h; n != nil; n = n.t {
		i += h.v
	}
	return
}

func mkSlice(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}
	return r
}

func sumSlice(s []int) (i int) {

	for _, v := range s {
		i += v
	}
	return
}

func Benchmarklist(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := mkList(1200)
		sumList(l)

	}
}

func Benchmarkslice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := mkSlice(1200)
		sumSlice(l)
	}
}
