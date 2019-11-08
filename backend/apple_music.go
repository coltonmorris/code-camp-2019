package main

import "fmt"

type Songs struct {
	artist string
	name   string
	album  string
}

type Playlist []Songs

type MusicService interface {
	GetPlaylist() Playlist
	CreatePlaylist(Playlist)
}

type AppleMusic struct {
}

func (a *AppleMusic) GetPlaylist() Playlist {
}

func (a *AppleMusic) CreatePlaylist() {
}
