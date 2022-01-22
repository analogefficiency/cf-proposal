package sqlite3helper

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	schemaPath = "./domain/ddl/schema"
)

var tables = map[string]string{
	/** No foreign key constraints */
	"URL": filepath.Join(schemaPath, "/URL.sql"),
}

func getTables() map[string]string {
	return tables
}

func initTables(db *sql.DB, filename string) {

	for table, filepath := range getTables() {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Creating %s table from %s", table, filepath)
		createTable(db, string(content))
	}
}

func createTable(db *sql.DB, s string) {
	statement, err := db.Prepare(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
