package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getAll := getCmd.Bool("all", false, "GET all animes")
	getID := getCmd.String("id", "", "GET Anime by ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addID := addCmd.String("id", "", "Anime ID")
	addName := addCmd.String("name", "", "Anime Name")
	addDescription := addCmd.String("description", "", "Anime Description")
	addImageUrl := addCmd.String("imageUrl", "", "Anime Image url")
	addYear := addCmd.Int("year", 0, "Published At")
	addRate := addCmd.Float64("rate", 0.0, "Anime Rate")

	delCmd := flag.NewFlagSet("del", flag.ExitOnError)
	delID := delCmd.String("id", "", "Anime ID")

	if len(os.Args) < 2 {
		fmt.Println("Expecter 'get' or 'add' or 'del' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		//ValidateAnime(addCmd, addID, addName, addDescription, addImageUrl, addYear, addRate)
		HandleAdd(addCmd, addID, addName, addDescription, addImageUrl, addYear, addRate)
	case "del":
		HandleDelete(delCmd, delID)

	default:
	}

}

func ValidateAnime(addCmd *flag.FlagSet, id *string, name *string, description *string, imageUrl *string, year *int, rate *float64) {
	addCmd.Parse(os.Args[2:])
	if *id == "" || *name == "" || *imageUrl == "" || *year == 0 || *rate == 0.0 || *description == "" {
		fmt.Print("All fields are required for adding an anime")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *string, name *string, description *string, imageUrl *string, year *int, rate *float64) {

	ValidateAnime(addCmd, id, name, description, imageUrl, year, rate)
	anime := anime{
		Id:          *id,
		Name:        *name,
		Year:        *year,
		Description: *description,
		ImageUrl:    *imageUrl,
		Rate:        *rate,
	}

	animes := GetAnime()
	animes = append(animes, anime)
	SaveAnime(animes)

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])
	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all anime")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	if *all {
		animes := GetAnime()
		fmt.Print("ID \t Name \t URL \t Year \t Rate \t Description \n")
		for _, anime := range animes {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \n", anime.Id, anime.Name, anime.ImageUrl, anime.Year, anime.Rate, anime.Description)
		}

		return
	}
	if *id != "" {
		animes := GetAnime()
		fmt.Print("ID \t Name \t URL \t Year \t Rate \t Description \n")
		for _, anime := range animes {
			if *id == anime.Id {
				fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \n", anime.Id, anime.Name, anime.ImageUrl, anime.Year, anime.Rate, anime.Description)
			}
		}

	}
}

func HandleDelete(delCmd *flag.FlagSet, id *string) {
	delCmd.Parse(os.Args[2:])
	if *id == "" {
		fmt.Print("id is required")
		delCmd.PrintDefaults()
		os.Exit(1)
	}
	if *id != "" {
		fmt.Println("handle delete: ", *id)
	}

}

// TODO: Delete anime by id
// TODO: Update an anime
// TODO: Show help
