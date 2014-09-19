package rafiki

import (
	"crypto/x509"
	//"encoding/hex"
	"encoding/pem"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
)

func ExportCSR(c *cli.Context) {

	log.Print("csr export")

	checkDB(c.String("db"))
	conn := createDBConn(c.String("db"))
	defer conn.Close()

	key := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	keyname := GetKeyName()

	ciphertext := SelectKey(conn, keyname)

	cleartext, err := DecryptString(key, ciphertext)
	log.Print(string(cleartext))
	ErrHandler(err)

}

func ImportCSR(c *cli.Context) {

	checkDB(c.String("db"))
	conn := createDBConn(c.String("db"))
	defer conn.Close()

	err := CheckFileFlag(c)
	ErrHandler(err)

	buf, err := ioutil.ReadFile(c.String("f"))
	ErrHandler(err)

	block, _ := pem.Decode(buf)

	CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
	ErrHandler(err)

	CSRName := CertificateRequest.Subject

	log.Print(CSRName.CommonName)
	//log.Print(CertificateRequest.SignatureAlgorithm)

	//log.Print(string(hex.Dump(CertificateRequest.Signature)))
	//log.Print(string(hex.EncodeToString(CertificateRequest.Signature)))

	key := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	ciphertext, err := EncryptString(key, block.Bytes)
	ErrHandler(err)

	InsertKey(conn, string(CSRName.CommonName), string(ciphertext))

}

func DeleteCSR(c *cli.Context) {

	log.Print("csr delete")

}

func ListCSR(c *cli.Context) {

	log.Print("csr list")
	checkDB(c.String("db"))
	conn := createDBConn(c.String("db"))
	defer conn.Close()

	ListKeys(conn)

}

func checkCSRFileSet(value bool) {
	if value == false {
		log.Print("No CSR file specified")
		os.Exit(1)
	}

	//if _, err := os.Stat(filename); os.IsNotExist(err) {
	//     log.Print("File not found")
	//     os.Exit(1)
	// }

}
