package main

import (
	"net/http"

	"github.com/mrpiggy97/golang-jwt-example/server"
)

func main() {
	var mainServer http.Server = server.GetServer()
	server.Runserver(mainServer)
}
