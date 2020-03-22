# spotify-cli

Manage Spotify playback through the command line written in Go

### Create a Spotify developer app
go to [spotifys developer site](https://developer.spotify.com/dashboard/)
create an a desktop app

### Clone the repo
``git clone https://github.com/rsbear/spotify-cli.git``
plop your client_id, client_secret, and username in a .env like so

```
CLIENT_ID=123123  
CLIENT_SECRET=lafjklsldkj 
USERNAME=goslittleoperator
```

cd into it
then run 

``go install``


#### Ball hard my friends
```
spotify help
spotify play artist <artist name>
spotify play track <track name>
spotify play album <album name>
spotify play list <playlist>
spotify play my <playlist name>

spotify pause
spotify next
spotify previous
spotify replay
spotify toggle shuffle
spotify toggle repeat
spotify quit
```
