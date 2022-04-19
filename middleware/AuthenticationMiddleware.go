package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mrpiggy97/golang-jwt-example/utils"
)

func AuthenticationMiddleware(writer http.ResponseWriter, req *http.Request) (*http.Request, error) {
	var authorizationHeader string = req.Header.Get("Authorization")
	var authenticationStatus IsAuthenticated = "isAuthenticated"
	// if length of Authorization header is equal to 0
	// that means no authorization header was provided so we
	// then return a context with isAuthenticated value set to false
	if len(authorizationHeader) == 0 {
		var newContext context.Context = context.WithValue(
			req.Context(),
			authenticationStatus,
			false,
		)
		req = req.Clone(newContext)
		return req, nil
	}

	token, parsingError := jwt.Parse(req.Header.Get("Authorization"), utils.TokenVerification)
	if parsingError != nil && parsingError.Error() != "Token is expired" {
		return nil, parsingError
	}

	if token.Valid {
		var newContext context.Context = context.WithValue(
			req.Context(),
			authenticationStatus,
			true,
		)
		req = req.Clone(newContext)
	} else {
		var newContext context.Context = context.WithValue(
			req.Context(),
			authenticationStatus,
			false,
		)
		req = req.Clone(newContext)
	}
	return req, nil
}
