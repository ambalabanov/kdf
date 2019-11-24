package kdf

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestKdf(t *testing.T) {
	key := hex.EncodeToString(Kdf("password", 32))
	if key != strings.ToLower("5F4DCC3B5AA765D61D8327DEB882CF992B95990A9151374ABD8FF8C5A7A0FE08") {
		t.Errorf("Key was incorrect")
	}
	key = hex.EncodeToString(Kdf("", 32))
	if key != strings.ToLower("D41D8CD98F00B204E9800998ECF8427E59ADB24EF3CDBE0297F05B395827453F") {
		t.Errorf("Key was incorrect")
	}
}
