package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/zmb3/spotify"
)

type API struct {
	sync.RWMutex
	Users              map[string]*LameUser
	SpotifyRedirectURI string
	YoutubeRedirectURI string
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
	this.LoginUser(user)
	lameUser, ok := this.GetUser(user)
	if !ok {
		return "", fmt.Errorf("could not get user when registering youtube")
	}

	ys, err := NewYoutubeService(context.Background(), user, this.YoutubeRedirectURI)
	if err != nil {
		return "", err
	}

	// add youtube service to user
	this.Lock()
	lameUser.ServiceAccounts["youtube"] = ys
	this.Unlock()

	return ys.authUrl, nil
}

func (this *API) YoutubeAuthCallback(r *http.Request) error {
	code := r.FormValue("code")
	user := r.FormValue("state")

	fmt.Println("youtube callback data: ")
	fmt.Println("code: ", code)
	fmt.Println("user: ", user)

	lameUser, ok := this.GetUser(user)
	if !ok {
		return fmt.Errorf("ERROR: could not find user: %s", user)
	}

	this.Lock()
	ys := lameUser.ServiceAccounts["youtube"].(*YoutubeService)
	if err := ys.Authenticate(code); err != nil {
		this.Unlock()
		return err
	}
	this.Unlock()

	return nil
}

func (this *API) RegisterSpotify(user string) string {
	auth := spotify.NewAuthenticator(this.SpotifyRedirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)
	return auth.AuthURL(user)
}

func (this *API) SpotifyAuthCallback(r *http.Request) error {
	auth := spotify.NewAuthenticator(this.SpotifyRedirectURI, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistModifyPublic)

	// we save our user in this parameter
	user := r.FormValue("state")
	tok, err := auth.Token(user, r)
	if err != nil {
		return err
	}

	client := auth.NewClient(tok)

	this.LoginUser(user)

	this.Lock()
	this.Users[user].ServiceAccounts["spotify"] = NewSpotifyUser(&client)
	this.Unlock()

	return nil
}
