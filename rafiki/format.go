package rafiki

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func hashStringToSHA256(input string) string {

	hash := sha256.New()
	hash.Write([]byte(input))
	chkSum := hash.Sum(nil)
	return hex.EncodeToString(chkSum)

}

func md5String(input string) string {

	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))

}

func formatMd5(input string) string {

	i := 0
	final := ""

	for _, c := range input {

		final = final + string(c)
		i++

		if i == len(input) {
			break
		}

		if i%2 == 0 {
			final = final + ":"
		}
	}

	return final
}
