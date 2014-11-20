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
	app.Version = "0.0.1"
	app.Usage = "Store SSL Certs securely-ish"
	app.Commands = []cli.Command{
		CSRCommand,
	}
	app.Run(os.Args)
}

var CSRCommand = cli.Command{
	Name:        "csr",
	Usage:       "csr",
	Description: "example CSR blah",
	Subcommands: []cli.Command{
		{
			Name:  "export",
			Usage: "Export a CSR from the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
			Action: rafiki.ExportCSR,
		},
		{
			Name:  "import",
			Usage: "Import a CSR into the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
			Action: rafiki.ImportCSR,
		},
		{
			Name:  "delete",
			Usage: "Delete a CSR from the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
			Action: rafiki.DeleteCSR,
		},
		{
			Name:  "list",
			Usage: "List all CSRs in the DB",
			Flags: []cli.Flag{
				DBLoc,
			},
            Action: func(c *cli.Context) {

               rafiki.InitDB(c)
               rafiki.InitPassword()
               rafiki.ListCSR(c)
               

            },
		},
	},
}

var FileLoc = cli.StringFlag{
	Name:  "f, file",
	Usage: "Location of the file",
}

var DBLoc = cli.StringFlag{
	Name:  "db, database",
	Value: "rafiki.db",
	Usage: "Location of the DB file",
}
