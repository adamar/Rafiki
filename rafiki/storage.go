package rafiki

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


func CreateDB() {

	// Create Database File
	//
	db, err := sql.Open("sqlite3", "./rafiki.db")
	ErrHandler(err)
	defer db.Close()

	// Generate Schema for DB
	//
	sqlStmt := `
        create table csrs (id integer not null primary key,
                          cn text,
                          csr blob);
        `
	_, err = db.Exec(sqlStmt)
	ErrHandler(err)

}

func listKeys(db *sql.DB) {

	rows, err := db.Query("select id, cn, csr from csrs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	rows.Close()
}

func checkDB() {

	// Check for SQLite Database, if unfound prompt to create
	//CreateDB()

	// Open DB Conn
	//
	//db, err := sql.Open("sqlite3", *dbFile)
	//errHandler(err)
	//defer db.Close()

	//listKeys(db)
}

func insertKeys(db *sql.DB) {

	_, err := db.Exec("INSERT INTO csrs (cn, csr) VALUES ('foo', 'asxasxas')")
	ErrHandler(err)

}
