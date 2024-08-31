package crypto

import (
	"crypto"
	"fmt"
)

func HashString(value string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(value))
	bytes := hash.Sum(nil)

	return fmt.Sprintf("%x", bytes)
}
