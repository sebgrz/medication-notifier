package crypto

import (
	"strconv"

	"golang.org/x/crypto/argon2"
)

func ComparePasswordWithHashedPassword(username, rawPassword, hashedPassword string, accountCreationTime int) bool {
	rawPasswordHashed := generatePassword(rawPassword, username+strconv.Itoa(accountCreationTime))
	return rawPasswordHashed == hashedPassword
}

func generatePassword(password, salt string) string {
	return string(argon2.IDKey([]byte(password), []byte(salt), 4, 64*1024, 4, 32))
}
