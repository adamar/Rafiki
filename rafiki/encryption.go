package rafiki

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
    "encoding/hex"
    "crypto/sha256"
)



func EncryptString(key []byte, clearText string, blockType string) (string, error) {

    encBuf := bytes.NewBuffer(nil)
    w, err := armor.Encode(encBuf, blockType, nil)
    if err != nil {
        log.Fatal(err)
    }

    plaintext, err := openpgp.SymmetricallyEncrypt(w, key, nil, nil)
    if err != nil {
        return "", err
    }
    message := []byte(clearText)
    _, err = plaintext.Write(message)

    plaintext.Close()
    w.Close()

    return encBuf.String(), nil

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


func shaString(originalString string) (string, error) {

   shastring := sha256.New()
   shastring.Write([]byte(originalString))
   outputString := shastring.Sum(nil)
   outputStringHex := hex.EncodeToString(outputString)
   // make proper error return
   return outputStringHex, nil

}


