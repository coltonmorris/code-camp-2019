package main

import (
	"os"
)

func init() {
	os.Setenv("YOUTUBE_API_KEY", "AIzaSyAmxriK-yE1-bc1lT8GSaszAUG8SP3J7tY")
	os.Setenv("YOUTUBE_CLIENT_SECRET_JSON", "{\"web\":{\"client_id\":\"612720349908-4qpltha782sgpooira264m51smkpqmr1.apps.googleusercontent.com\",\"project_id\":\"synclist-blake\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"EezHkcB0gFrycMJe0Vl7tn_4\",\"redirect_uris\":[\"http://synclist.tech/callback/youtube\",\"http://localhost:8080/callback/youtube\"],\"javascript_origins\":[\"http://synclist.tech\",\"http://localhost:8080\"]}}")

	os.Setenv("SPOTIFY_ID", "65ab513777a4414483b9b8a7275879d1")
	os.Setenv("SPOTIFY_SECRET", "dce8883dcca14ffba6611a876f313432")
}

func main() {
	spotifyRedirectURI := os.Getenv("SPOTIFY_REDIRECT")
	if spotifyRedirectURI == "" {
		spotifyRedirectURI = "http://localhost:8080/callback/spotify"
	}

	youtubeRedirectURI := os.Getenv("YOUTUBE_REDIRECT")
	if youtubeRedirectURI == "" {
		youtubeRedirectURI = "http://localhost:8080/callback/youtube"
	}

	api := &API{
		Users:              make(map[string]*LameUser, 0),
		SpotifyRedirectURI: spotifyRedirectURI,
		YoutubeRedirectURI: youtubeRedirectURI,
	}

	RunHttpServer(api)
}
