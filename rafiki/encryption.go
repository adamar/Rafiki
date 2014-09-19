package rafiki

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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
