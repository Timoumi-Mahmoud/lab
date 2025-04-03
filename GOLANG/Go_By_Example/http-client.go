package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:6969/heiiiiillo")
	if err != nil {
		log.Fatal("unable to send request : ", err)
	}
	defer resp.Body.Close()
	fmt.Println("request status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println("request body: ", scanner.Text())
	}
}
