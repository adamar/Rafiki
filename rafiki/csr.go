package rafiki

import (
	"crypto/x509"
	//"encoding/hex"
	"encoding/pem"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
    "bufio"
)

func ExportCSR(c *cli.Context, db *sql.DB, password string) {

    err := CheckFileFlag(c)
    if err != nil {
        log.Print(err)
    }

	keyname := GetKeyName()
    log.Print(keyname)

	ciphertext := SelectKey(db, keyname)

	cleartext, err := DecryptString([]byte(password), ciphertext)
        err = ioutil.WriteFile(c.String("file"), []byte(cleartext), 0644)
        if err != nil {
            panic(err)
        }

}


func ListCSR(c *cli.Context, db *sql.DB, password string) {

    PrintOrange("List of CSRs")
	err := ListKeys(db, "csr")
    if err != nil {
        log.Print(err)
    }

}


func checkCSRFileSet(value bool) {

	if value == false {
		log.Print("No CSR file specified")
		os.Exit(1)
	}

}
