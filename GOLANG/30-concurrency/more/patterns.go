package main

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	price    int
	category string
}

func main() {
	fmt.Println("vim-go")
	c := gen(
		item{10, "shirt"},
		item{20, "shoe"},
		item{30, "shoe"},
		item{40, "drink"},
	)
	/*	for k := range c {
		fmt.Println("this is c:", k)
	}*/
	//out := discount(c)
	c1 := discount(c)
	c2 := discount(c)
	out := FanIn(c1, c2)
	for processed := range out {
		fmt.Println("Category:", processed.category, "Price: ", processed.price)
	}

}

func gen(items ...item) <-chan item {
	out := make(chan item, len(items))
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}

func discount(items <-chan item) <-chan item {
	//fmt.Println("items:", items)
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(time.Second / 2)
			if i.category == "shoe" {
				i.price = 666
			}
			out <- i

		}

	}()
	return out
}

func FanIn(channels ...<-chan item) <-chan item {
	var wg sync.WaitGroup
	out := make(chan item)
	output := func(c <-chan item) {
		defer wg.Done()
		for i := range c {
			out <- i
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

/*

Concurrency patterns:
1. Pipeline:
[Routine] ->[Routine]->[Routine]

2. Fan In/Fan Out

[Routine] ->[Routine]->[Routine]->[Routine]
				  ->[Routine]->


*/
