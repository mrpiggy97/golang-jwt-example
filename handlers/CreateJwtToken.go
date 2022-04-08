package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

type UserRequest struct {
	Username string
	Email    string
	Name     string
}

type UserClaims struct {
	UserRequest
	jwt.RegisteredClaims
}

func CreateJwtToken(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// read request body
	jsonBody, _ := io.ReadAll(req.Body)
	// create UserRequest type instance and populate its fields with
	// json info decoded from body
	var incomingUser *UserRequest = new(UserRequest)
	var decodingErr error = json.Unmarshal(jsonBody, incomingUser)
	if decodingErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	// to create token we first need to createa UserClaims type instance,
	// its fields will come primarly from UserRequest type instance
	var claims UserClaims = UserClaims{
		UserRequest: UserRequest{
			Username: incomingUser.Username,
			Email:    incomingUser.Email,
			Name:     incomingUser.Name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		},
	}

	// we also need a jwt.SigningMethod type instance to create token
	var signingMethod jwt.SigningMethod = jwt.GetSigningMethod(jwt.SigningMethodHS256.Alg())
	// finally we can create a jwt.Token type instance with the signing method,
	// and claims created above
	token := jwt.NewWithClaims(signingMethod, claims)
	// the final step is to sign the token with a key, in this case the
	// key was created ouside this project, keygenerator.io is highly
	// recommended
	stringToken, tokenErr := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if tokenErr != nil {
		panic(tokenErr)
	}

	// now we send token via http request
	var data map[string]string = make(map[string]string)
	data["token"] = stringToken
	jsonResponse, _ := json.Marshal(data)
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(jsonResponse)
	defer req.Body.Close()
}