package main

import ()

type Song struct {
	artist string
	name   string
	album  string
}

type Playlist struct {
	Name  string
	Songs []Song
}

type User struct {
	email           string
	gToken          []byte
	serviceAccounts map[string]UserServiceAccount
}

type UserServiceAccount struct {
	token     []byte
	playlists []Playlist
}

type MusicService interface {
	GetPlaylist(string) (Playlist, error)
	FindAddSong(string, Song) bool
	CreatePlaylist(string) (string, error)
	GetName() string
}

type Storage interface{}

type UserService interface {
	GetMusicServiceAccount(string) (MusicService, error)
}

// API responses
type PlaylistsResponse struct {
	serviceAccount string
	playlists      []string
}

type GetAllPlaylistsResponse struct {
	accountPlaylists []PlaylistResponse
}

func main() {
}
