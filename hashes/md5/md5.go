package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Encrypt(content string) string {

	message := md5.Sum([]byte(content))
	hash := hex.EncodeToString(message[:])
	return hash
}
