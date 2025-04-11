package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Anime struct {
	Title       string `redis:"title"`
	Year        int    `redis:"year"`
	MyRate      string `redis:"myRate"`
	NbrEpisodes int    `redis:"nbrEpisodes"`
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()

	_, err = conn.Do(
		"HMSET",
		"anime:1",
		"title",
		"Berzerk",
		"year",
		1996,
		"nbrEpisodes",
		36,
		"myRate",
		"10/10",
	)
	checkError(err)
	title, err := redis.String(conn.Do("HGET", "anime:1", "title"))
	checkError(err)
	year, err := redis.String(conn.Do("HGET", "anime:1", "year"))
	checkError(err)

	fmt.Println("title: ", title)
	fmt.Println("year: ", year)
	fmt.Println()

	values, err := redis.StringMap(conn.Do("HGETALL", "anime:1"))
	checkError(err)

	for k, v := range values {
		fmt.Printf("key:%v  ,value: %v\n ", k, v)
	}

	fmt.Println()
	fmt.Println()

	reply, err := redis.Values(conn.Do("HGETALL", "anime:1"))
	checkError(err)

	var anime Anime
	err = redis.ScanStruct(reply, &anime)
	checkError(err)

	fmt.Printf("Anime: %v\n", anime)
}
