package rafiki

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func TestCreateDBConn(t *testing.T) {

	dbFileName := "test_rafiki_db_name"

	CreateDB(dbFileName)

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		t.Error(err)
	}

	var now string
	err = db.QueryRow("select 1").Scan(&now)
	if err != nil {
		t.Error(err)
	}

	if now != "1" {
		t.Error("SQLite Error")
	}

	os.Remove(dbFileName)

}
