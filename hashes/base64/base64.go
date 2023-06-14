package base64

import (
	"encoding/base64"
)

func Encode(content string) string {

	hash := base64.StdEncoding.EncodeToString([]byte(content))
	return hash
}
