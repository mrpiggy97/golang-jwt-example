package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/golang-jwt-example/utils"
)

func VerifyToken(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	token, parsingErr := jwt.Parse(req.Header.Get("Authorization"), utils.TokenVerification)
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
