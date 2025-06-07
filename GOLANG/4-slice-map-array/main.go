package main

import "fmt"

func main() {
	var arr [4]int

	fmt.Printf("%8T\n", arr)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	temp := arr
	arr[0] = 999
	fmt.Println("ARR: ", arr)
	fmt.Println("temp: ", temp)

	fmt.Println("--------------")

	var a []int
	a = append(a, 69)
	a = append(a, 68)
	a = append(a, 67)
	b := a
	fmt.Println("b: ", b)
	b = append(b, 9999)

	a = append(a, 1555)

	fmt.Println("a: ", a)
	d := make([]int, 5)

	fmt.Println("d: ", d)
	fmt.Println("--------------")
	t := []byte("string")

	fmt.Println(len(t), t)
	fmt.Println(t[2])
	fmt.Println(t[3:5], len(t[3:5]))

	fmt.Println("--------------")

	var w = [...]int{1, 2, 3} // array of len(3)
	var x = []int{0, 0, 0}    // slice of len(3

	y := do(w, x)
	fmt.Println(w, x, y)

	/*j
	var m map[string]int // nil , no storage: I can read from a nil map but inserting will panic
	m["one"] = 1
	fmt.Println("m: ", m)
	*/
	p := make(map[string]int)
	aa := p["bara"]
	fmt.Println("aa: ", aa)

	p["bara"] = 99

	fmt.Println("p: ", p)
	fmt.Println("aa: ", aa)

}
func do(a [3]int, b []int) []int {
	a[0] = 4            // w unchanged =>>>>>>>>>> because array passed by value
	b[0] = 3            // x changed =>>>>>>>>>>>> because slice are passed by reference
	c := make([]int, 5) // []int{0, 0, 0, 0, 0}
	c[4] = 42
	copy(c, b) // copies only 3 elts
	return c
}
