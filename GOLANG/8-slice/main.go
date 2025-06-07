package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s []int
	t := []int{}
	u := make([]int, 5)    //len = 5
	v := make([]int, 0, 5) //  len, cap

	fmt.Printf("%d , %d, %T, %5t,-> %#[3]v\n ", len(s), cap(s), s, s == nil)
	fmt.Printf("%d , %d, %T, %5t,-> %#[3]v\n ", len(t), cap(t), t, t == nil)
	fmt.Printf("%d , %d, %T, %5t,-> %#[3]v\n ", len(u), cap(u), u, u == nil)
	fmt.Printf("%d , %d, %T, %5t,-> %#[3]v\n ", len(v), cap(v), v, v == nil)
	a := []int{1, 2, 3, 4}
	j1, _ := json.Marshal(a)
	fmt.Println(string(j1))
	j2, _ := json.Marshal(s)
	fmt.Println(string(j2))
	j3, _ := json.Marshal(t)
	fmt.Println(string(j3))

	u = append(u, 69)
	j4, _ := json.Marshal(u)
	fmt.Println(string(j4))

	v = append(v, 44)
	j5, _ := json.Marshal(v)
	fmt.Println(string(j5))

	fmt.Printf("%d , %d, %T, %5t,-> %#[3]v\n ", len(v), cap(v), v, v == nil)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	aa := [3]int{1, 2, 3}
	bb := aa[:1]

	fmt.Println("aa = ", aa)
	fmt.Println("bb = ", bb)
	cc := bb[0:2]

	fmt.Println("cc = ", cc) //wtf

	fmt.Println(len(bb))
	fmt.Println(cap(bb))

	fmt.Println(len(cc))
	fmt.Println(cap(cc))

	dd := bb[0:1:1] //[i:j:k] len j-1 , cap k-1
	fmt.Println("dd = ", dd)
	fmt.Println(len(dd))
	fmt.Println(cap(dd))
	fmt.Println()
	fmt.Println()

	aaa := [...]int{1, 2, 3}
	bbb := aaa[0:1]
	ccc := bbb[0:2]
	fmt.Printf("aaa[%p] = %v\n", &aaa, aaa)
	fmt.Printf("bbb[%p] = %[1]v\n", bbb)

	ccc = append(ccc, 1000)
	ccc = append(ccc, 10)
	ccc[0] = 999

	fmt.Printf("aaa[%p] = %v\n", &aaa, aaa)
	fmt.Printf("ccc[%p] = %[1]v\n", ccc)

}
