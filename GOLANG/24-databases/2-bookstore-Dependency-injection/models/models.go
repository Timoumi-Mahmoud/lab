package models

import (
	"database/sql"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// conncetion pool as a parameter
func AllBooks(db *sql.DB) ([]Book, error) {
	result, err := db.Query("select * from book")
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
