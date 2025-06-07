package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	out := make(chan string)
	go makeOrders("Snikers", out)
	go makeOrders("T-shirt", out)
	go makeOrders("Jeans", out)
	for k := range out {
		fmt.Println(k)
	}
}

func makeOrders(item string, out chan<- string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Second / 2)
		out <- fmt.Sprintf("order %d: %s", i, item)
	}
	close(out)
}
