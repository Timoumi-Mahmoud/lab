package main

import (
	"fmt"
	"regexp"
)

var phre = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`
var pfmt = regexp.MustCompile(phre)

func main() {
	orig := "call me at (214) 514-3232 today"
	match := pfmt.FindStringSubmatch(orig)
	fmt.Printf("%q\n", match)
	fmt.Printf("=>%q\n", match[0])
	intl := pfmt.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")
	fmt.Println(intl)

}
