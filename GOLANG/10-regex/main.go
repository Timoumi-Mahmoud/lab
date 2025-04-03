package main

import (
	"fmt"
	"regexp"
)

var uure = `^(http|https)://([a-zA-Z0-9\-\.]+[a-zA-Z]{2,4})` +
	`(?::([0-9]+))?\/?([a-zA-Z0-9\-\._\?\,\'\/\\\+&amp;%\$#\=~]*)$`

func main() {

	ufmt := regexp.MustCompile(uure)
	fmt.Println(ufmt.MatchString("http://matt.com:8080/hello?a=1&b=2"))
	fmt.Println(ufmt.MatchString("http://localhost:6060/pkg/regexp/"))
}
