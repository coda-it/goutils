package hash

import (
	"crypto/sha1"
	"fmt"
)

// EncryptString - transform string into hash
func EncryptString(input string) string {
	val := []byte(input)
	h := sha1.New()
	h.Write(val)

	return fmt.Sprintf("%x", h.Sum(nil))
}
