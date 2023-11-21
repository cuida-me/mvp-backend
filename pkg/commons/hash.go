package commons

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func NewJwt(sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "secrets",
		"sub": sub,
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 500).Unix(),
	})

	return token.SignedString([]byte("jfdjsbfjdsbfjhdsbhjfbdsjhbfdsjh"))
}

func ValidateJwt(token string) (string, string, error) {
	tkn, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("jfdjsbfjdsbfjhdsbhjfbdsjhbfdsjh"), nil
	})
	if err != nil {
		return "", "", err
	}

	if tkn.Valid && err == nil {
		claims := tkn.Claims.(jwt.MapClaims)
		sub := claims["sub"].(string)
		split := strings.Split(sub, "_")
		return split[0], split[1], err
	}

	return "", "", err
}

func GenerateToken(length int) string {
	uuidObj, _ := uuid.NewUUID()

	uuidStr := uuidObj.String()

	if length < len(uuidStr) {
		uuidStr = uuidStr[:length]
	}

	return uuidStr
}
