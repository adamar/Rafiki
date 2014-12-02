package rafiki

import (
	"bufio"
	"crypto/x509"
    "crypto/rsa"
	"database/sql"
    "strings"
	"encoding/pem"
    "encoding/hex"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
)

type Rafiki struct {
	FileLoc  string
	Password string
	DB       *sql.DB
}

func NewRafikiInit(c *cli.Context) (raf *Rafiki) {

	db := InitDB(c)
	password, _ := InitPassword(db)

	raf = &Rafiki{
		FileLoc:  c.String("f"),
		Password: password,
		DB:       db,
	}

	return

}

// Generic Import function
//
func (raf *Rafiki) Import(rtype string) {

	buf, err := ReadFile(raf.FileLoc)
	if err != nil {
		log.Print(err)
	}

	var commonName string

	switch rtype {
	case "sslcert":

		block, _ := pem.Decode(buf)
		Certificate, err := x509.ParseCertificate(block.Bytes) //Requires Go 1.3+
		if err != nil {
			log.Print(err)
		}
		commonName = string(Certificate.Subject.CommonName)


	case "sslkey":

		block, _ := pem.Decode(buf)

        key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
        if err != nil {
            log.Print(err)
        }

        rsakey := key.(*rsa.PrivateKey)

        prefix := "Modulus="
        suffix := "\n" 
        modulus := strings.ToUpper(hex.EncodeToString(rsakey.N.Bytes()))

        commonName = md5String(prefix + modulus + suffix)



	case "csr":

		block, _ := pem.Decode(buf)
		CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
		if err != nil {
			log.Print(err)
		}
		commonName = string(CertificateRequest.Subject.CommonName)

	}

	ciphertext, err := EncryptString([]byte(raf.Password), string(buf))

	InsertKey(raf.DB, commonName, rtype, ciphertext)

}

func (raf *Rafiki) Delete() {

	newReader := bufio.NewReader(os.Stdin)
	log.Print("Please enter the Key ID to Delete:")
	kId, _ := newReader.ReadString('\n')
	DeleteKey(raf.DB, kId)
	log.Print(kId)

}

func (raf *Rafiki) List(rtype string) {

	PrintOrange(rtype + " List")
	err := ListKeys(raf.DB, rtype)
	if err != nil {
		log.Print(err)
	}

}

func (raf *Rafiki) Export() {

	//err := CheckFileFlag(c)
	//if err != nil {
	//    log.Print(err)
	//}

	keyname := GetKeyName()

	ciphertext := SelectKey(raf.DB, keyname)

	cleartext, err := DecryptString([]byte(raf.Password), ciphertext)
	err = ioutil.WriteFile(raf.FileLoc, []byte(cleartext), 0644)
	if err != nil {
		panic(err)
	}

}
