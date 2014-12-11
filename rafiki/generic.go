package rafiki

import (
	"bufio"
    "strings"
	"io/ioutil"
	"log"
	"os"
    "path"   
	"crypto/x509"
    "crypto/rsa"
	"encoding/pem"
    "encoding/hex"
	"github.com/codegangsta/cli"
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
)


const (
    SSLCERT = iota
    SSLCSR  = iota
    SSLKEY  = iota
    SSHKEY  = iota
    ECPKEY  = iota
)


type Rafiki struct {
    FileLoc  string
    Password string
    DB       *sql.DB
}


type Key struct {
    Type                         int
    FileContents                 []byte
}


func NewRafikiKey(buf []byte) *Key {

    block, _ := pem.Decode(buf)

    // SSL Certificate
    _, err := x509.ParseCertificate(block.Bytes); if err == nil {
       return &Key{Type: SSLCERT, FileContents: block.Bytes}
    }

    // SSL Certificate Signing Request
    _, err = x509.ParseCertificateRequest(block.Bytes); if err == nil {
       return &Key{Type: SSLCSR, FileContents: block.Bytes}
    }

    // SSL Private Key
    _, err = x509.ParsePKCS8PrivateKey(block.Bytes); if err == nil {
       return &Key{Type: SSLKEY, FileContents: block.Bytes}
    }

    // RSA Private Key
    _, err = x509.ParsePKCS1PrivateKey(block.Bytes); if err == nil {
       return &Key{Type: SSHKEY, FileContents: block.Bytes}
    }

    // EC Private Key
    _, err = x509.ParseECPrivateKey(block.Bytes); if err == nil {
       return &Key{Type: ECPKEY, FileContents: block.Bytes}
    }

    return &Key{}

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

    _, fileName := path.Split(raf.FileLoc)

	buf, err := ReadFile(raf.FileLoc)
	if err != nil {
		log.Print(err)
	}

	var commonName string

    myKey := NewRafikiKey(buf)

	switch myKey.Type {
	case SSLCERT:

		Certificate, err := x509.ParseCertificate(block.Bytes) //Requires Go 1.3+
		if err != nil {
			log.Print(err)
		}
		commonName = string(Certificate.Subject.CommonName)


	case SSLKEY:

		block, _ := pem.Decode(buf)

        key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
        if err != nil {
            log.Print(err)
        }

        rsakey := key.(*rsa.PrivateKey)

        commonName = calcThumbprint(rsakey.N.Bytes())

	case SSLCSR:

		block, _ := pem.Decode(buf)
		CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
		if err != nil {
			log.Print(err)
		}
		commonName = string(CertificateRequest.Subject.CommonName)



    case SSHKEY:
  
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

	InsertKey(raf.DB, commonName, rtype, ciphertext, fileName)

}

func (raf *Rafiki) Delete() {

	newReader := bufio.NewReader(os.Stdin)
	log.Print("Please enter the Key ID to Delete:")
	kId, _ := newReader.ReadString('\n')
	DeleteKey(raf.DB, kId)
	log.Print(kId)

}

func (raf *Rafiki) List(rtype string) {

    ClearScreen()
	PrintOrange("\n" + strings.Title(rtype) + " List" + "\n")
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
