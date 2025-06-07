package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanich", false, "Use spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use spanish language.")
}

func main() {
	flag.Parse()
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
	if spanish == true {
		fmt.Println("Hola %s!\n", *name)

	} else {

		fmt.Println("Hello %s!\n", *name)
	}
}
