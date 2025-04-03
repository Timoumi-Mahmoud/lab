package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const HELLO = "hello"

func main() {
	rs := make(map[string]int)
	for _, fname := range os.Args[1:] {
		var lc, cc, wc int
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text() //scan by line
			///	fmt.Println(s)
			cc += len(s)
			wc += len(strings.Fields(s))
			lc++
		}
		fmt.Printf("ls: %#d \n", lc)
		fmt.Printf("cc: %#d \n", cc)
		fmt.Printf("wc: %#d \n", wc)

		rs["lc"] += lc
		rs["wc"] += wc
		rs["cc"] += cc

		file.Close()
	}

	fmt.Printf("Total: %#v \n", rs)
}

/*
file, err := os.Open("test.txt")
if err != nil {
	log.Fatal(err)
}
var count int

data := make([]byte, 10000)
count, err = file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])


*/

/*
a, b := 69, 96
fmt.Printf("vim-go: %#v\n", HELLO)
fmt.Printf("vim-go: %v\n", HELLO)
fmt.Printf("a: %d, b: %d\n", a, b)
fmt.Printf("HEX: a: %#x, b: %#x\n", a, b)
fmt.Println()
fmt.Printf("|%10d|,  |%10d|\n", a, b)
fmt.Printf("|%010d|,  |%010d|\n", a, b)
fmt.Printf("|%-10d|,  |%-10d|\n", a, b)
fmt.Println()
s := []int{1, 2, 3}
arr := [3]rune{'a', 'b', 'c'}

fmt.Printf("%T \n", s)
fmt.Printf("%#v \n", s)

fmt.Printf("%T \n", arr)
fmt.Printf("%v \n", arr)
fmt.Printf(" q: %q \n", arr)
fmt.Printf("%#v \n", arr)
fmt.Printf("%v \n", arr)

fmt.Println("\n")
m := map[string]int{"and": 1, "or": 2}

fmt.Printf("%#v \n", m)

fmt.Println(" byte: \n")
ss := "a string"
b := []byte(ss)
fmt.Printf("%T \n", b)
fmt.Printf("%v \n", b)
fmt.Printf(" q: %q \n", b)
fmt.Printf("%#v \n", b)
fmt.Printf("%v \n", b)
fmt.Printf("%v \n", string(b)) //convert bytes to string


*/
