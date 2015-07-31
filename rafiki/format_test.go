package rafiki

import (
	"testing"
)

func Testmd5String(t *testing.T) {

	input := "test string"
	output := md5String(input)

	if output != "6f8db599de986fab7a21625b7916589c" {
		t.Error("MD5 Formatting failed")
	}

}

func TestformatMd5(t *testing.T) {

	input := "6f8db599de986fab7a21625b7916589c"
	output := formatMd5(input)

	if output != "6f:8d:b5:99:de:98:6f:ab:7a:21:62:5b:79:16:58:9c" {
		t.Error("MD5 Formatting failed")
	}

}
