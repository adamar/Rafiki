package rafiki

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
    "github.com/bndr/gotabulate"
	"log"
	"os"
    "crypto/sha256"
    "encoding/hex"
    "github.com/codegangsta/cli"
)

var db *sql.DB


func InitDB(c *cli.Context) *sql.DB {

    fname := c.String("db")

    if _, err := os.Stat(fname); os.IsNotExist(err) {
        log.Print("db doesnt exit")
        PromptToCreateDB()
    }
    
    db := createDBConn(fname)

    return db

}


func PromptToCreateDB() {

	msg := "No DB Specified, Y/N to create a new one"
	var i string
	fmt.Println(msg)
	fmt.Scan(&i)
	if i == "y" {
		CreateDB()
	} else {
		os.Exit(1)
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

	// Generate Schema for DB
	//
	sqlStmt := `
        create table files (id integer not null primary key,
                          identity text,
                          type text,
                          file blob);
        `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	// Create password table
        //
	sqlStmt = `CREATE TABLE password (
                hashed_password text NOT NULL);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil

}


func ListKeys(db *sql.DB) error {

    new := [][]string{}

    rows, err := db.Query("select id, identity from files")

    if err != nil {
        return err
    }

    for rows.Next() {
        var cn string
        var id string
        rows.Scan(&cn, &id)
        new = append(new, []string{cn,id})
    }
    rows.Close()

    tabulate := gotabulate.Create(new)
    tabulate.SetHeaders([]string{"ID", "CommonName"})

    if len(new) > 0 {
        fmt.Println(tabulate.Render("grid"))
    } else {
        fmt.Println("No Keys to Print\n")
    }


    return nil
}


func checkDB(fname string) (password string, err error) {

	if _, err := os.Stat(fname); os.IsNotExist(err) {
		log.Print("db doesnt exit")
		PromptToCreateDB()
		//password, err := setPassword()
		return password, err
	} else {
		//password, err := checkPassword()
		return password, err
	}
	return password, nil

}

func InsertKey(db *sql.DB, identity string, fileContents string) error {

    fileType := "csr"

	_, err := db.Exec("INSERT INTO files (identity, type, file) VALUES (?, ?, ?)", identity, fileType, fileContents)
	if err != nil {
		return err
	}

	return nil

}


func DeleteKey(db *sql.DB, kId string) error {

    _, err := db.Exec("DELETE FROM files WHERE id = ?", kId)
    if err != nil {
        return err
    }

    return nil

}


func SelectKey(db *sql.DB, id string) string {

	rows, err := db.Query("SELECT file from files WHERE id = ? ", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

    var csr string = ""
	for rows.Next() {
		rows.Scan(&csr)
	}

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

	return pass, nil

}

func CheckIsPasswordSet(db *sql.DB) (string, error) {

    var count string

    err := db.QueryRow("SELECT COUNT(hashed_password) from password").Scan(&count)
    if err != nil {
        log.Fatal(err)
    }
    //defer rows.Close()

    //var pass string
    //for rows.Next() {
    //    rows.Scan(&pass)
    //}

    return count, nil

}


func createDBConn(fname string) *sql.DB {

	db, err := sql.Open("sqlite3", fname)
        if err !=  nil {
            log.Print(err)
        }

	return db
}



func hashStringToSHA256(input string) string {

       hash := sha256.New()
       hash.Write([]byte(input))
       chkSum := hash.Sum(nil)
       return hex.EncodeToString(chkSum)

}

