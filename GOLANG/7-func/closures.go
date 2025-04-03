package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func do(d func()) {
	d()
}
func main() {
	f, g := fib(), fib()
	/*	for x := f(); x < 100; x = f() {
		fmt.Println("vim-go", x)
	}*/

	fmt.Println("vim-go", f(), f(), f(), f(), f(), f())
	fmt.Println("vim-go", g(), g(), g(), g(), g(), g())
	/* f and g gets different a and b variables*/

	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
		do(v)
	}

	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}

}
