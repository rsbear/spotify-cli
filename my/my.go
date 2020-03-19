package my

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/parnurzeal/gorequest"
	"github.com/rsbear/spotify/auth"
	"github.com/tidwall/gjson"
)

func PlayList(search string, username string) string {
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

	filter := fmt.Sprintf("%s%v%s", "items.#(name=", search, ").uri")
	result := gjson.Get(body, filter)
	fmt.Println(result)

	playYes := fmt.Sprintf(`"play track \"%v\""`, result)
	fmt.Println(playYes)

	yes, err := mack.Tell("Spotify", playYes)
	// yes, err := mack.Tell("Spotify", "play track \"spotify:playlist:5r0SyLIoCOh4TS3lFE6bdN\"")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yes)

	return "play"
}
