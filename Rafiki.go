package main

import (
	"./rafiki"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

    csrCli := rafiki.CLI("csr")
    sslcertCli := rafiki.CLI("sslcert")
    sslkeyCli := rafiki.CLI("sslkey")

    // CLI parsing is done here
    //
    app := cli.NewApp()
    app.Name = "Rafiki"
    app.Version = "0.0.2"
    app.Usage = "Store SSL Certs and CSRs securely"
    app.Commands = []cli.Command{
        csrCli,
        sslcertCli,
        sslkeyCli,
    }

    // Start Application
    //
    app.Run(os.Args)
}

