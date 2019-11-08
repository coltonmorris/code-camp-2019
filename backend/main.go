package main

type Song struct {
	artist string
	name   string
	album  string
}

type Playlist struct {
	name  string
	songs []song
}

type MusicService interface {
	GetPlaylist() Playlist
	CreatePlaylist(Playlist)
}
