package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		//fmt.Println("scan:", scan.Text())
		words[scan.Text()]++
	}
	fmt.Println("words", words)
	type kv struct {
		key string
		val int
	}
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	fmt.Println("ss is :", ss)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
