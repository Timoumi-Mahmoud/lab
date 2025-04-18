package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Username string `json:"username"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

func main() {
	f, err := os.OpenFile("write.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	//os.O_TRUNC
	if err != nil {
		log.Fatal(err)
	}

	n, err := f.Write([]byte("blblblba\n"))
	if err != nil {
		log.Fatal("write err:", err)
	}

	fmt.Println(f)
	fmt.Println(n)

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	u := user{Username: "khadija", Lastname: "blbbl", Password: "0000"}
	if err := enc.Encode(u); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
	f.Write(buf.Bytes())

}
