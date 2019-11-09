package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/zmb3/spotify"
)

var SpotifyRedirectURI string = "http://localhost:8080/callback/spotify"

type API struct {
	sync.RWMutex
	Users map[string]*LameUser
}

type LameUser struct {
	ServiceAccounts map[string]ServiceAccount
}

func (this *API) GetUser(username string) (*LameUser, bool) {
	this.RLock()
	result, ok := this.Users[username]
	this.RUnlock()
	return result, ok
}

// LoginUser has some unique functionality. We decided to allow anyone to login with any username. First check if the user exists, if it doesn't, create it.
func (this *API) LoginUser(username string) {
	if _, exists := this.GetUser(username); exists == true {
		// the user already exists, ideally we have stronger account security :P
		return
	}

	this.Lock()
	this.Users[username] = &LameUser{
		ServiceAccounts: make(map[string]ServiceAccount),
	}
	this.Unlock()

	return
}

func (this *API) RegisterYoutube(user string) (string, error) {
	// ys, err := NewYoutubeService(context.Background())
	// if err != nil {
	// 	return "", err
	// }

	// return ys.authUrl, nil
	return "", nil
}

func (this *API) RegisterSpotify(user string) string {
	// if this env var is available, use it instead
	redirect := os.Getenv("SPOTIFY_REDIRECT")
	if redirect != "" {
		SpotifyRedirectURI = redirect
		fmt.Println("using spotify redirect: ", SpotifyRedirectURI)
	}

	auth := spotify.NewAuthenticator(SpotifyRedirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
	return auth.AuthURL(user)
}

func (this *API) callbackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["service"] == "spotify" {
		auth := spotify.NewAuthenticator(SpotifyRedirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
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
