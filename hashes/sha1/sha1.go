package sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encrypt(content string) string {

	message := sha1.Sum([]byte(content))
	hash := hex.EncodeToString(message[:])
	return hash
}
