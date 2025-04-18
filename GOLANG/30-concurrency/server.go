package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d<h1>", <-nextID) //reader
	//	nextID++

}
func counter() {
	for i := 0; ; i++ {
		nextID <- i // writer
	}
}
func main() {
	fmt.Println("vim-go")
	go counter()
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":6969", nil))
}
