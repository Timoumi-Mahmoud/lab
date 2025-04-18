package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func fileReader() {
	f, err := os.Open("small.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	readToStout(f, 4)

}

func stringReader() {
	s := strings.NewReader("very short message")
	readToStout(s, 10)
}

func connReader() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprint(conn, "GET HTTP 1.1\r\n\r\n")
	readToStout(conn, 10)
}

func readToStout(r io.Reader, bufSize int) {

	buf := make([]byte, bufSize)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
			//			fmt.Println(buf[:n])
		}
	}

}
func main() {
	fileReader()
	stringReader()
	connReader()
}
