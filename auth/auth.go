package auth

import (
	"encoding/base64"
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
)

type Auth struct {
	clientId     string
	clientSecret string
	accessToken  string
	//	redirectURI        string
}

const (
	SPOTIFY_SEARCH_API = "https://api.spotify.com/v1/search"
	SPOTIFY_TOKEN_URI  = "https://accounts.spotify.com/api/token"
)

var clientId string
var clientSecret string
var token string

func New(id string, secret string, user string) Auth {
	auth := Auth{
		clientId:     id,
		clientSecret: secret,
	}
	clientId = id
	clientSecret = secret
	return auth
}

func Authorize() string {

	data := fmt.Sprintf("%v:%v", clientId, clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	authcode := fmt.Sprintf("Basic %s", encoded)

	// create a new request to get our access_token
	// and send our Keys on Authorization Header
	request := gorequest.New()
	request.Post(SPOTIFY_TOKEN_URI)
	request.Set("Authorization", authcode)
	request.Send("grant_type=client_credentials")

	_, body, errs := request.End()

	if errs != nil {
		fmt.Printf("something went wrong")
	}

	// Parse response to simplejson object
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("[Authorize] Error parsing Json!")
		errs = []error{err}
	}

	// check whether we got the access_token or not.
	jsToken, exists := js.CheckGet("access_token")
	if exists {
		// If we got it then assign it to the auth object.
		token, err = jsToken.String()
		if err != nil {
			fmt.Println("[Authorize] Error Getting Access Token from Json!")
			errs = []error{err}
		}
	}

	return token
}
