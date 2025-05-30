package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}
	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing word")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strings.ToUpper(word))

}
func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/upper", upperCaseHandler)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
