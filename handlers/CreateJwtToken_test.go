package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/mrpiggy97/golang-jwt-example/server"
	"github.com/mrpiggy97/golang-jwt-example/utils"
)

func testCreateJwtToken(testCase *testing.T) {
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(2)

	var responsesChannel chan utils.ResponseWrapper = make(chan utils.ResponseWrapper, 1)
	// run test server
	go server.Runserver()
	//give time for server to get up
	time.Sleep(time.Second * 1)
	go utils.MakeRequestToCreateJwtToken(waiter, responsesChannel)
	// receive Response
	var res utils.ResponseWrapper = <-responsesChannel
	// now we can close the testing server
	waiter.Done()
	if res.ResError != nil {
		testCase.Error(res.ResError)
	} else if res.Res.StatusCode != 202 {
		testCase.Error(res.Res.Status)
	} else {
		decodedBody, _ := io.ReadAll(res.Res.Body)
		var tokenResponse *utils.DecodedResponse = new(utils.DecodedResponse)
		_ = json.Unmarshal(decodedBody, tokenResponse)
		fmt.Println(tokenResponse.Token)
	}
	waiter.Wait()
}

func TestRunCreateJwtToken(testCase *testing.T) {
	testCase.Run("action=create-jwt-token", testCreateJwtToken)
}
