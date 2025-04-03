package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Propriety struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	// PublishAt   time.Time `json:""`
}

type ProprietyList []Propriety

var JSON = `
		{
				"id": 0,
				"title": "Apartment in Russia",
				"price": 1000.25	,
				"location": "Russia, Moscow",
				"description": "lorem ipsum ........................"
		}
	`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %v\n ", r.URL.Path[1:])
}

func GetProperityHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal("Unable to open data.json file", err)
	}
	defer file.Close()
	output := make([]byte, 1024)
	data, err := file.Read(output)
	if err != nil {
		log.Fatal("Unable to Read data.json file", err)
	}
	fmt.Println("file content:", string(output[:data]))

	var ps ProprietyList

	//	err = json.NewDecoder(string(output[:data])).Decode(&ps)

	err = json.Unmarshal(output[:data], &ps)
	if err != nil {
		log.Fatal("err in Decode: :", err)
	}

	/***********************************************/
	var p Propriety
	err = json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		log.Fatal("err in Unmarshal: :", err)
	}

	fmt.Fprintf(w, " %v\n ", ps)
}

func main() {

	log.Println("Starting the server on 6969")
	http.HandleFunc("/", handler)
	http.HandleFunc("GET /propriety", GetProperityHandler)
	log.Fatal(http.ListenAndServe(":6969", nil))

}
