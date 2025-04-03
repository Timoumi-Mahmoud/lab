package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	fmt.Println("vim-go")
	d1 := []byte("Fuck the world \n")
	fmt.Printf("%#v\n", d1)
	err := os.WriteFile("/tmp/data1", d1, 0644)
	checkErr(err)
	f, err := os.Stat("/tmp/data1")
	checkErr(err)
	println("f", &f)

	fmt.Printf("%#v\n", f.Size())
	f2, err := os.Create("/tmp/data2")
	checkErr(err)
	defer f2.Close()

	d2 := []byte("Fuck the world another time \n")
	n2, err := f2.Write(d2)
	checkErr(err)
	fmt.Printf("worte %d bytes\n", n2)

	o, err := os.Open("/tmp/data2")

	checkErr(err)

	fmt.Printf("open %#v \n", o)
	defer o.Close()

	//var fs []byte
	fs := make([]byte, 10)
	byteCount, err := o.Read(fs)
	checkErr(err)
	fmt.Printf("content : %#v\n", string(fs[:byteCount]))

	w := bufio.NewWriter(f2)
	n4, err := w.WriteString("buffered\n")
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()

}
