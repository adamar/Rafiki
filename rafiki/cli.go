package rafiki

import (
	"github.com/codegangsta/cli"
)

// Generic File Flag
//
var FileLoc = cli.StringFlag{
	Name:  "f, file",
	Usage: "Location of the file",
}

// Generic DB location Flag
//
var DBLoc = cli.StringFlag{
	Name:  "db, database",
	Value: "rafiki.db",
	Usage: "Location of the DB file",
}

// GenericCLI lays out the available commands
//
var GenericCLI = []cli.Command{
	{
		Name:      "list",
		ShortName: "l",
		Usage:     "List all Keys stored in Rafiki",
		Flags: []cli.Flag{
			FileLoc,
			DBLoc,
		},
		Action: func(c *cli.Context) {
			raf := NewRafikiInit(c)
			raf.List()
		},
	},
	{
		Name:      "delete",
		ShortName: "d",
		Usage:     "List all Keys stored in Rafiki",
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
		Name:      "import",
		ShortName: "i",
		Usage:     "Import a Key into Rafiki",
		Flags: []cli.Flag{
			FileLoc,
			DBLoc,
		},
		Action: func(c *cli.Context) {
			raf := NewRafikiInit(c)
			raf.Import()
		},
	},
	{
		Name:      "export",
		ShortName: "e",
		Usage:     "Export a Key out of Rafiki",
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
		Name:      "profile",
		ShortName: "p",
		Usage:     "Profile a Key",
		Flags: []cli.Flag{
			FileLoc,
			DBLoc,
		},
		Action: func(c *cli.Context) {

			raf := NewRafikiInit(c)
			raf.Profile()

		},
	},
}
