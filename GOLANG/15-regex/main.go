package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// strings pkg
	s := "hello world"

	fmt.Println("To upper", strings.ToUpper(s))
	fmt.Println(strings.HasPrefix(s, "re"))
	fmt.Println(strings.HasSuffix(s, "!!"))
	fmt.Println(strings.Contains(s, "rl"))
	fmt.Println(strings.LastIndex(s, "d"))
	fmt.Println(strings.LastIndex(s, "l"))
	fmt.Println(strings.LastIndexByte(s, 'l'))
	fmt.Println(strings.Replace(s, "hello", "bye", 1))
	fmt.Println(strings.ReplaceAll(s, "l", "$"))

	fmt.Println(s)
	fmt.Println()
	fmt.Println("Regex")
	fmt.Println()

	te := "aba abba abbba"
	re := regexp.MustCompile("a+") // +: one or more
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Println(mm)
	fmt.Println(id)
	fmt.Printf("type of id: %T\t, type of mm: %T\n", id, mm)

	up := re.ReplaceAllStringFunc(te, strings.ToUpper)
	fmt.Println(up)

}

/*
.    : is any character
.*   : is zero or more
.?   : is zero or one(prefer one) maximal much

a{n}    : is n repetitions of the letter "a"
a{n, m} : is n to m repetitions of the letter "a"
[a-z]   : is a character class (here letters a-z)
[^a-z]  : si an negated class (here anything except a-z)

xy    : is "x" followed by "y"
x|y   : is either "x" or "y"
^x    : is "x" at the beginning
x$    : is "x" at the end
^x$   : is "x" by itself
\b    :is word boundray
\bx\b : is word "x" by itself
(x)   :is a capture group


\d is a decimal digit
\w is a word character ([0-9A-Za-z])
\s is whitespace
[[:alpha:]]
[[:alnum:]]
[[:punct:]]
[[:xdigit:]]
*/
