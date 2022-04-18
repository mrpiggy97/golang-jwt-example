package server

import (
	"fmt"
	"net/http"
)

func Runserver() {
	var mainServer http.Server = GetServer()
	fmt.Printf("server listening at address %v\n", mainServer.Addr)
	mainServer.ListenAndServe()
}
