package main

import "fmt"

type Customer struct {
	FirstName string
	LastName  string
	Age       int16
}

func getCustomers() (customers []Customer) {

	bara := Customer{Age: 5, FirstName: "bara", LastName: "hazmi"}
	mahmoud := Customer{"mahmoud", "timoumi", 29}
	customers = append(customers, bara)
	customers = append(customers, mahmoud)
	return customers
}

func main() {
	fmt.Println("vim-go inside Docker")
	c := getCustomers()
	fmt.Println(c)
	fmt.Println()
	animes := getAnime()
	for i, _ := range animes {
		animes[i].Description = animes[i].Description + "\n description here"
	}
	fmt.Println(animes)
	saveAnime(animes)

}
