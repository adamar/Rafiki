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
        SSLCommand,
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
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.Export(c, db, password)

            },
		},
		{
			Name:  "import",
			Usage: "Import a CSR into the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.Import(c, db, password, "csr")

            }, 
		},
		{
			Name:  "delete",
			Usage: "Delete a CSR from the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.Delete(c, db, password)
            },
		},
		{
			Name:  "list",
			Usage: "List all CSRs in the DB",
			Flags: []cli.Flag{
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.List(c, db, password, "csr")
               
            },
		},
	},
}

var SSLCommand = cli.Command{
	Name:        "ssl",
	Usage:       "ssl",
	Description: "example SSL blah",
	Subcommands: []cli.Command{
		{
			Name:  "export",
			Usage: "Export a CSR Cert from the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.Export(c, db, password)

            },
		},
		{
			Name:  "import",
			Usage: "Import an SSL Cert into the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.Import(c, db, password, "sslkey")

            }, 
		},
		{
			Name:  "delete",
			Usage: "Delete an SSL Cert from the DB",
			Flags: []cli.Flag{
				FileLoc,
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.Delete(c, db, password)

            },
		},
		{
			Name:  "list",
			Usage: "List all SSL Certs in the DB",
			Flags: []cli.Flag{
				DBLoc,
			},
            Action: func(c *cli.Context) {

               db := rafiki.InitDB(c)
               password, _ := rafiki.InitPassword(db)
               rafiki.List(c, db, password, "sslkey")
               
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
