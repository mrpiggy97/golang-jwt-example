package server

import (
	"fmt"
	"net/http"
)

func Runserver(server http.Server) {
	fmt.Printf("server listening at address %v\n", server.Addr)
	server.ListenAndServe()
}
