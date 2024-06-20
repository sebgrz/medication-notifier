package crypto

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "lorem_ipsum"

type TokenError int

const (
	TOKEN_EXPIRED TokenError = iota
	TOKEN_INVALID
	OTHER
)

func (err TokenError) Error() string {
	return strconv.Itoa(int(err))
}

func GenereteToken(userId string, durationInMinutes time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * durationInMinutes)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	authToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return authToken, nil
}

func ValidateTokenAndReturnUserId(authToken string) (string, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return "", TOKEN_EXPIRED
		}

		return "", OTHER
	}

	if !token.Valid {
		return "", TOKEN_INVALID
	}

	subject, _ := token.Claims.GetSubject()
	return subject, nil
}
