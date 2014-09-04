package main

import (
        //"encoding/base64"
	"flag"
        "log"
        "database/sql"
	_ "github.com/mattn/go-sqlite3"
        "fmt"
        "os"
)

func main() {


	var (
                dbFile = flag.String("db", "rafiki.db", "Location of Your Rafiki DB File")
	)

        // Check for SQLite Database, if unfound prompt to create
        //
	flag.Parse()
        log.Print(dbFile)
	if flag.NFlag() == 0 {
           var i string
           fmt.Println("No DB specified, y/n to create a new one")
           fmt.Scan(&i)
           if i == "y" {
              CreateDB()
           } else {
              os.Exit(0)
           }
	}

        // Open DB Conn
        //
        db, err := sql.Open("sqlite3", *dbFile)
        errHandler(err)
        defer db.Close()

        listKeys(db)
}

func insertKeys(db *sql.DB) {

        _, err = db.Exec("INSERT INTO csrs (cn, csr) VALUES ('foo', 'asxasxas')")
        errHandler(err)

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


func errHandler(err error) {
        if err != nil {
                log.Print(err)
        }
}

func CreateDB() {

        // Create Database File
        //
        db, err := sql.Open("sqlite3", "./rafiki.db")
        errHandler(err)
        defer db.Close()


        // Generate Schema for DB
        //
        sqlStmt := `
        create table csrs (id integer not null primary key, 
                          cn text,
                          csr blob);
        `
        _, err = db.Exec(sqlStmt)
        errHandler(err)

}
