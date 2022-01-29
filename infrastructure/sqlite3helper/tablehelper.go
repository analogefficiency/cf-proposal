package sqlite3helper

import (
	"cf-proposal/common/logservice"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	schemaPath = "./domain/ddl/schema"
	viewPath   = "./domain/ddl/views"
)

var tables = map[string]string{
	/** No foreign key constraints */
	"URL":        filepath.Join(schemaPath, "/URL.sql"),
	"HISTORY":    filepath.Join(schemaPath, "/HISTORY.sql"),
	"STATISTICS": filepath.Join(viewPath, "/statistics.sql"),
}

func getTables() map[string]string {
	return tables
}

func initTables(db *sql.DB, filename string) {

	logservice.LogInfo("Seeding database from ddl")
	for table, filepath := range getTables() {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		logservice.LogInfo(fmt.Sprintf("Executing %s script from %s", table, filepath))
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
