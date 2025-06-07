package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushBack(2)
	mylist.PushBack(3)
	mylist.PushFront(69)

	el := mylist.Front()
	mylist.Remove(el)
	for element := mylist.Front(); element != nil; element = element.Next() {
		if element.Value == 2 {
			mylist.Remove(element)
		}

	}

	for element := mylist.Front(); element != nil; element = element.Next() {

		fmt.Println(element.Value)
	}
}
