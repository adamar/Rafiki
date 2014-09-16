package rafiki

import (
	"log"
	"os"
    "github.com/codegangsta/cli"
    "errors"
)


func ErrHandler(err error) {
	if err != nil {
		log.Print(err)
        os.Exit(1)
	}
}




func CheckFileFlag(c *cli.Context) error {

    // Check File Flag set
    //
    if c.IsSet("f") == false {
       return errors.New("File Flag not set")
    }

    // Check File exists
    //
    if _, err := os.Stat(c.String("f")); os.IsNotExist(err) {
       return err
    }

    return nil

}



