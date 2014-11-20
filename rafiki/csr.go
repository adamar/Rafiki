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
)

func ExportCSR(c *cli.Context) {

	password, err := checkDB(c.String("db"))
        if err != nil {
            log.Print(err)
        }

	conn := createDBConn(c.String("db"))
	defer conn.Close()

        err = CheckFileFlag(c)
        if err != nil {
            log.Print(err)
        }

	keyname := GetKeyName()
        log.Print(keyname)

	ciphertext := SelectKey(conn, keyname)

	cleartext, err := DecryptString([]byte(password), ciphertext)
        err = ioutil.WriteFile(c.String("file"), []byte(cleartext), 0644)
        if err != nil {
            panic(err)
        }

}

func ImportCSR(c *cli.Context) {

	password, _ := checkDB(c.String("db"))
    log.Print(password)
	conn := createDBConn(c.String("db"))
	defer conn.Close()

	err := CheckFileFlag(c)
        if err != nil {
            log.Print(err)
        }

	buf, err := ioutil.ReadFile(c.String("f"))
        if err != nil {
	    log.Print(err)
        }

	block, _ := pem.Decode(buf)

	CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
        if err != nil {
	    log.Print(err)
        }

	CSRName := CertificateRequest.Subject

	log.Print(CSRName.CommonName)

	ciphertext, err := EncryptString([]byte(password), string(buf))

        log.Print(ciphertext)
	InsertKey(conn, string(CSRName.CommonName), ciphertext)

}

func DeleteCSR(c *cli.Context) {

	log.Print("csr delete")

}

func ListCSR(c *cli.Context, db *sql.DB) {

	log.Print("csr list")
	//checkDB(c.String("db"))
	//conn := createDBConn(c.String("db"))
	//defer conn.Close()

	ListKeys(db)

}

func checkCSRFileSet(value bool) {

	if value == false {
		log.Print("No CSR file specified")
		os.Exit(1)
	}

}
