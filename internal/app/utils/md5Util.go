package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Encryption(plaintext string) (encryption string) {
	h := md5.New()
	h.Write([]byte(plaintext))
	md5Hash := h.Sum(nil)

	encryption = hex.EncodeToString(md5Hash)
	return encryption
}
