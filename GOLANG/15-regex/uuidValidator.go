package main

import (
	"fmt"
	"regexp"
)

var uure = `^[[:xdigit:]]{8}-[[:xdigit:]]{4}-` +
	`[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`

var ufmt = regexp.MustCompile(uure)
var test = []string{
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbblkjlm", // fails because i added ^ and $
	"072665ee-a034-6cc3-a2e8-9f1822c4ebbb",      // wrong version
	"072665ee-a034-4cc3-72e8-9f1822c4ebbb",      // wrong format
}

func main() {
	for i, t := range test {
		if !ufmt.MatchString(t) {
			fmt.Println(i, t, "fails")
		}
	}
}
