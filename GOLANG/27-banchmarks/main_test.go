package main

import "testing"

func Fib(n int, recursive bool) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if recursive {
			return Fib(n-1, true) + Fib(n-2, true)
		}
		a, b := 0, 1
		for i := 1; i < n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

func BenchmarkFib20T(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20, true) // run the Fib function b.N times
	}
}

func BenchmarkFib20F(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20, false) // run the Fib function b.N times
	}
}
