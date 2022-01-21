package sqlite3helper

import (
	"fmt"
	"os"
	"testing"
)

func TestInitDbOk(t *testing.T) {
	dbname := "initDbOk"
	filename := fmt.Sprintf("./%s.db", dbname)
	_, err := os.Stat(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Errorf("File should not exist")
		}
	}

	InitDb(dbname)

	_, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("%s database should exist", filename)
		}
	}

	os.Remove(filename)
}
