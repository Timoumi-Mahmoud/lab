package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

// define server Routes
func Routes() {
	http.HandleFunc("/user/get", SendJSON)
	http.HandleFunc("/ping", PingHandler)
}

func SendJSON(w http.ResponseWriter, r *http.Request) {
	u := User{
		"mahmoudXXXX",
		"mahmoud@gmail.com",
	}

	j, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(j)

}
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong ......\n"))
}
