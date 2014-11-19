package rafiki

import (
	"errors"
	"github.com/codegangsta/cli"
	"os"
)


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
