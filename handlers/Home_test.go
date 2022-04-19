package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/mrpiggy97/golang-jwt-example/handlers"
	"github.com/mrpiggy97/golang-jwt-example/server"
	"github.com/mrpiggy97/golang-jwt-example/utils"
)

// testintHomeWithoutAuthorization header will test response
// of endpoint when no authorization header is provided
func testingHomeWithoutAuthoriationHeader(testCase *testing.T) {
	// start testing server
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(1)
	go server.Runserver()
	//give time for server to be up
	time.Sleep(time.Second * 1)
	//start request
	request, _ := http.NewRequest(
		"GET",
		"http://localhost:8000/api/home",
		nil,
	)
	var client http.Client = http.Client{}
	response, responseError := client.Do(request)
	waiter.Done()
	if responseError != nil {
		testCase.Error(responseError.Error())
	} else {
		decodedBody, _ := io.ReadAll(response.Body)
		fmt.Println(string(decodedBody))
	}
	waiter.Wait()
	defer response.Body.Close()
}

// testingHomeWithExpiredToken will test response with
// with authorization header but with its token expired
func testingHomeWithExpiredJwtToken(testCase *testing.T) {
	//set up testing server
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(2)
	go server.Runserver()
	// give time for server to be up
	time.Sleep(time.Second * 1)
	// first get token from /api/auth/token endpoint
	var responsesChannel chan utils.ResponseWrapper = make(chan utils.ResponseWrapper, 1)
	go utils.MakeRequestToCreateJwtToken(waiter, responsesChannel)
	var createTokenResponse utils.ResponseWrapper = <-responsesChannel
	decodedBody, _ := io.ReadAll(createTokenResponse.Res.Body)
	var tokenResponse *utils.DecodedResponse = new(utils.DecodedResponse)
	_ = json.Unmarshal(decodedBody, tokenResponse)
	// now that we have token we wait for it to expire, in this case 5 seconds
	time.Sleep(time.Second * 6)
	// now we make request to home endpoint with expired token
	newRequest, _ := http.NewRequest(
		"GET",
		"http://localhost:8000/api/home",
		nil,
	)
	newRequest.Header.Add("Authorization", tokenResponse.Token)
	var client http.Client = http.Client{}
	response, responseError := client.Do(newRequest)
	waiter.Done()
	if responseError != nil {
		testCase.Error(responseError.Error())
	} else {
		var expectedResult string = "request is not authenticated"
		decodedBody, _ := io.ReadAll(response.Body)
		fmt.Println(response.Status)
		var homeRes *handlers.HomeResponse = new(handlers.HomeResponse)
		_ = json.Unmarshal(decodedBody, homeRes)
		if homeRes.Message != expectedResult {
			testCase.Errorf("expected result to be %v, instead got %v", expectedResult, homeRes.Message)
		}
	}
	waiter.Wait()
	defer createTokenResponse.Res.Body.Close()
	defer response.Body.Close()
}

// testingHomeWithValidToken will test endpoint response
// with authorization header that has a valid token
func testingHomeWithValidToken(testCase *testing.T) {
	// run testing server
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(2)
	go server.Runserver()
	//give time for server to be up
	time.Sleep(time.Second * 1)
	// get jwt token
	var responsesChannel chan utils.ResponseWrapper = make(chan utils.ResponseWrapper, 1)
	go utils.MakeRequestToCreateJwtToken(waiter, responsesChannel)
	var responseWrapper utils.ResponseWrapper = <-responsesChannel
	decodedBody, _ := io.ReadAll(responseWrapper.Res.Body)
	var tokenResponse *utils.DecodedResponse = new(utils.DecodedResponse)
	_ = json.Unmarshal(decodedBody, tokenResponse)
	// now we make request to home with token in Authorization header
	newRequest, _ := http.NewRequest(
		"GET",
		"http://localhost:8000/api/home",
		nil,
	)
	newRequest.Header.Add("Authorization", tokenResponse.Token)
	var client http.Client = http.Client{}
	response, responseErr := client.Do(newRequest)
	waiter.Done()
	if responseErr != nil {
		testCase.Error(responseErr.Error())
	} else {
		if response.StatusCode != http.StatusAccepted {
			testCase.Errorf(
				"status expected to be %v, got %v instead",
				http.StatusAccepted, response.StatusCode,
			)
		}
		var expectedResult string = "request is authenticated"
		decodedBody, _ := io.ReadAll(response.Body)
		var unmarshaledResponse *handlers.HomeResponse = new(handlers.HomeResponse)
		_ = json.Unmarshal(decodedBody, unmarshaledResponse)
		fmt.Println(string(decodedBody))

		if unmarshaledResponse.Message != expectedResult {
			testCase.Errorf(
				"expected response in body to be %v, go %v instead",
				expectedResult, unmarshaledResponse.Message,
			)
		}
	}
	waiter.Wait()
	defer response.Body.Close()
	defer responseWrapper.Res.Body.Close()
}

func TestHomeHandler(testCase *testing.T) {
	testCase.Run("action=test-home-without-authorization-header", testingHomeWithoutAuthoriationHeader)
	testCase.Run("action=test-home-with-expired-token", testingHomeWithExpiredJwtToken)
	testCase.Run("action=test-home-with-valid-token", testingHomeWithValidToken)
}
