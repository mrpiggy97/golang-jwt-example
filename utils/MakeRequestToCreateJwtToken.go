package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

func MakeRequestToCreateJwtToken(waiter *sync.WaitGroup, sendResponse chan<- ResponseWrapper) {
	// once server is ready to accept requests we start one
	var newUser NewUserRequest = NewUserRequest{
		Username: "puerquis",
		Email:    "puerquisEmail@email.com",
		Name:     "john adams",
	}
	jsonNewUser, _ := json.Marshal(newUser)
	var buferrer *bytes.Buffer = bytes.NewBuffer(jsonNewUser)
	request, _ := http.NewRequest(
		"POST",
		"http://localhost:8000/api/auth/token",
		buferrer,
	)
	var client http.Client = http.Client{}
	httpRes, httpResError := client.Do(request)
	var resWrapper ResponseWrapper = ResponseWrapper{
		Res:      httpRes,
		ResError: httpResError,
	}
	// once we get response from server we send a ResponseWrapper
	// instance through a channel and then we close that channel
	sendResponse <- resWrapper
	close(sendResponse)
	defer waiter.Done()
}
