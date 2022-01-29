package sqlite3helper

import (
	"cf-proposal/common/logservice"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DbConn *sql.DB

func InitDb(dbname string) {
	filename := fmt.Sprintf("./sqlite/%s.db", dbname)
	buildSchema := false
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			logservice.LogInfo("No Sqlite3 database found in project root, creating.")
			file, err := os.Create(filename)
			if err != nil {
				log.Fatal(err.Error())
			}
			file.Close()
			logservice.LogInfo("Sqlite3 database created")
			buildSchema = true
		}
	} else {
		logservice.LogInfo("Existing sqlite3 database found in project root")
	}

	DbConn, err = sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}

	if buildSchema {
		initTables(DbConn, filename)
	}
	logservice.LogInfo("Sqlite3 database ready!")
}
