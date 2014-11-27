package rafiki

import (
	"errors"
	"github.com/codegangsta/cli"
	"os"
    "io/ioutil"
)


func ReadFile(c *cli.Context) ([]byte, error) {

	// Check File Flag set
	//
	if c.IsSet("f") == false {
		return nil, errors.New("File Flag not set")
	}

	// Check File exists
	//
	if _, err := os.Stat(c.String("f")); os.IsNotExist(err) {
		return nil, err
	}

    // Open file and Read Contents
    //
    buf, err := ioutil.ReadFile(c.String("f"))
    if err != nil {
        return nil, err   
    }

	return buf, nil

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