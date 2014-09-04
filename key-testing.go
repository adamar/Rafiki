package main

import (
	"crypto/x509"
        "crypto/x509/pkix"
	"encoding/pem"
        //"encoding/base64"
	"flag"
	"fmt"
        "log"
	"io/ioutil"
	"os"
)

func main() {

	var (
		//pem      = flag.String("pem", "cert.pem", "X509 Pem File")
		//rsakey = flag.String("rsakey", "id_rsa", "RSA Private Key")
		//crt    = flag.String("crt", "cert.crt", "X509 CRT File")
                csr    = flag.String("csr", "cert.csr", "X509 CSR File")
	)

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	//if crt != nil {
	//	crtVerify(*crt)
	//}

	//if rsakey != nil {
	//	checkRSAKey(*rsakey)
	//}

        if csr != nil {
                CertificateRequest := checkCSR(*csr)
                Printr(CertificateRequest.Subject)
        }

}


func Printr(ssd pkix.Name) {

    log.Print(ssd.)


}


func checkRSAKey(fname string) {

	buf, err := ioutil.ReadFile(fname)
	errHandler(err)
        block, _ := pem.Decode(buf)

        file, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

        fmt.Println(file)

}


func checkCSR(fname string) x509.CertificateRequest {

        buf, err := ioutil.ReadFile(fname)
        errHandler(err)

        block, _ := pem.Decode(buf)

        CertificateRequest, err := x509.ParseCertificateRequest(block.Bytes) //Requires Go 1.3+
        errHandler(err)

        return *CertificateRequest

}




//func readPrivateKey(path string) {
    //file, _ := os.Open(path)

    //stat, err := file.Stat()

    //buf := make([]byte, stat.Size())
    //file.Read(buf)


    //block, _ := pem.Decode(buf)
    //new := x509.ParsePKCS1PrivateKey(block.Bytes)

    //fmt.Println(new)

//}



func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getPrivateKey(fname string) { //*rsa.PrivateKey {
	buf, err := ioutil.ReadFile(fname)
	errHandler(err)

	block, _ := pem.Decode(buf)
	if block == nil {
		fmt.Println("Problem with your key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	errHandler(err)

	err = key.Validate()
	errHandler(err)

	//return key
}

func crtVerify(file string) {

	certPEM, _ := ioutil.ReadFile(file)
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	fmt.Println(cert.DNSNames)
	fmt.Println(cert.EmailAddresses)
	fmt.Println(cert.IssuingCertificateURL)
	fmt.Println(cert.OCSPServer)

	//if _, err := cert.Verify(x509.VerifyOptions{}); err != nil {
	//        panic("failed to verify certificate: " + err.Error())
	//}

}


func csrVerify(file string) {


        certPEM, _ := ioutil.ReadFile(file)
        block, _ := pem.Decode([]byte(certPEM))
        if block == nil {
                panic("failed to parse certificate PEM")
        }

        log.Print(block)
        //newi, err := x509.ParseCertificateRequest(block.Bytes)

        log.Print("Need Go Version 1.3")


}
