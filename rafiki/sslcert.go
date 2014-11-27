
package rafiki

import (
    "crypto/x509"
    "log"
    "encoding/pem"
    "io/ioutil"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/codegangsta/cli"
    )



func ImportSSLKey(c *cli.Context, db *sql.DB, password string){

    err := CheckFileFlag(c)
        if err != nil {
            log.Print(err)
        }

    buf, err := ioutil.ReadFile(c.String("f"))
        if err != nil {
        log.Print(err)
        }

    block, _ := pem.Decode(buf)

    Certificate, err := x509.ParseCertificate(block.Bytes) //Requires Go 1.3+
        if err != nil {
        log.Print(err)
        }

    CSRName := Certificate.RawSubject
   
    log.Print("rawsubject", CSRName)

    //ciphertext, err := EncryptString([]byte(password), string(buf))

    //InsertKey(db, string(CSRName.CommonName), "csr",ciphertext)

}

