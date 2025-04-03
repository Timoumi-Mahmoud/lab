package main

import (
	"fmt"
)

type album1 struct {
	title string `json:"name"`
	//year  int
}
type album2 struct {
	title string `json:"name"`
}
type album3 struct {
	title  string `json:"name"`
	copies int
}

func main() {
	var a1 = album1{
		"the red album",
		//1999,
	}

	var a2 = album2{
		"the pink album",
	}
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	//a1 = a2 should be converted first (see conversion rule for struct bellow)
	a1 = album1(a2)
	fmt.Println("a1: ", a1)

	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The white album",
		"The Beatles",
		1968,
		100_000_000,
	}
	var pAlbum = &album

	album.year = 2000
	fmt.Println()

	fmt.Printf(" %+[1]v \n", album)
	fmt.Printf(" %+[1]v \n", pAlbum)

	fmt.Println()
	var a3 = album3{
		"the devil album",
		10000000,
	}

	fmt.Printf(" %+[1]v \n", a3)
	soldAnother(a3)
	fmt.Printf(" %+[1]v \n", a3)
	soldAnotherWithPointer(&a3)
	fmt.Printf(" %+[1]v \n", a3)

	isPresent := map[int]bool{} // a set typ (instead of bool)
	isPresent[1] = true
	isPresent[1] = true
	isPresent[2] = true
	fmt.Println("isPresent: ", isPresent)

	isPresentWithStruct := map[int]struct{}{} // a set typ (instead of bool)
	isPresentWithStruct[1] = emptyStruct{}
	isPresentWithStruct[1] = emptyStruct{}
	isPresentWithStruct[3] = emptyStruct{}
	fmt.Println("isPresentWithStruct: ", isPresentWithStruct)

}

type emptyStruct struct {
}

func soldAnother(a album3) {
	a.copies++
}
func soldAnotherWithPointer(a *album3) {
	a.copies++ //equivalent  to (*a).copies
}

/*

struct type are compatible if :
 - the fields have the same types and names.
 - in the same order
 - and the same tags  /// not valid after go 1.8

*/

/*
a struct with no fields is useful: it takes up no space:
	1) set type (instead of bool)
	var isPresent map[int]struct{} // a set type
	See Above the example of isPresent and isPresentWithStruct

	2) a very cheap channel type
	done := make(chan struct{})


	Struct{} is a singleton : go keeps a variable of type struct{} somewhere
*/
