
package rafiki




func checkPassword() (passwd string, err error){

    pass, err := gopass.GetPass("Please enter your Password:")
    return pass, err

}

