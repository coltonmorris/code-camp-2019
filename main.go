package main

import (
	"os"
)

func init() {
	os.Setenv("YOUTUBE_API_KEY", "AIzaSyB6zO_vHL4r3bjzjoBGEDpirBOa_ozoRkM")
	os.Setenv("YOUTUBE_CLIENT_SECRET_JSON", "{\"web\":{\"client_id\":\"151708764487-rp0lbvppvfudv9p24miqv6lm6jf2o3kt.apps.googleusercontent.com\",\"project_id\":\"coce-camp-2019\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"qR-UNuoMgVcdxFkR4njlB9PZ\",\"redirect_uris\":[\"http://synclist.tech/youtube_callback\",\"http://localhost:8080/youtube_callback\",\"http://www.synclist.tech/youtube_callback\"],\"javascript_origins\":[\"http://synclist.tech\",\"http://www.synclist.tech\",\"http://localhost:8080\",\"http://localhost:8080/callback/youtube\"]}}")

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
