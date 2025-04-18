package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	fmt.Println("vim-go")
	s := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go s.Inc("blbla")
	}
	time.Sleep(time.Second)
	fmt.Println(s.Value("blbla"))

}
