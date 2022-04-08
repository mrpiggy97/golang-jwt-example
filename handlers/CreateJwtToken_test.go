package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/mrpiggy97/golang-jwt-example/server"
)

type NewUserRequest struct {
	Username string
	Email    string
	Name     string
}

func TestCreateJwtToken(testCase *testing.T) {
	var testingServer http.Server = server.GetServer()
	go server.Runserver(testingServer)
	var username *string = new(string)
	var email *string = new(string)
	var name *string = new(string)
	fmt.Println("enter Username:")
	fmt.Scanln(username)
	fmt.Println("enter email:")
	fmt.Scanln(email)
	fmt.Println("enter name:")
	fmt.Scanln(name)
	var newUser NewUserRequest = NewUserRequest{
		Username: *username,
		Email:    *email,
		Name:     *name,
	}
	jsonNewUser, _ := json.Marshal(newUser)
	var buferrer *bytes.Buffer = bytes.NewBuffer(jsonNewUser)
	request, reqError := http.NewRequest(
		"POST",
		"http://localhost:8000/api/auth/token",
		buferrer,
	)
	var client http.Client = http.Client{}
	if reqError != nil {
		testCase.Error(reqError)
	}
	res, resError := client.Do(request)
	if resError != nil {
		testCase.Error(resError)
	}
	decodedBody, _ := io.ReadAll(res.Body)
	fmt.Println(string(decodedBody))
	defer testingServer.Close()
}
