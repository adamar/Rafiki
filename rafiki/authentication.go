package rafiki

import (
	"code.google.com/p/gopass"
	"fmt"
    "errors"
    "log"
)


func InitPassword() {

        log.Print("initdb")

}


func checkPassword() (pass string, err error) {

	pass, err = gopass.GetPass("Please enter your Password:")

    hashedPassword, err := SelectPassword(db)

    hashedPassAttempt := hashStringToSHA256(pass) 


    if hashedPassword != hashedPassAttempt {

        return "", errors.New("Wrong Password")

    }
	return pass, nil

}

func setPassword() (passwd string, err error) {

	pass_attempt_one, err := gopass.GetPass("Please Enter the Password For your New Database:")
	if err != nil {
		return "", err
	}

	pass_attempt_two, err := gopass.GetPass("Please re-enter your new Password:")
	if err != nil {
		return "", err
	}

	if pass_attempt_one != pass_attempt_two {

		passwd := ""
		err = fmt.Errorf("Sorry, the Passwords you entered dont match")
		return passwd, err

	} else {

        err = InsertPassword(db , pass_attempt_one)
        if err != nil {
            panic(err)
        }
		return pass_attempt_one, err
	}
}
