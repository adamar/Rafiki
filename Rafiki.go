package main

import (
	"./rafiki"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	// CLI parsing is done here
	//
	app := cli.NewApp()
	app.Name = "Rafiki"
	app.Version = "0.0.2"
	app.Usage = "Store SSL Certs and CSRs securely"
	app.Commands = rafiki.GenericCLI

	// Start Application
	//
	app.Run(os.Args)
}
