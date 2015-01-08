
package rafiki


import  (
        "testing"
        "strings"
        )

func TestEncryptString(t *testing.T) {

    key := []byte("My Encryption Key")
    clearText := `My Important Data to Encrypt`
    output, err := EncryptString(key, clearText)
    
    if err != nil {
        t.Error("Encrypt String Failed")
    }

    expected_prefix := `-----BEGIN PGP SIGNATURE-----`

    if strings.HasPrefix(output, expected_prefix)  != true {
        t.Error("Encrypted String doesnt have PGP Signature")
    }

}


func TestDecryptString(t *testing.T) {


    encrypted_string := `-----BEGIN PGP SIGNATURE-----

wx4EBwMCqZVMs/V7IV1gtsRHjTqbn8S3vXyYGd/Yd9HS4AHkzUsIvAqMuowNAwS0
ocirtuGEgOCw4Cnh9Ivg/+KMvOah4N7kvlyAqXsJrVaejXWjmLbHM+D141+imeSf
jRrR4JTiMQpe2eCa5JakEfxBOb2pw1ud0q8g+gjiR+eWGuEfjwA=
=W3OG
-----END PGP SIGNATURE-----`
 
    key := []byte("My Encryption Key")

    clearText := `My Important Data to Encrypt`

    output, err := DecryptString(key, encrypted_string)
    if err != nil {
        t.Error("Decrypt String Failed")
    }

    if output != clearText {
        t.Error("Decrypted String doesnt match expected output")
    }

}
