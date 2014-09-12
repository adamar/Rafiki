package rafiki

import (
	"log"
	"os"
)

func ErrHandler(err error) {
	if err != nil {
		log.Print(err)
	}
}

func noFileSet() {

	log.Print("no file set")
	os.Exit(1)

}
