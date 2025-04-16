package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"berserk.com/m/models"
	_ "github.com/lib/pq"
)

type Env struct {
	books models.BookModel
}

func main() {
	var err error
	fmt.Println("rolling")
	fmt.Println("drivers: ", sql.Drivers())

	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/books?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to establish connection with database: ", err)
	}

	// Initalise Env with models.BookModel instance (which in turn warps the connection pool).
	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksList)
	http.ListenAndServe(":6969", nil)
}
func (env *Env) booksList(w http.ResponseWriter, r *http.Request) {
	bks, err := env.books.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
