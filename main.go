package main

import (
	"os"
)

func init() {
	// os.Setenv("YOUTUBE_API_KEY", "AIzaSyB6zO_vHL4r3bjzjoBGEDpirBOa_ozoRkM")
	os.Setenv("YOUTUBE_CLIENT_SECRET_JSON", "{\"web\":{\"client_id\":\"966586910350-1nfo5mp4cs44smqag8jb3c04jedfh29c.apps.googleusercontent.com\",\"project_id\":\"mono-207103\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"pvA96vgMDTzxtABoEAixpEen\",\"redirect_uris\":[\"http://synclist.tech/callback/youtube\",\"http://www.synclist.tech/callback/youtube\",\"http://localhost:8080/callback/youtube\"],\"javascript_origins\":[\"http://synclist.tech\",\"http://www.synclist.tech\",\"http://localhost:8080\"]}}")

	// os.Setenv("SPOTIFY_ID", "462eace056d94eaaa00f678b93d9bd0d")
	// os.Setenv("SPOTIFY_SECRET", "4b6e16b7fa2349938b9f58f1c685b269")
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
