/*
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zmb3/spotify"
)

func init() {
	os.Setenv("SPOTIFY_ID", "462eace056d94eaaa00f678b93d9bd0d")
	os.Setenv("SPOTIFY_SECRET", "4b6e16b7fa2349938b9f58f1c685b269")
}

type API struct {
	Users map[string]*LameUser
}

type LameUser struct {
	ServiceAccounts map[string]ServiceAccount
}

const redirectURI = "http://localhost:8080/callback/spotify"

func main() {
	// api := &API{
	// 	Users: make(map[string]*LameUser, 0),
	// }

	RunHttpServer()
}

func (this *API) SimpleHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "HELLLO")
}

func (this *API) registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	url := ""
	fmt.Println("register hit", params)
	if params["service"] == "spotify" {
		url = this.registerSpotify(params["username"])
	}

	fmt.Println(url)
	type urlResponse struct {
		url string
	}
	io.WriteString(w, url)
}

func (this *API) registerSpotify(user string) string {
	auth := spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
	return auth.AuthURL(user)
}

func (this *API) callbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["service"] == "spotify" {
		auth := spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
		st := r.FormValue("state")
		tok, err := auth.Token(st, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}
		client := auth.NewClient(tok)
		this.Users[st] = &LameUser{
			ServiceAccounts: make(map[string]ServiceAccount),
		}
		this.Users[st].ServiceAccounts["spotify"] = NewSpotifyUser(&client)
	}

	type response struct {
		ok bool
	}
	json.NewEncoder(w).Encode(response{ok: true})
}
