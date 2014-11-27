
package rafiki

import (
    "crypto/x509"
    "log"
    "encoding/pem"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/codegangsta/cli"
    )



func ImportSSLKey(c *cli.Context, db *sql.DB, password string){

    buf, err := ReadFile(c)
    if err != nil {
         log.Print(err)
    }

    block, _ := pem.Decode(buf)

    Certificate, err := x509.ParseCertificate(block.Bytes) //Requires Go 1.3+
    if err != nil {
        log.Print(err)
    }

    commonName := string(Certificate.Subject.CommonName)

    ciphertext, err := EncryptString([]byte(password), string(buf))

    InsertKey(db, commonName, "sslkey", ciphertext)

}

