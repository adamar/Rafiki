package rafiki

import (
	"bufio"
    "strings"
	"io/ioutil"
	"log"
	"os"
	"crypto/x509"
    "crypto/rsa"
	"encoding/pem"
    "encoding/hex"
	"github.com/codegangsta/cli"
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
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

    err := CheckFileFlag(raf.FileLoc)
    if err != nil {
        log.Print("No --file flag set")
        os.Exit(1)
    }

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

        commonName = calcThumbprint(rsakey.N.Bytes())

	case "csr":

		block, _ := pem.Decode(buf)
		CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
		if err != nil {
			log.Print(err)
		}
		commonName = string(CertificateRequest.Subject.CommonName)



    case "sshkey":
  
        block, _ := pem.Decode(buf)

        switch block.Type {
        case "RSA PRIVATE KEY":
                key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
                log.Print(key.PublicKey.N.Bytes())
                commonName = calcThumbprint(key.PublicKey.N.Bytes())
        case "EC PRIVATE KEY":
                //key, _ := x509.ParseECPrivateKey(block.Bytes)
                //log.Print(key.PublicKey.N.Bytes())
                commonName = "ec"
        //case "DSA PRIVATE KEY":
                //return ParseDSAPrivateKey(block.Bytes)
                
        }


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

func calcThumbprint(input []byte) string {

    prefix := "Modulus="
    suffix := "\n"
    modulus := strings.ToUpper(hex.EncodeToString(input))
    return md5String(prefix + modulus + suffix)

}
