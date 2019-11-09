package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
func init() {
	os.Setenv("SPOTIFY_ID", "462eace056d94eaaa00f678b93d9bd0d")
	os.Setenv("SPOTIFY_SECRET", "4b6e16b7fa2349938b9f58f1c685b269")
}

const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopePlaylistReadPrivate)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
	ch = make(chan *spotify.Client)
	state = "abc123"

	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	sUser := NewSpotifyUser("bob", client)

	syncedSongs, err := TransferPlaylist("C0d3 k4Mp", "bob", sUser, sUser)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println()
	fmt.Println("ACCEPTED!")
	fmt.Println(syncedSongs.AcceptedSongs)
	fmt.Println("FAILED!")
	fmt.Println(syncedSongs.FailedSongs)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}
