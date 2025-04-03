package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {

	log.Println("Server :Hello Handler Started")
	ctx := r.Context()

	defer log.Println("Server :Hello Handler Ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)

	}
}
func main() {

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":6969", nil))
}

/*
context.Context : control cancellation, carries deadlines, cancel signals..
A context.Context is created for each request by the net/http machinery it's available in r.Context()


*/
