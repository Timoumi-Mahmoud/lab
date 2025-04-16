package main

import (
	"fmt"
	"regexp"
)

var uure = `^(http|https)://([a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,4})` +
	`(?::([0-9]+))?\/?([a-zA-Z0-9\-\._\?\,\'\/\\\+&amp;%\$#\=~]*)$`
var ufmt = regexp.MustCompile(uure)
var test = []string{
	"http://mahmoud.com/hello",
	"http://mahmoud.com:8080/hello/",
	"http://mahmoud.com:8080/hello?a=1&b=2",
	"http://mahmoud.com:8080/hello?a=1&b=2&c=3",
	"mahmoud.com",
}

func main() {
	for i, t := range test {
		match := ufmt.FindStringSubmatch(t)
		fmt.Printf("%d: %q\n", i, match)
	}
}
