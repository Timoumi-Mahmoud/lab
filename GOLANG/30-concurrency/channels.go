package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
	// A sender can close a channel to indicate that no more values wil be sent ; only the sender can close the receive can't

}

// select lets a goroutine wait on multiple communication operations
// a select blocs until one of its cases can run, that it execute that case
func fibWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}

	}
}

func main() {
	s := []int{10, 5, 4, 1, 3, 2, 2, 3}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println("sum:", x, y, y+x)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch2 := make(chan int, 10)
	fmt.Println("cap(ch2), ch2", cap(ch2), ch2)
	go fib(cap(ch2), ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("<-c", <-ch3)
			if i == 8 {

				quit <- 0
			}
		}
	}()
	fibWithSelect(ch3, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*
ch <-v // send v to channel ch
v:= <-ch // receive from ch

*/
