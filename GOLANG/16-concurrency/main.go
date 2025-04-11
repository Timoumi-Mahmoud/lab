package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	go say("yes")
	fmt.Println()
	say("no")
	fmt.Println()

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	x := <-c
	y := <-c
	fmt.Println(x, y, x+y)

	fmt.Println()
	fmt.Println("Buffered Channels ")
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Printf("%T, lent: %d\n", ch, len(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println()
	fmt.Println("Range and Close")

}

/*
- Goroutine : is a lightweight thread managed by the go runtime.
- Channels: are typed conduit in which I can send and receive value with(<-)
* like maps and slices, channels must be created before use.
ch := make(chan type)
* By default sends and receive block until the other side is ready=> this allows goroutines to synchronize without explicit locks or condition values.

* Channels can be buffered:  	ch := make(chan int, 5)



*/
