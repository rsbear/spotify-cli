package my

import (
	"fmt"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/parnurzeal/gorequest"
	"github.com/rsbear/spotify/auth"
	"github.com/tidwall/gjson"
)

func PlayList(username string, search ...string) string {
	name := strings.Join(search[:], " ")
	fmt.Println(fmt.Sprintf("%s%v", "Searching my playlists for ", name))
	token := auth.Authorize()
	authcode := fmt.Sprintf("Bearer %s", token)

	SEARCH_URI := fmt.Sprintf("%s%v%s", "https://api.spotify.com/v1/users/", username, "/playlists?offset=0&limit=20")
	request := gorequest.New()
	request.Get(SEARCH_URI)
	request.Set("Accept", "application/json")
	request.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Set("Authorization", authcode)

	_, body, errs := request.End()

	if errs != nil {
		fmt.Println(errs)
	}

	filter := fmt.Sprintf("%s%v%s", "items.#(name=", name, ").uri")
	result := gjson.Get(body, filter)

	stringResult := fmt.Sprintf("play track \"%v%v", result, "\"")

	mack.Tell("Spotify", stringResult)

	return "play"
}
