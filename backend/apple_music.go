package main

import "fmt"

type Songs struct {
	artist string
	name string
	album string
}

type Playlist []Songs

type MusicService interface {
	GetPlaylist() Playlist 
	CreatePlaylist(Playlist) 
}
