
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

// Generic function args
//
func CLI(fileType string) cli.Command { 

 Command := cli.Command{
    Name:        fileType,
    Usage:       fileType,
    Description: "example blah",
    Subcommands: []cli.Command{
        {
            Name:  "export",
            Usage: "Export from the DB",
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
            Usage: "Import into the DB",
            Flags: []cli.Flag{
                FileLoc,
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.List(fileType)

            }, 
        },
        {
            Name:  "delete",
            Usage: "Delete from the DB",
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
            Usage: "List all in the DB",
            Flags: []cli.Flag{
                DBLoc,
            },
            Action: func(c *cli.Context) {

               raf := NewRafikiInit(c)
               raf.List(fileType)
               
            },
        },
    },
 }

 return Command

}
