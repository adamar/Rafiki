package rafiki

import (
	"errors"
	"io/ioutil"
	"os"
)

// ReadFile takes a files locationm reads the contents
// and returns a byte array of the file
//
func ReadFile(fileLoc string) ([]byte, error) {

	// Check File Flag set
	//
	//if c.IsSet("f") == false {
	//	return nil, errors.New("File Flag not set")
	//}

	// Check File exists
	//
	if _, err := os.Stat(fileLoc); os.IsNotExist(err) {
		return nil, err
	}

	// Open file and Read Contents
	//
	buf, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		return nil, err
	}

	return buf, nil

}

// CheckFileFlag checks if the file flag is set
//
func CheckFileFlag(fileLoc string) error {

	// Check File Flag set
	//
	if fileLoc == "" {
		return errors.New("File Flag not set")
	}

	// Check File exists
	//
	if _, err := os.Stat(fileLoc); os.IsNotExist(err) {
		return err
	}

	return nil

}
