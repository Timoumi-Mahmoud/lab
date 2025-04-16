package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	fmt.Println("rolling")
	fmt.Println("drivers: ", sql.Drivers())

	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/myDB?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to establish connection with database: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	result, err := db.Query("select * from product")
	if err != nil {
		log.Fatal("Error while fetching product table: ", err)
	}

	for result.Next() {
		var (
			id    int
			name  string
			price int
		)

		err = result.Scan(&id, &name, &price)
		if err != nil {
			log.Fatal("Unable to parse row: ", err)
		}
		fmt.Printf("id: %d, name: %s , price: %d \n", id, name, price)
	}

	var (
		id    int
		name  string
		price int
	)

	err = db.QueryRow("select * from product where id=1").Scan(&id, &name, &price)
	if err != nil {
		log.Fatal("Error while fetching product by id: ", err)
	}

	fmt.Printf("id: %d, name: %s , price: %d \n", id, name, price)
	/*
	   	products := []struct {
	   		id    int
	   		name  string
	   		price int
	   	}{

	   		{3, "water", 1000},
	   		{4, "coka", 1500},
	   	}

	   stmt, err := db.Prepare("Insert into product values($1,$2,$3)")

	   	if err != nil {
	   		log.Fatal("Unable to prepare statement: ", err)
	   	}

	   	for _, product := range products {
	   		_, err = stmt.Exec(product.id, product.name, product.price)
	   		if err != nil {
	   			log.Fatal("Unable to execute statement: ", err)
	   		}
	   	}
	*/
}
