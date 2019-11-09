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
	"flag"
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

const redirectURI = "http://localhost:8080/callback/spotify"

func main() {
	// heroku creates this env var automagically
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("p", "8080", "port to serve on"))
	}
	flag.Parse()

	directory := "./build"

	// api := &API{
	// 	Users: make(map[string]*LameUser, 0),
	// }

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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

type API struct {
	Users map[string]*LameUser
}

type LameUser struct {
	ServiceAccounts map[string]ServiceAccount
}
