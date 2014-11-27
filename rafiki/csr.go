package rafiki

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
)



func checkCSRFileSet(value bool) {

	if value == false {
		log.Print("No CSR file specified")
		os.Exit(1)
	}

}
