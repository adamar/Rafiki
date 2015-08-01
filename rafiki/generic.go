package rafiki

import (
	"bufio"
	"bytes"
	"code.google.com/p/go.crypto/openpgp"
	"code.google.com/p/go.crypto/ssh"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Rafiki struct {
	RequireAuth bool
	FileLoc     string
	Password    string
	DB          *sql.DB
}

type Key struct {
	Identifier   string
	Type         string
	FileContents []byte
	ParsedKey    interface{}
}

func NewRafikiKey(buf []byte) *Key {

	block := checkPemDecode(buf)

	switch {
	case validCSR(block):
		sslcsr, _ := x509.ParseCertificateRequest(block)
		return &Key{
			Identifier:   string(sslcsr.Subject.CommonName),
			Type:         "sslcsr",
			FileContents: block,
			ParsedKey:    sslcsr,
		}

	case validCert(block):
		sslcert, _ := x509.ParseCertificate(block)
		return &Key{
			Identifier:   string(sslcert.Subject.CommonName),
			Type:         "sslcert",
			FileContents: block,
			ParsedKey:    sslcert,
		}

	case validSSLKey(block):
		sslkey, _ := x509.ParsePKCS8PrivateKey(block)
		return &Key{
			Identifier:   calcThumbprint(sslkey.(*rsa.PrivateKey).N.Bytes()),
			Type:         "sslkey",
			FileContents: block,
			ParsedKey:    sslkey,
		}

	case validRSAKey(block):
		sshkey, _ := x509.ParsePKCS1PrivateKey(block)
		return &Key{
			Identifier:   calcThumbprint(sshkey.N.Bytes()),
			Type:         "sshkey",
			FileContents: block,
			ParsedKey:    sshkey,
		}

	case validECKey(block):
		ecpkey, _ := x509.ParseECPrivateKey(block)
		return &Key{
			Identifier:   "ec", // Require proper identifier here
			Type:         "ecpkey",
			FileContents: block,
			ParsedKey:    ecpkey,
		}

	case validPublicKey(block):
		publickey, _ := ssh.ParsePublicKey(block)
		return &Key{
			Identifier:   calcThumbprint(publickey.Marshal()),
			Type:         "publickey",
			FileContents: block,
			ParsedKey:    publickey,
		}

	case validPGPKey(block):
		keyReader := bytes.NewReader(block)
		keyRing, _ := openpgp.ReadArmoredKeyRing(keyReader)
		return &Key{
			Identifier:   fmt.Sprintf("0x0%x", keyRing[0].PrimaryKey.KeyId),
			Type:         "pgpkey",
			FileContents: block,
			ParsedKey:    keyRing,
		}
	}

	return &Key{}

}

func validCSR(input []byte) bool {

	_, err := x509.ParseCertificateRequest(input)
	if err != nil {
		return false
	}
	return true

}

func checkPemDecode(buf []byte) []byte {

	decoded, contents := pem.Decode(buf)
	if decoded == nil {
		return contents
	}
	return decoded.Bytes

}

func validCert(input []byte) bool {

	_, err := x509.ParseCertificate(input)
	if err != nil {
		return false
	}
	return true

}

func validSSLKey(input []byte) bool {

	_, err := x509.ParsePKCS8PrivateKey(input)
	if err != nil {
		return false
	}
	return true

}

func validRSAKey(input []byte) bool {

	_, err := x509.ParsePKCS1PrivateKey(input)
	if err != nil {
		return false
	}
	return true

}

func validECKey(input []byte) bool {

	_, err := x509.ParseECPrivateKey(input)
	if err != nil {
		return false
	}
	return true

}

func validPGPKey(input []byte) bool {

	keyReader := bytes.NewReader(input)
	_, err := openpgp.ReadArmoredKeyRing(keyReader)
	if err != nil {
		return false
	}
	return true

}

func validPublicKey(input []byte) bool {

	_, err := ssh.ParsePublicKey(input)
	if err != nil {
		return false
	}
	return true

}

func validBotoKey(input []byte) bool {

	_, err := ini.Load(input)
	if err != nil {
		return false
	}
	return true

}

func NewRafikiInit(c *cli.Context, checkAuth bool) (raf *Rafiki) {

	var filePath string
	var password string

	if c.String("f") != "" {
		filePath = c.String("f")
	}

	dbPath := ".rafiki.db"

	if os.Getenv("HOME") != "" {
		dbPath = os.Getenv("HOME") + "/" + dbPath
	}

	if c.IsSet("db") == true {
		dbPath = c.String("db")
	}

	db, _ := InitDB(dbPath)

	if checkAuth == true {
		password, _ = InitPassword(db)
	}

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

	myKey := NewRafikiKey(buf)

	ciphertext, err := EncryptString([]byte(raf.Password), string(buf))

	InsertKey(raf.DB, myKey.Identifier, myKey.Type, ciphertext, fileName)

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

	err := CheckFileFlag(raf.FileLoc)
	if err != nil {
		log.Print("No --file flag set")
		os.Exit(1)
	}

	buf, err := ReadFile(raf.FileLoc)
	if err != nil {
		log.Print(err)
	}

	myKey := NewRafikiKey(buf)

	PrintOrange(raf.FileLoc)
	PrintOrange(myKey.Type)
	PrintOrange(myKey.Identifier)

}

func calcThumbprint(input []byte) string {

	prefix := "Modulus="
	suffix := "\n"
	modulus := strings.ToUpper(hex.EncodeToString(input))
	return formatMd5(md5String(prefix + modulus + suffix))

}
