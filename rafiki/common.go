package rafiki

import (
	"errors"
	"log"
)

func startUp() ([]byte, error) {

	log.Print("Starting")

	password, err := checkPassword()
	if err != nil {
		return nil, errors.New("Wrong Password")
	}

	key := []byte(password)

	return key, nil

}
