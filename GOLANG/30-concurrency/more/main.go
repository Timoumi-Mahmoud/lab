package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(counter *int, wg *sync.WaitGroup) {
	*counter++
	wg.Done()
}

func decrement(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	*counter--
	mu.Unlock()
	wg.Done()
}

func main() {

	num := 99

	var wg sync.WaitGroup
	wg.Add(6)
	go increment(&num, &wg)
	go increment(&num, &wg)
	go increment(&num, &wg)
	go increment(&num, &wg)
	go increment(&num, &wg)
	go increment(&num, &wg)
	wg.Wait()
	fmt.Println("vim-go")

	var mu sync.Mutex
	wg.Add(5)
	go decrement(&num, &wg, &mu)
	go decrement(&num, &wg, &mu)
	go decrement(&num, &wg, &mu)
	go decrement(&num, &wg, &mu)
	go decrement(&num, &wg, &mu)
	wg.Wait()

	wg.Add(3)
	go func() {
		num++
		wg.Done()
	}()
	go func() {
		num++
		wg.Done()
	}()
	go func() {
		num++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(num)
	fmt.Println()

	channel := make(chan int, 5)

	numTwo := 5
	channel <- numTwo

	close(channel)
	response, closed := <-channel

	fmt.Println(closed)

	if !closed {
		fmt.Println("the channel is closed")
	} else {
		fmt.Println("the channel is open")
	}
	response, closed = <-channel

	fmt.Println(closed)
	if !closed {
		fmt.Println("the channel is closed")
	} else {
		fmt.Println("the channel is open")
	}

	fmt.Printf("%T\n", channel)
	fmt.Println("response:", response)

	fmt.Println()
	fmt.Println()
	out := make(chan int)
	oo := make(chan int)
	go count(out)
	//	for v := range out {
	//		fmt.Println("Message from count:", v)
	//	}

	go mod(out, oo)
	for k := range oo {
		fmt.Println("Message from mod:", k)
	}
}

func count(out chan<- int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second / 2)
		out <- i
	}
	close(out)
}

func mod(input <-chan int, out chan<- int) {
	for k := range input {
		out <- k + 100
	}

	close(out)
}

/*

Goroutine: are way to have section of code attempts to execute
it's not like thread in Java, I can have 100s or goroutine executing on the same processor

Concurrency vs parallelism: 2 processes run in the same time for sure in parallelism, but in concurrency is not for sure maybe one process is running but it can be interrupted and then resume executing.


*/
