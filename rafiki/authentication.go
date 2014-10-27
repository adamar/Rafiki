package rafiki

import (
	"code.google.com/p/gopass"
	"fmt"
)

func checkPassword() (passwd string, err error) {

	pass, err := gopass.GetPass("Please enter your Password:")
	return pass, err

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
		err = fmt.Errorf("Passwords dont match")
		return passwd, err

	} else {

		return pass_attempt_one, err
	}
}
