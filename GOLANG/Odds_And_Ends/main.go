package main

import "fmt"

type shoe int

const (
	tennis shoe = iota
	dress
	sandal
	clog
)

func sum(nums ...int) int { //only the last variable can be a variable list : x ...int, nums ...int is wrong
	var total int
	for _, num := range nums {
		total += num
	}
	return total

}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4))
	s := []int{1, 2, 3, 4}
	fmt.Println(sum(s...))
	s = append(s, s...)

	fmt.Println("s: ", s)
	fmt.Println(sum(s...))
	a, b := uint16(65535), uint16(281)
	fmt.Printf("%016b %#04[1]x\n", a)         // 1111111111111111 0xffff
	fmt.Printf("%016b %#04[1]x\n", a&^0b1111) // 1111111111110000 0xfff0
	fmt.Printf("%016b %#04[1]x\n", a&0b1111)  // 0000000000001111 0x000f
	fmt.Printf("%016b %#04[1]x\n", b)         // 0000000100011001 0x0119
	fmt.Printf("%016b %#04[1]x\n", ^b)        // 1111111011100110 0xfee6
	fmt.Printf("%016b %#04[1]x\n", b|0b1111)  // 0000000100011111 0x011f
	fmt.Printf("%016b %#04[1]x\n", b^0b1111)  // 0000000100010110 0x0116

}
