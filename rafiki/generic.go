
package rafiki

import (
    "crypto/x509"
    "log"
    "encoding/pem"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/codegangsta/cli"
    )



func Import(c *cli.Context, db *sql.DB, password string, Rtype string){

    buf, err := ReadFile(c)
    if err != nil {
         log.Print(err)
    }

    var commonName string

    switch Rtype {
        case "sslkey":
            
            block, _ := pem.Decode(buf)
            Certificate, err := x509.ParseCertificate(block.Bytes) //Requires Go 1.3+
            if err != nil {
                log.Print(err)
            }
            commonName = string(Certificate.Subject.CommonName)

        case "csr":

            block, _ := pem.Decode(buf)
            CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
            if err != nil {
                log.Print(err)
            }
            commonName = string(CertificateRequest.Subject.CommonName)

        }

    ciphertext, err := EncryptString([]byte(password), string(buf))

    InsertKey(db, commonName, Rtype, ciphertext)

}




func Delete(c *cli.Context, db *sql.DB, password string) {
    
    newReader := bufio.NewReader(os.Stdin)
    log.Print("Please enter the Key ID to Delete:")
    kId, _ := newReader.ReadString('\n')
    DeleteKey(db, kId)
    log.Print(kId)
                        
}

