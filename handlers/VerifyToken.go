package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

func tokenVerification(token *jwt.Token) (interface{}, error) {
	// check if signing method is of interface jwt.SigningMethodHMAC
	// this is called type assertation and works only for interfaces
	_, signingMethodIsCorrect := token.Method.(*jwt.SigningMethodHMAC)
	if !signingMethodIsCorrect {
		return nil, errors.New("invallid signing method")
	}
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

func VerifyToken(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	token, parsingErr := jwt.Parse(req.Header.Get("Authorization"), tokenVerification)
	if parsingErr != nil {
		fmt.Println(parsingErr)
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		if token.Valid {
			var data map[string]string = make(map[string]string)
			data["message"] = "token is valid"
			jsonData, _ := json.Marshal(data)
			writer.WriteHeader(http.StatusAccepted)
			writer.Write(jsonData)
		}
	}
}
