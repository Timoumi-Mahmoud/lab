package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"berserk.com/m/models"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	fmt.Println("rolling")
	fmt.Println("drivers: ", sql.Drivers())

	models.DB, err = sql.Open("postgres", "postgresql://root:password@localhost:5433/books?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to establish connection with database: ", err)
	}
	http.HandleFunc("/books", booksList)
	http.ListenAndServe(":6969", nil)
}
func booksList(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}

}
