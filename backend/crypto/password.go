package crypto

import (
	"encoding/base64"
	"strconv"

	"golang.org/x/crypto/argon2"
)

func ComparePasswordWithHashedPassword(username, rawPassword, hashedPassword string, accountCreationTime int) bool {
	rawPasswordHashed := generatePassword(rawPassword, username+strconv.Itoa(accountCreationTime))
	return rawPasswordHashed == hashedPassword
}

func GeneratePasswordHash(password, username string, creationTime int) string {
	return generatePassword(password, username+strconv.Itoa(creationTime))
}

func generatePassword(password, salt string) string {
	return base64.RawStdEncoding.EncodeToString(argon2.IDKey([]byte(password), []byte(salt), 4, 64*1024, 4, 32))
}
