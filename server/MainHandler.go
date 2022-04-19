package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/golang-jwt-example/middleware"
)

type MainHandler struct {
	Router *httprouter.Router
}

func (handlerInstance *MainHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	req, middlewareError := middleware.AuthenticationMiddleware(writer, req)
	if middlewareError != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		handlerInstance.Router.ServeHTTP(writer, req)
	}
}
