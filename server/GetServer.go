package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/golang-jwt-example/handlers"
)

func GetServer() http.Server {
	var router *httprouter.Router = httprouter.New()
	router.POST("/api/auth/token", handlers.CreateJwtToken)
	router.GET("/api/auth/verify", handlers.VerifyToken)
	var address string = "0.0.0.0:8000"
	return http.Server{Addr: address, Handler: router}
}
