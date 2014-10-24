package rafiki

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func CheckCreateDB() {

	msg := "No DB Specified, Y/N to create a new one"
	var i string
	fmt.Println(msg)
	fmt.Scan(&i)
	if i == "y" {
		CreateDB()
	} else {
		os.Exit(0)
	}

}

func GetKeyName() string {

	msg := "Which key?"
	var i string
	fmt.Println(msg)
	fmt.Scan(&i)

	return i

}

func CreateDB() error {

	// Create Database File
	//
	db, err := sql.Open("sqlite3", "./rafiki.db")
	if err != nil {
		return err
	}
	//defer db.Close()

	// Generate Schema for DB
	//
	sqlStmt := `
        create table csrs (id integer not null primary key,
                          cn text,
                          csr blob);
        `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	// Create password table
	sqlStmt = `CREATE TABLE password (
                hashed_password UNSIGNED BIG INT NOT NULL);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil

}

func ListKeys(db *sql.DB) error {

	rows, err := db.Query("select cn from csrs")
	if err != nil {
		return err
	}
	//defer rows.Close()
	for rows.Next() {
		var cn string
		rows.Scan(&cn)
		fmt.Println(cn)
	}
	rows.Close()

	return nil
}

func checkDB(fname string) (password string, err error) {

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		log.Print("db doesnt exit")
		CheckCreateDB()
		password, err := setPassword()
		return password, err
	} else {
		password, err := checkPassword()
		return password, err
	}
	return password, nil

}

func InsertKey(db *sql.DB, cn string, csr string) error {

	_, err := db.Exec("INSERT INTO csrs (cn, csr) VALUES (?, ?)", cn, csr)
	if err != nil {
		return err
	}

	return nil

}

func SelectKey(db *sql.DB, id string) string {

	rows, err := db.Query("SELECT csr from csrs WHERE id = ? ", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

    var csr string = ""
	for rows.Next() {
		rows.Scan(&csr)
	}

	ErrHandler(err)

	return csr

}

func InsertPassword(db *sql.DB, password string) error {

	_, err := db.Exec("INSERT INTO password (hashed_password) VALUES (?)", password)
	if err != nil {
		return err
	}

	return nil

}

func SelectPassword(db *sql.DB) (string, error) {

	rows, err := db.Query("SELECT hashed_password from password LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pass string
	for rows.Next() {
		rows.Scan(&pass)
	}

	ErrHandler(err)

	return pass, nil

}

func createDBConn(fname string) *sql.DB {

	db, err := sql.Open("sqlite3", fname)
	ErrHandler(err)
	//defer db.Close()

	return db
}
