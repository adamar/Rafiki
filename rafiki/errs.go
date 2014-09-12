package rafiki

import (
	"log"
	"os"
    "github.com/codegangsta/cli"
    "errors"
    "fmt"
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


func CheckCreateDB() {

           msg := "No DB Specified, Y/N to create a new one"
           var i string
           fmt.Println(msg)
           fmt.Scan(&i)
           if i == "y" {
              CreateDB()
           } else {
              os.Exit(0)
           }

}


