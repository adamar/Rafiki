package rafiki

import (
	"os"
	"testing"
)

func TestCreateDBConn(t *testing.T) {

	dbFileName := "test_rafiki_db_name"

	db := createDBConn(dbFileName)
	defer db.Close()

	var now string
	err := db.QueryRow("select 1").Scan(&now)
	if err != nil {
		t.Error(err)
	}

	if now != "1" {
		t.Error("SQLite Error")
	}

	os.Remove(dbFileName)

}
