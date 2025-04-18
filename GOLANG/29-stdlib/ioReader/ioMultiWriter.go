package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Company struct {
	Name string
	ID   int
}

func main() {
	companies := []Company{
		{Name: "Apple", ID: 1},
		{Name: "Googl", ID: 2},
		{Name: "Debian", ID: 3},
		{Name: "Red Hat", ID: 4},
	}

	f, err := os.Create("companies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bf, err := os.Create("companies-backup.json")
	if err != nil {
		log.Fatal(err)
	}
	defer bf.Close()

	w := io.MultiWriter(f, bf, os.Stdout)

	enc := json.NewEncoder(w)
	for _, c := range companies {
		if err := enc.Encode(c); err != nil {
			log.Fatal(err)
		}
	}
	/*
		bf, err := os.Create("companies-backup.json")
		if err != nil {
			log.Fatal(err)
		}
		defer bf.Close()

		enc = json.NewEncoder(bf)
		for _, c := range companies {
			if err := enc.Encode(c); err != nil {
				log.Fatal(err)
			}
		}

	*/
}
