package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args !")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]

	fmt.Println("vim-go", old, new)
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		fmt.Println(strings.Split(scan.Text(), old))

		t := strings.Join(s, new)
		fmt.Println(t)
	}

	fmt.Println("Result", old, new)
}

/*

$ go run main.go [value_to_replace] [value_to_replace_with] < file

*/
