package sqlite3helper

import (
	"log"
	"os"
)

func InitDb() {
	_, err := os.Stat("./shortener.db")
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("No Sqlite3 database found in project root, creating...")
			file, err := os.Create("shortener.db")
			if err != nil {
				log.Fatal(err.Error())
			}
			file.Close()
			log.Printf("Sqlite3 database created")
		}
	} else {
		log.Printf("Exiting sqlite3 database found in project root")
	}
}
