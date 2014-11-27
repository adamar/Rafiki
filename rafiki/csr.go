package rafiki

import (
    _ "github.com/mattn/go-sqlite3"
	"github.com/codegangsta/cli"
)


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

               raf := NewRafikiInit(c)
               raf.Export()

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

               raf := NewRafikiInit(c)
               raf.List("csr")

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

               raf := NewRafikiInit(c)
               raf.Delete()

            },
        },
        {
            Name:  "list",
            Usage: "List all CSRs in the DB",
            Flags: []cli.Flag{
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.List("csr")
               
            },
        },
    },
}

