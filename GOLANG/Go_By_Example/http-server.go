package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello from %#v\n ", r.URL.Path[1:])

}
func main() {

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
