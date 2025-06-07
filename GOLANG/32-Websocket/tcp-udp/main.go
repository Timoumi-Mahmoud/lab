package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("vim-go")
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
	fmt.Fprintf(conn, "Echo - "+string(buf[:n]))

}
