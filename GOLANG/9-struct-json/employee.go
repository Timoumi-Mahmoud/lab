package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	var e Employee
	fmt.Println("vim-go")
	c := map[string]*Employee{}
	fmt.Printf("%T %+[1]v \n", e)

	bara := Employee{Name: "bara", Number: 112, Hired: time.Now()}
	fmt.Printf("Bara: %+v \n", bara)

	salim := Employee{Name: "salim", Number: 666, Boss: &bara, Hired: time.Now()}
	fmt.Printf("salim: %+v \n", salim)

	j1, err := json.Marshal(salim)
	if err != nil {
		log.Println("error marshaling json: ", err)
	}
	fmt.Println(string(j1))

	start := time.Date(2023, 03, 25, 12, 0, 0, 0, time.UTC)
	fmt.Println("start: ", start)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("bara: ", bara.Name)

	c["lamine"] = &Employee{Name: "Lamine", Number: 69, Hired: time.Now()}

	c["lamine"].Number++
	c[salim.Name] = &salim
	fmt.Println("company lamine: ", c["lamine"])
	fmt.Println("company bara: ", c["salim"])

}
