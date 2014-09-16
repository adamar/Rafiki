package rafiki

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
        "os"
)


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

func checkDB(fname string) {

    if _, err := os.Stat(fname); os.IsNotExist(err) {
       log.Print("db doesnt exit")
       CheckCreateDB()
    }

}

func InsertKeys(db *sql.DB, cn string, csr string) {

	_, err := db.Exec("INSERT INTO csrs (cn, csr) VALUES ('foo', 'asxasxas')")
	ErrHandler(err)

}


func createDBConn(fname string) *sql.DB {

        db, err := sql.Open("sqlite3", fname)
        ErrHandler(err)
        //defer db.Close()

        return db
}


