package main

import (
	"fmt"
	"log"
	"os"

	// "strings"

	"github.com/andybrewer/mack"
	"github.com/joho/godotenv"
	"github.com/rsbear/spotify/auth"
	"github.com/rsbear/spotify/my"
	"github.com/rsbear/spotify/play"
	"github.com/urfave/cli/v2"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	clientId, exists := os.LookupEnv("CLIENT_ID")
	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	username, exists := os.LookupEnv("USERNAME")
	if exists != true {
		log.Fatal("Check your credentials")
	}
	auth.New(clientId, clientSecret, username)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "play",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "album",
						Usage: "play music by album",
						Action: func(c *cli.Context) error {
							// search := strings.Split()
							search := c.Args().Slice()
							// fmt.Println(search)
							play.PlayNow("album", search...)
							return nil
						},
					},
					{
						Name:  "artist",
						Usage: "play music by artist",
						Action: func(c *cli.Context) error {
							// search := strings.Split()
							search := c.Args().Slice()
							// fmt.Println(search)
							play.PlayNow("artist", search...)
							return nil
						},
					},
					{
						Name:  "playlist",
						Usage: "play by playlist",
						Action: func(c *cli.Context) error {
							search := c.Args().Slice()
							play.PlayNow("playlist", search...)
							return nil
						},
					},
					{
						Name:  "song",
						Usage: "play by song",
						Action: func(c *cli.Context) error {
							search := c.Args().Slice()
							play.PlayNow("track", search...)
							return nil
						},
					},
					{
						Name:  "my",
						Usage: "My playlists",
						Action: func(c *cli.Context) error {
							search := c.Args().Slice()
							my.PlayList(username, search...)
							return nil
						},
					},
				},
			},
			{
				Name:    "pause",
				Aliases: []string{"t"},
				Usage:   "pause Spotify",
				Action: func(c *cli.Context) error {
					mack.Tell("Spotify", "playpause")
					fmt.Println("Pausing Spotify...")
					return nil
				},
			},
			{
				Name:    "next",
				Aliases: []string{"t"},
				Usage:   "Play next track",
				Action: func(c *cli.Context) error {
					mack.Tell("Spotify", "next track")
					fmt.Println("Onto the next track...")
					return nil
				},
			},
			{
				Name:    "previous",
				Aliases: []string{"t"},
				Usage:   "Play previous track",
				Action: func(c *cli.Context) error {
					mack.Tell("Spotify", "previous track")
					fmt.Println("Going back to previous track...")
					return nil
				},
			},
			{
				Name:    "quit",
				Aliases: []string{"t"},
				Usage:   "Play previous track",
				Action: func(c *cli.Context) error {
					mack.Tell("Spotify", "quit")
					fmt.Println("Quitting Spotify")
					return nil
				},
			},
			{
				Name:    "replay",
				Aliases: []string{"t"},
				Usage:   "Replay track",
				Action: func(c *cli.Context) error {
					mack.Tell("Spotify", "set player position to 0")
					fmt.Println("Replaying the current track")
					return nil
				},
			},
			{
				Name:    "toggle",
				Aliases: []string{"t"},
				Usage:   "Toggle shuffle or repeat",
				Action: func(c *cli.Context) error {
					//shuffle
					if c.Args().First() == "shuffle" {
						cur, err := mack.Tell("Spotify", "shuffling")
						if err != nil {
							fmt.Println("sometin wong")
						}
						if cur != "true" {
							fmt.Println("Shuffle on...")
							mack.Tell("Spotify", "set shuffling to not shuffling")
						} else {
							fmt.Println("Shuffle off...")
							mack.Tell("Spotify", "set shuffling to not shuffling")
						}
					}
					//repeat
					if c.Args().First() == "repeat" {
						cur, err := mack.Tell("Spotify", "repeating")
						if err != nil {
							fmt.Println("sometin wong")
						}
						if cur != "true" {
							fmt.Println("Repeat on...")
							mack.Tell("Spotify", "set repeating to not repeating")
						} else {
							fmt.Println("Repeat off...")
							mack.Tell("Spotify", "set repeating to not repeating")
						}
					}
					return nil
				},
			},
			{
				Name:    "status",
				Aliases: []string{"t"},
				Usage:   "Show the artist, track, and album that are playing",
				Action: func(c *cli.Context) error {
					artist, err := mack.Tell("Spotify", "artist of current track as string")
					if err != nil {
						fmt.Println("something wrong")
					}

					album, err := mack.Tell("Spotify", "album of current track as string")
					if err != nil {
						fmt.Println("something wrong")
					}

					track, err := mack.Tell("Spotify", "name of current track as string")
					if err != nil {
						fmt.Println("something wrong")
					}

					if artist != "" && album != "" && track != "" {
						fmt.Println("Artist.. ", artist)
						fmt.Println("Album... ", album)
						fmt.Println("Song.... ", track)
					}

					return nil
				},
			},
		},
	}

	errors := app.Run(os.Args)
	if errors != nil {
		log.Fatal(errors)
	}
}
