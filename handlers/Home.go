package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/golang-jwt-example/middleware"
)

type HomeResponse struct {
	Message string
}

// Home handler was merely created to test AuthenticationMiddleware
func Home(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var key middleware.IsAuthenticated = "isAuthenticated"
	var isAuthenticated interface{} = req.Context().Value(key)
	var data HomeResponse = HomeResponse{}

	if isAuthenticated == true {
		data.Message = "request is authenticated"
	} else {
		data.Message = "request is not authenticated"
	}
	jsonData, _ := json.Marshal(data)
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(jsonData)
}
