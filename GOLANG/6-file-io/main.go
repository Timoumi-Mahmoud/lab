package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rs := make(map[string]int)
	for _, fname := range os.Args[1:] {
		var lc, cc, wc int
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			cc += len(s)
			wc += len(strings.Fields(s))
			lc++
		}
		fmt.Printf("lc: %#d \n", lc)
		fmt.Printf("cc: %#d \n", cc)
		fmt.Printf("wc: %#d \n", wc)

		rs["lc"] += lc
		rs["wc"] += wc
		rs["cc"] += cc

		file.Close()
	}

	fmt.Printf("Total: %#v \n", rs)
}
