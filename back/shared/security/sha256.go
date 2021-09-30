package security

import (
	"crypto/sha512"
	"fmt"
)

// HashString takes a string in parameter and returns the same string hashed with sha512
func HashString(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	bytesHash := h.Sum(nil)
	hexString := fmt.Sprintf("%x", bytesHash)

	return hexString
}
