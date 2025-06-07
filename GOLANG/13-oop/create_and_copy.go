package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}
func main() {
	var c ByteCounter

	f1, err := os.Open("a.txt")
	checkErr(err)

	f2, err := os.Create("b.txt")
	checkErr(err)

	n, err := io.Copy(f2, f1)
	checkErr(err)

	fmt.Println("copied: ", n, " bytes")
	fmt.Println(c)

}
