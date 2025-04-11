package main

import (
	"fmt"
	"strings"
)

func ReverseString(input string) string {

	var tmp []string

	for i := len(input) - 1; i >= 0; i-- {
		tmp = append(tmp, string(input[i]))
	}
	justString := strings.Join(tmp, "")

	return justString
}

func main() {

	fmt.Println("Rolling ...")
	stringInput := "hello"
	fmt.Println("=> after ", ReverseString(stringInput))

}

/*
- given a string a return the reverse of that string: example hello=> olleh

*/
