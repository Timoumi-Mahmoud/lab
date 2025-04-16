package models

import (
	"database/sql"
)

// create and export global variable to hold the db connetion pool
var DB *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]Book, error) {
	result, err := DB.Query("select * from book")
	if err != nil {
		return nil, err
	}

	var bks []Book
	for result.Next() {

		var bk Book

		err = result.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	return bks, nil

}
