
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


