package play

import (
	"fmt"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/parnurzeal/gorequest"
	"github.com/rsbear/spotify/auth"
	"github.com/tidwall/gjson"
)

func PlayNow(searchType string, search ...string) string {
	name := strings.Join(search[:], " ")
	fmt.Println(fmt.Sprintf("Searching for %v%s%v", searchType, "... ", name))
	token := auth.Authorize()
	authcode := fmt.Sprintf("Bearer %s", token)

	SEARCH := fmt.Sprintf("%s%s%v%s%v%s", "https://api.spotify.com/v1/search?", "q=", search, "&type=", searchType, "&limit=20")
	request := gorequest.New()
	request.Get(SEARCH)
	request.Set("Accept", "application/json")
	request.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Set("Authorization", authcode)

	_, body, errs := request.End()

	if errs != nil {
		fmt.Printf("something went wrong")
	}

	searchLine := fmt.Sprintf("%v%s%v%s", searchType, "s.items.#(name =", name, ").uri")
	result := gjson.Get(body, searchLine)
	stringResult := fmt.Sprintf("play track \"%v%v", result, "\"")
	mack.Tell("Spotify", stringResult)

	return "play"
}
