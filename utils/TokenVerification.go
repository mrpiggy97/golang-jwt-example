package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func TokenVerification(token *jwt.Token) (interface{}, error) {
	// check if signing method is of interface jwt.SigningMethodHMAC
	// this is called type assertation and works only for interfaces
	_, signingMethodIsCorrect := token.Method.(*jwt.SigningMethodHMAC)
	if !signingMethodIsCorrect {
		return nil, errors.New("invallid signing method")
	}
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
