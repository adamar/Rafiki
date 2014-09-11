package main

import (
	//"encoding/base64"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
        "github.com/codegangsta/cli"
)


var CSRCommand = cli.Command{
        Name:        "csr",
        Usage:       "csr",
        Description: "example CSR blah",
        Subcommands: []cli.Command{
                 {
                 Name: "export",
                 Usage: "Export a CSR from the DB",
                 Flags: []cli.Flag{
                     FileLoc,
                     },
                 Action: doMain,
                  },
                 {
                 Name: "import",
                 Usage: "Import a CSR into the DB",
                 Flags: []cli.Flag{
                     FileLoc,
                     },
                 Action: doMain,
                  },
                 {
                 Name: "delete",
                 Usage: "Delete a CSR from the DB",
                 Flags: []cli.Flag{
                     FileLoc,
                     },
                 Action: doMain,
                  },
                 {
                 Name: "list",
                 Usage: "List all CSRs in the DB",
                 Action: doMain,
                 },
        },
}



var FileLoc = cli.StringFlag{
    Name: "f, file",
    Usage: "Location of the file",
}


var DBLoc = cli.StringFlag{
    Name: "db, database",
    Value: "rafiki.db",
    Usage: "Location of the DB file",
}



func main() {
    app := cli.NewApp()
    app.Name = "Rafiki"
    app.Version = "0.0.1"
    app.Usage = "Store SSL Certs securely-ish"
    app.Action = doMain
    app.Flags = []cli.Flag{
         DBLoc,
    }
    app.Commands = []cli.Command{
          CSRCommand,
    }
    app.Run(os.Args)
}




func doMain(c *cli.Context) {

        log.Print(c.String("db"))

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
