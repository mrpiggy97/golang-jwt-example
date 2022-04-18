package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/mrpiggy97/golang-jwt-example/server"
	"github.com/mrpiggy97/golang-jwt-example/utils"
)

func testingVerifyToken(testCase *testing.T) {
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(2)
	// run testing server
	go server.Runserver()
	// give time for server to get up
	time.Sleep(time.Second * 1)
	// make http request
	var responseChannel chan utils.ResponseWrapper = make(chan utils.ResponseWrapper, 1)
	go utils.MakeRequestToCreateJwtToken(waiter, responseChannel)
	// recieve ResponseWrapper instance
	var response utils.ResponseWrapper = <-responseChannel
	if response.ResError != nil {
		testCase.Error(response.ResError)
	} else {
		// decode response body to get token
		decodedBody, _ := io.ReadAll(response.Res.Body)
		var tokenResponse *utils.DecodedResponse = new(utils.DecodedResponse)
		_ = json.Unmarshal(decodedBody, tokenResponse)
		// make http request to /api/auth/verify
		var url *string = new(string)
		*url = "http://localhost:8000/api/auth/verify"
		newRequest, _ := http.NewRequest(
			"GET",
			"http://localhost:8000/api/auth/verify",
			nil,
		)
		newRequest.Header.Add("Authorization", tokenResponse.Token)
		var client http.Client = http.Client{}
		res, resError := client.Do(newRequest)
		// terminate testing server
		waiter.Done()
		if resError != nil {
			testCase.Error(resError)
		} else if res.StatusCode != 202 {
			testCase.Error(res.Status)
		} else {
			decodedBody, _ := io.ReadAll(res.Body)
			fmt.Println(string(decodedBody))
		}
	}
}

func TestVerifyToken(testCase *testing.T) {
	testCase.Run("action=verify-jwt-token", testingVerifyToken)
}
