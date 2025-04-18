package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {

		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		//		fmt.Println(resp.Body)
		resp.Body.Close()

	}

}
func main() {
	stopper := time.After(3 * time.Millisecond)
	list := []string{
		"http://localhost:6969",
		"http://localhost:6969",
		"http://localhost:6969",
		"http://localhost:6969",
		"http://notFound",
	}
	results := make(chan result)
	for _, url := range list {
		go get(url, results)
	}

	for range list {
		select {
		case r := <-results:
			if r.err != nil {
				log.Printf("%-20 %s\n", r.url, r.err)
			} else {
				log.Printf("%-20 %s\n", r.url, r.latency)
			}
		case <-stopper:
			log.Fatal("timeout")
		}

	}
}
