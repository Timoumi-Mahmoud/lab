package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong ...........\n")
}
