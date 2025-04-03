package main

// ToDo open network socket pul down (amazon website)
// Ignore JavaScript

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
	<body>
		<h1>My First Heading </h1>
		<p>My first paragraph. </p>
		<p>HTML images are defined with the img tab: </p>
		<img src="xxx.jpg" width="144" height="142">
	</body>
</html>
`

func main() {
	fmt.Println(raw)
	fmt.Printf("%T \n", raw)
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------------")
	fmt.Println(doc)
	/*
		for v := range raw {
			//fmt.Println(string(v))
			fmt.Println(html.EscapeString(string(v)))

		}
	*/
}
