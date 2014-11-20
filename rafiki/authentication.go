package rafiki

import (
	"code.google.com/p/gopass"
	"fmt"
    "errors"
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)


func InitPassword(db *sql.DB) (string, error) {

        log.Print("initpassword")

        passIsSet := CheckStoredPassword(db)

        var password = ""

        if passIsSet == false {
            password, _ := setPassword(db)
            return password, nil
        } else {
            password, _ := checkPassword(db)
            return password, nil
        }

        return password, nil

}


// Check if a password is tored in the DB
//
func CheckStoredPassword(db *sql.DB) bool {


    return false


}


// Compare the provided password against the hash in the DB
//
func checkPassword(db *sql.DB) (pass string, err error) {

	pass, err = gopass.GetPass("Please enter your Password:")

    hashedPassword, err := SelectPassword(db)

    hashedPassAttempt := hashStringToSHA256(pass) 


    if hashedPassword != hashedPassAttempt {

        return "", errors.New("Wrong Password")

    }
	return pass, nil

}


// Prompt the user to enter a new password (twice)
//
func setPassword(db *sql.DB) (passwd string, err error) {

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
