package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "èlite"
	fmt.Println("s", s[3])
	sv := "lite"
	fmt.Println("sv", sv[3])

	fmt.Printf("%44T %[1]v\n", s)
	fmt.Printf("%4T %[1]v , len: %d\n", []rune(s), len(s))

	b := []byte(s)
	fmt.Printf("%4T %[1]v , len: %d\n", b, len(b))

	ss := "hello, world"
	hello := ss[:5]
	world := ss[7:]
	fmt.Println(ss, hello, world)
	t := ss
	fmt.Println(t)

	fox := "the quick brown fox"

	fox += "es"
	slowFox := fox[0:4] + "slow" + fox[9:]
	fmt.Println(fox)
	fmt.Println(slowFox)

	//////Strings has many function  on strings
	f := "a string f"
	fmt.Println(
		strings.Contains(f, "g"),
		strings.Contains(f, "x"),
		strings.HasPrefix(f, "a"),
		strings.Index(f, "f"),
	)

	f = strings.ToUpper(f)
	fmt.Println(f)
	fSlice := []string{f}

	newJoin := strings.Join(fSlice, ",")
	newSplit := strings.Split(newJoin, "---------")

	fmt.Printf(" -Type fSlice : %T \n", fSlice)
	fmt.Printf(" -Type newJoin: %T\n", newJoin)

	fmt.Println(newJoin)
	fmt.Println(newSplit)

	words := strings.Split("a:b:c", ":")
	fmt.Println(words)
	fmt.Println(strings.Join(words, ":"))
}
