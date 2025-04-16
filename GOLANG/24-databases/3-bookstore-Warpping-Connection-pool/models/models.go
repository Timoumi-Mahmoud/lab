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

// Create a custom BookModel type which warps the sql.DB connection pool.
type BookModel struct {
	DB *sql.DB
}

// Use a method on the custom BookModel type to run the sql Query
func (m BookModel) AllBooks() ([]Book, error) {
	result, err := m.DB.Query("select * from book")
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
