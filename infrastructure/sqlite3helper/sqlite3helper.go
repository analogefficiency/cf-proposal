package sqlite3helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DbConn *sql.DB

func InitDb(dbname string) {
	filename := fmt.Sprintf("./%s.db", dbname)
	buildSchema := false
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("No Sqlite3 database found in project root, creating.")
			file, err := os.Create(filename)
			if err != nil {
				log.Fatal(err.Error())
			}
			file.Close()
			log.Printf("Sqlite3 database created")
			buildSchema = true
		}
	} else {
		log.Printf("Existing sqlite3 database found in project root")
	}

	DbConn, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}

	if buildSchema {
		initTables(DbConn, filename)
	}
	log.Printf("Sqlite3 database ready!")
}
