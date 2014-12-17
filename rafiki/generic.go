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
    ParsedKey                    interface{}
}


func NewRafikiKey(buf []byte) *Key {

    block, _ := pem.Decode(buf)

    // SSL Certificate
    sslcert, err := x509.ParseCertificate(block.Bytes); if err == nil {
       return &Key{Type: SSLCERT, FileContents: block.Bytes, ParsedKey: sslcert}
    }

    // SSL Certificate Signing Request
    sslcsr, err := x509.ParseCertificateRequest(block.Bytes); if err == nil {
       return &Key{Type: SSLCSR, FileContents: block.Bytes, ParsedKey: sslcsr}
    }

    // SSL Private Key
    sslkey, err := x509.ParsePKCS8PrivateKey(block.Bytes); if err == nil {
       return &Key{Type: SSLKEY, FileContents: block.Bytes, ParsedKey: sslkey}
    }

    // RSA Private Key
    sshkey, err := x509.ParsePKCS1PrivateKey(block.Bytes); if err == nil {
       return &Key{Type: SSHKEY, FileContents: block.Bytes, ParsedKey: sshkey}
    }

    // EC Private Key
    ecpkey, err := x509.ParseECPrivateKey(block.Bytes); if err == nil {
       return &Key{Type: ECPKEY, FileContents: block.Bytes, ParsedKey: ecpkey}
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

	return raf

}


// Generic Import function
//
func (raf *Rafiki) Import() {

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

	var commonName,keyType string

    myKey := NewRafikiKey(buf)

	switch myKey.Type {
	  case SSLCERT:

        sslcert := myKey.ParsedKey.(*x509.Certificate)
		commonName = string(sslcert.Subject.CommonName)
        keyType = "sslcert"


	  case SSLKEY:

        rsakey := myKey.ParsedKey.(*rsa.PrivateKey)
        commonName = calcThumbprint(rsakey.N.Bytes())
        keyType = "sslkey"


	  case SSLCSR:

        sslcsr := myKey.ParsedKey.(*x509.CertificateRequest)
		commonName = string(sslcsr.Subject.CommonName)
        keyType = "sslcsr"


      case SSHKEY:

        sshkey := myKey.ParsedKey.(*rsa.PrivateKey)
        commonName = calcThumbprint(sshkey.N.Bytes())
        keyType = "sshkey"


      case ECPKEY:

        commonName = "ec"
        keyType = "ecpkey"
                
	}

	ciphertext, err := EncryptString([]byte(raf.Password), string(buf))

	InsertKey(raf.DB, commonName, keyType, ciphertext, fileName)

}


func (raf *Rafiki) Delete() {

	newReader := bufio.NewReader(os.Stdin)
	log.Print("Please enter the Key ID to Delete:")
	kId, _ := newReader.ReadString('\n')
	DeleteKey(raf.DB, kId)
	log.Print(kId)

}


func (raf *Rafiki) List() {

    ClearScreen()
	PrintOrange("\n" + strings.Title("Key ") + " list" + "\n")
	err := ListKeys(raf.DB, "csr")
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
