
package rafiki


import (
        "log"
        "errors"
        )
        


func startUp() ([]byte, error) {

    log.Print("Starting")

    password, err:= checkPassword()
    if err != nil {
        return nil, errors.New("Wrong Password")
    }

    key := []byte(password)

    return key, nil

}

