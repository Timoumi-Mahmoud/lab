package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type anime struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Year        int16   `json:"year"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Rate        float64 `json:"rate"`
}

func getAnime() (animes []anime) {
	fileByte, err := ioutil.ReadFile("./anime.json")
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Println("b: ", string(fileByte))
	// unmarshal convert []byte ==> go type []anime
	// marshal convert go type []anime ==> []byte

	err = json.Unmarshal(fileByte, &animes)
	if err != nil {
		log.Fatal(err)
	}

	return animes
	/*
	   	anime1 := anime{
	   		Id:       "ehoijfn",
	   		Name:     "Berzerk",
	   		Year:     1996,
	   		ImageUrl: "https://google/berzerk",
	   		Rate:     9.9,
	   	}

	   	anime2 := anime{
	   		Id:       "ehoijfndkjf",
	   		Name:     "Monster",
	   		Year:     1990,
	   		ImageUrl: "https://google/monster",
	   		Rate:     9.4,
	   	}

	   	anime3 := anime{
	   		Id:       "oedfdoihoijfn",
	   		Name:     "Cowboy bebop",
	   		Year:     1999,
	   		ImageUrl: "https://google/Cowboy",
	   		Rate:     8.7,
	   	}

	   return []anime{anime1, anime2, anime3}
	*/
}

func saveAnime(anime []anime) {

	animeByte, err := json.Marshal(anime)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./anime-update.json", animeByte, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
