package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "not enough args !")
		os.Exit(-1)
	}

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		fmt.Println(scan.Text())
	}

}

/*
read a file from the standard input

$ go run main.go < file

*/
