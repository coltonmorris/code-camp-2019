package main

import (
	"os"
)

func init() {
	os.Setenv("SPOTIFY_ID", "462eace056d94eaaa00f678b93d9bd0d")
	os.Setenv("SPOTIFY_SECRET", "4b6e16b7fa2349938b9f58f1c685b269")
}

func main() {
	spotifyRedirectURI := os.Getenv("SPOTIFY_REDIRECT")
	if spotifyRedirectURI == "" {
		spotifyRedirectURI = "http://localhost:8080/callback/spotify"
	}

	api := &API{
		Users:              make(map[string]*LameUser, 0),
		SpotifyRedirectURI: spotifyRedirectURI,
		// TODO make sure youtube is getting its redirectURI correctly
	}

	RunHttpServer(api)
}
