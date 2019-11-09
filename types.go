package main

type Song struct {
	Artist string
	Album  string
	Name   string
}

type PlayCount struct {
	playlistName string
	songCoung    int
}

type PlaylistId string
type Playlists map[PlaylistId]Playlist

type Playlist struct {
	name      string
	songs     []Song
	songCount int
}

type SyncedSongs struct {
	AcceptedSongs []Song
	FailedSongs   []Song
}

type ServiceAccount interface {
	GetName() string
	GetPlaylist(string) ([]Song, error)
	CreatePlaylist(string, string) (string, error)
	AddSongs(string, []Song) (SyncedSongs, error)
}

type LameUser struct {
	ServiceAccounts map[string]ServiceAccount
}
