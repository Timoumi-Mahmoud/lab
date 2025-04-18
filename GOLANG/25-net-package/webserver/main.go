package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"
)

type anime struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Year        int     `json:"year"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Rate        float64 `json:"rate"`
}

func checkError(err error) {
	if err != nil {

		log.Fatal(err)
	}
}

func GetAnime() (animes []anime) {
	fileByte, err := ioutil.ReadFile("./anime.json")
	checkError(err)

	// unmarshal convert []byte ==> go type []anime
	// marshal convert go type []anime ==> []byte

	err = json.Unmarshal(fileByte, &animes)
	if err != nil {
		log.Fatal(err)
	}

	return animes
}

func animeHandler(w http.ResponseWriter, r *http.Request) {
	/*
	   animes := GetAnime()
	   response, err := json.Marshal(animes)
	   checkError(err)
	*/
	fileByte, err := ioutil.ReadFile("./anime.json")
	checkError(err)

	w.Write(fileByte)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func main() {

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/anime", animeHandler)
	mux.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
}

/*
- Webserver : is a program that listens on a port for HTTP requests and responds accordingly.
- mux : muliplexer it can handle request for multiple resources paths while routing the request to the appropriate handler func.
it's passed as second arg to ListenAndServe, so then determines wich func to call based on the requested resource path.


Client ---GET /anime --> Go webserver ----> mux ("/anime"=>animeHandler) ---> animeHandler

- be aware of paths that ends with /


*/
