package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var (
	ErrorWordNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot perform operation on word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if !ok {
		return "", ErrorWordNotFound
	}
	return result, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

func main() {

	dictionary := Dictionary{"test": "this is just a test"}
	mTwo := Dictionary{"testTwo": "this is just a testTwo"}
	fmt.Printf("%T\n", dictionary)
	fmt.Printf("%v\n", dictionary)

	mTwo["testThree"] = "this is just a testThree"
	fmt.Printf("%v\n", mTwo)
	delete(mTwo, "testTwo")
	fmt.Printf("%v\n", mTwo)
	/*
		mTwo["hello"] = "world" // panic: assignment to entry in nil map
		fmt.Printf("%T\n", mTwo)
		fmt.Printf("%v\n", mTwo)
	*/
	/*
		var dictionary = map[string]string{}
		// OR
		var dictionary = make(map[string]string)
	*/

}
