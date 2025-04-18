package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	var i int
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Request ID: ", ctx.Value("requestID"))
				fmt.Println("Context canceled....")
				fmt.Println("=>", ctx.Err())
				close(out)
				return
			case out <- i:
				i++
			}
		}
	}()
	return out

}
func main() {
	//	d := time.Now().Add(5 * time.Millisecond)
	ctx := context.Background() //parent context
	//ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithDeadline(ctx, d)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
	ctx = context.WithValue(ctx, "requestID", "123456")
	for i := range gen(ctx) {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("gen: ", i)
		if i == 10 {
			break
		}
	}
	cancel()
	time.Sleep(1 * time.Second)

}
