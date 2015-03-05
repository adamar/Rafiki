package rafiki

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/hex"
	"encoding/pem"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const (
	SSLCERT = iota
	SSLCSR  = iota
	SSLKEY  = iota
	SSHKEY  = iota
	ECPKEY  = iota
)

type Rafiki struct {
	RequireAuth bool
	FileLoc     string
	Password    string
	DB          *sql.DB
}

type Key struct {
	Type         int
	FileContents []byte
	ParsedKey    interface{}
}

func NewRafikiKey(buf []byte) *Key {

	block, _ := pem.Decode(buf)

	// SSL Certificate
	sslcert, err := x509.ParseCertificate(block.Bytes)
	if err == nil {
		return &Key{Type: SSLCERT, FileContents: block.Bytes, ParsedKey: sslcert}
	}

	// SSL Certificate Signing Request
	sslcsr, err := x509.ParseCertificateRequest(block.Bytes)
	if err == nil {
		return &Key{Type: SSLCSR, FileContents: block.Bytes, ParsedKey: sslcsr}
	}

	// SSL Private Key
	sslkey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		return &Key{Type: SSLKEY, FileContents: block.Bytes, ParsedKey: sslkey}
	}

	// RSA Private Key
	sshkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return &Key{Type: SSHKEY, FileContents: block.Bytes, ParsedKey: sshkey}
	}

	// EC Private Key
	ecpkey, err := x509.ParseECPrivateKey(block.Bytes)
	if err == nil {
		return &Key{Type: ECPKEY, FileContents: block.Bytes, ParsedKey: ecpkey}
	}

	return &Key{}

}

func NewRafikiInit(c *cli.Context, checkAuth bool) (raf *Rafiki) {

	var filePath string

	if c.String("f") != "" {
		filePath = c.String("f")
	}

	dbPath := ".rafiki.db"

	if os.Getenv("HOME") != "" {
		dbPath = os.Getenv("HOME") + "/" + dbPath
	}

	if c.String("db") != "" {
		dbPath = c.String("db")
	}

	db, _ := InitDB(dbPath)

	password, _ := InitPassword(db)

	raf = &Rafiki{
		RequireAuth: checkAuth,
		FileLoc:     filePath,
		Password:    password,
		DB:          db,
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

	var commonName, keyType string

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

	PrintOrange("Imported " + fileName)

}

func (raf *Rafiki) Delete() {

	newReader := bufio.NewReader(os.Stdin)
	PrintOrange("Please enter the Key ID to Delete:")
	kId, _ := newReader.ReadString('\n')
	DeleteKey(raf.DB, kId)
	ClearScreen()
	PrintOrange(" Deleted key " + kId)

}

func (raf *Rafiki) List() {

	ClearScreen()
	PrintOrange("\n Key list \n")
	err := ListKeys(raf.DB, "")
	if err != nil {
		log.Print(err)
	}

}

func (raf *Rafiki) Export() {

	err := ListKeys(raf.DB, "")

	keyname := GetKeyName()

	ciphertext, filename := SelectKey(raf.DB, keyname)

	cleartext, err := DecryptString([]byte(raf.Password), ciphertext)
	err = ioutil.WriteFile(filename, []byte(cleartext), 0644)

	if err != nil {
		panic(err)
	}

	PrintOrange("Exported " + filename)

}

func (raf *Rafiki) Profile() {

	log.Print("Not implemented yet")

}

func calcThumbprint(input []byte) string {

	prefix := "Modulus="
	suffix := "\n"
	modulus := strings.ToUpper(hex.EncodeToString(input))
	return formatMd5(md5String(prefix + modulus + suffix))

}
