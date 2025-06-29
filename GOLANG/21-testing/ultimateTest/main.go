package main

import (
	"log"
	"net/http"

	handler "bbb.ts/handlers"
)

func main() {

	handler.Routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
