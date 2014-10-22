
package rafiki




func startUp() (string, error) {

    log.Print("Starting")

    checkDB(c.String("db"))
    conn := createDBConn(c.String("db"))
    defer conn.Close()

    password, err:= checkPassword()
    if err != nil {
        log.Print("Password entry failed")
    }

    key := []byte(password)


}

