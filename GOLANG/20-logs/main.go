package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func main() {
	InfoLogger.Println("This is info")
	WarningLogger.Println("This is very important")
	ErrorLogger.Println("Something went wrong")
}

/*

	log.SetFlags(log.LstdFlags | log.Lshortfile)
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)

	log.Println("vim go")
//log.Fatal("Log fatal  ") // log the message and call os.Exit(1)
	//log.Panic("Log panic  ") // log the message and call runtime panic

	// check github.com/sirupsen/logrus => more feature: json, color based on log level, or even create my own formatter.
*/
