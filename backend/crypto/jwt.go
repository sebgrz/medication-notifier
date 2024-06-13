package crypto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "lorem_ipsum"

func GenereteToken(userId string) (string, error) {
	claims := jwt.RegisteredClaims {
		Subject: userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	authToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return authToken, nil
}
