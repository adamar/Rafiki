package rafiki

import (
	"crypto/md5"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/bndr/gotabulate"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDB(dbPath string) (*sql.DB, error) {

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		ClearScreen()
		PromptToCreateDB()
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil

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

	msg := "Which key ID?"
	var i string
	fmt.Println(msg)
	fmt.Scan(&i)

	return i

}

func CreateDB() error {

	// Set the DB Path to a hidden file
	//
	dbPath := "./rafiki.db"
	if os.Getenv("HOME") != "" {
		dbPath = os.Getenv("HOME") + "/.rafiki.db"
	}

	// Create Database File
	//
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// Generate Schema for DB
	//
	sqlStmt := `
        create table files (id integer not null primary key,
                          identity text,
                          type text,
                          filename text,
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

func ListKeys(db *sql.DB, fileType string) error {

	new := [][]string{}

	var query string
	if fileType != "" {
		query = fmt.Sprintf("select id, identity, filename, type from files WHERE type = %s", fileType)
	} else {
		query = fmt.Sprintf("select id, identity, filename, type from files order by type")
	}

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	for rows.Next() {
		var id string
		var identity string
		var filename string
		var ftype string
		rows.Scan(&id, &identity, &filename, &ftype)
		new = append(new, []string{id, identity, filename, ftype})
	}
	rows.Close()

	tabulate := gotabulate.Create(new)
	tabulate.SetHeaders([]string{"ID", "CommonName / Fingerprint", "Filename", "File Type"})

	if len(new) > 0 {
		fmt.Println(tabulate.Render("grid"))
	} else {
		fmt.Println("No Keys to Print")
	}

	return nil
}

//func checkDB(fname string) (password string, err error) {
//
//	if _, err := os.Stat(fname); os.IsNotExist(err) {
//		log.Print("db doesnt exit")
//		PromptToCreateDB()
//		//password, err := setPassword()
//		return password, err
//	} else {
//		//password, err := checkPassword()
//		return password, err
//	}
//
//}

func InsertKey(db *sql.DB, identity string, fileType string, fileContents string, fileName string) error {

	_, err := db.Exec("INSERT INTO files (identity, type, filename, file) VALUES (?, ?, ?, ?)", identity, fileType, fileName, fileContents)
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

func SelectKey(db *sql.DB, id string) (string, string) {

	rows, err := db.Query("SELECT file, filename from files WHERE id = ? ", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var fid string
	var filename string
	for rows.Next() {
		rows.Scan(&fid, &filename)
	}

	return fid, filename

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

	return count, nil

}

//func createDBConn(fname string) *sql.DB {
//
//	db, err := sql.Open("sqlite3", fname)
//	if err != nil {
//		log.Print(err)
//	}
//
//	return db
//}

func hashStringToSHA256(input string) string {

	hash := sha256.New()
	hash.Write([]byte(input))
	chkSum := hash.Sum(nil)
	return hex.EncodeToString(chkSum)

}

func md5String(input string) string {

	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))

}

func formatMd5(input string) string {

	i := 0
	final := ""

	for _, c := range input {

		final = final + string(c)
		i++

		if i == len(input) {
			break
		}

		if i%2 == 0 {
			final = final + ":"
		}
	}

	return final
}
