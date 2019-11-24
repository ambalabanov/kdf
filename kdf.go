package kdf

import "crypto/md5"

// Kdf key-derivation function from original Shadowsocks
// https://github.com/shadowsocks/go-shadowsocks2/blob/master/core/cipher.go
func Kdf(password string, keyLen int) []byte {
	var b, prev []byte
	h := md5.New()
	for len(b) < keyLen {
		h.Write(prev)
		h.Write([]byte(password))
		b = h.Sum(b)
		prev = b[len(b)-h.Size():]
		h.Reset()
	}
	return b[:keyLen]
}
