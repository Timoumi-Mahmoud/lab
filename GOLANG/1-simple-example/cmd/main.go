package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}

/*
$ go run cmd/main.go or go run ./cmd
$ go test
$ go build ./cmd/main.go
*/
