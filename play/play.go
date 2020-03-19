package play

import (
	"fmt"

	// "github.com/andybrewer/mack"
	"github.com/parnurzeal/gorequest"
	"github.com/rsbear/spotify/auth"
	"github.com/tidwall/gjson"
)

type Play struct {
	uri string
}

type Obj struct {
	uri string
}

func PlayNow(search string, searchType string) string {
	token := auth.Authorize()
	authcode := fmt.Sprintf("Bearer %s", token)

	SEARCH := fmt.Sprintf("%s%s%v%s%v%s", "https://api.spotify.com/v1/search?", "q=", search, "&type=", searchType, "&limit=30")
	request := gorequest.New()
	request.Get(SEARCH)
	request.Set("Accept", "application/json")
	request.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Set("Authorization", authcode)

	_, body, errs := request.End()

	if errs != nil {
		fmt.Printf("something went wrong")
	}

	filter := fmt.Sprintf("%v%s", searchType, "s")
	// fmt.Println(filter)
	result := gjson.Get(body, filter)
	// mack.Tell("Spotify", "play")
	fmt.Println(result)

	return "play"
}
