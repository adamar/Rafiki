package rafiki

import (
	"code.google.com/p/gopass"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func InitPassword(db *sql.DB) (string, error) {

	ClearScreen()

	passIsSet := CheckStoredPassword(db)

	var password = ""

	if passIsSet == false {
		password, _ := setPassword(db)
		return password, nil
	} else {
		password, err := checkPassword(db)
		if err != nil {
			PrintOrange("Sorry, your password appears to be incorrect!")
			os.Exit(1)
		}
		return password, nil
	}

	return password, nil

}

// Check if a password is tored in the DB
//
func CheckStoredPassword(db *sql.DB) bool {

	res, _ := CheckIsPasswordSet(db)

	// Should attempt change to int before checking
	if res == "0" {
		return false
	} else {
		return true
	}


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

		hashedPassword := hashStringToSHA256(pass_attempt_one)
		err = InsertPassword(db, hashedPassword)
		if err != nil {
			panic(err)
		}
		return pass_attempt_one, err
	}
}
