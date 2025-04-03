package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Page  int      `json:"pageee"`
	Words []string `json:"words,omitempty"` // lower case is not encoded (either exported)
}

func main() {
	r := &Response{Page: 1, Words: []string{"up", "in", "out"}}
	r2 := &Response{Page: 1, Words: []string{}}

	fmt.Printf("%#v\n ", r)

	fmt.Println()
	j1, err := json.Marshal(r)
	if err != nil {
		log.Println("err: ", err)
	}

	fmt.Println(string(j1))
	fmt.Println()

	fmt.Println()
	j2, err := json.Marshal(r2)
	if err != nil {
		log.Println("err: ", err)
	}

	fmt.Println(string(j2))
	fmt.Println()

}
