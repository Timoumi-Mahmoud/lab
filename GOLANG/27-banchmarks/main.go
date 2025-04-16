package main

import "fmt"

func fib(n int) (int, int) {
	i := 0
	a, b := 0, 1
	for i < n {
		a, b = b, a+b
		i++

	}
	return a, b
}

func main() {
	x, y := fib(6)
	fmt.Printf("fib(6): x=%d, y=%d \n", x, y)

}

/*
Benchmarks: performance, allocations
* lives in test files ending with _test.go
* run Benchmarks with go test -bench
* go only runs the BenchmarkXXXX funcs

$ go test -bench=.
goos: linux
goarch: amd64
pkg: bench.tn
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
												(iterations)    (time ns/op)
BenchmarkFib20T-8   	   29776	     41040 ns/op
BenchmarkFib20F-8   	100000000	        10.16 ns/op
PASS
ok  	bench.tn	2.662s

*/
