package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type UserRequest struct {
	Username string
	Email    string
	Name     string
}

type UserClaims struct {
	UserRequest
	jwt.RegisteredClaims
}

type NewUserRequest struct {
	Username string
	Email    string
	Name     string
}

type ResponseWrapper struct {
	Res      *http.Response
	ResError error
}

type DecodedResponse struct {
	Token string
}
