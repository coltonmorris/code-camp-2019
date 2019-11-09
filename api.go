package main

import (
	"net/http"
	"sync"

	"github.com/zmb3/spotify"
)

type API struct {
	sync.RWMutex
	Users              map[string]*LameUser
	SpotifyRedirectURI string
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
	return "", nil

	// TODO uncomment when blake pushes.
	// lameUser, exists := this.GetUser(user)
	// if exists == false {
	// 	return "", fmt.Errorf("user does not exist")
	// }

	// TODO Add user to NewYoutubeService. Also move env vars to main.go
	// ys, err := NewYoutubeService(context.Background(), user)
	// if err != nil {
	// 	return "", err
	// }

	// // add youtube service to user
	// this.Lock()
	// lameUser.ServiceAccounts["youtube"] = ys
	// this.Unlock()

	// return ys.authUrl, nil
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
