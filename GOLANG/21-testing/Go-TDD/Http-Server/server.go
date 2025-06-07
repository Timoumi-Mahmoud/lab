package main

import (
	"fmt"
	"log"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.PathValue("name")
	fmt.Println("name: ", player)
	fmt.Fprint(w, GetPlayerScore(player))

}

func GetPlayerScore(name string) string {

	fmt.Println("name: ==>", name)
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""

}

func main() {
	log.Println("Starting the server ..............")
	http.HandleFunc("/player/{name}", PlayerServer)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
