package rafiki

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
    "code.google.com/p/gopass"
    "fmt"
)

func EncryptString(key, ClearText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	ErrHandler(err)

	bse64 := base64.StdEncoding.EncodeToString(ClearText)
	CipherText := make([]byte, aes.BlockSize+len(bse64))
	iv := CipherText[:aes.BlockSize]

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(CipherText[aes.BlockSize:], []byte(bse64))

	return CipherText, nil

}

func DecryptString(key, CipherText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	ErrHandler(err)

	iv := CipherText[:aes.BlockSize]
	CipherText = CipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(CipherText, CipherText)

	data, err := base64.StdEncoding.DecodeString(string(CipherText))
	ErrHandler(err)

	return data, nil

}




func setPassword() (passwd string, err error){

    pass_attempt_one, err := gopass.GetPass("Please enter your new Password:")
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




