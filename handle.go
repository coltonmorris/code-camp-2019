package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Song struct {
	Artist string
	Album  string
	Name   string
}

type PlayCount struct {
	playlistName string
	songCoung    int
}

type Playlist struct {
	name  string
	songs []Song
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

func TransferPlaylist(newPlaylistName, playlistName string, ogAccount, acceptingAccount ServiceAccount) (SyncedSongs, error) {
	songs, err := ogAccount.GetPlaylist(playlistName)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"playlistName":     playlistName,
			"newName":          newPlaylistName,
			"OriginalAccount":  ogAccount.GetName(),
			"acceptingAccount": acceptingAccount.GetName(),
		})
		return SyncedSongs{}, err
	}

	if newPlaylistName == "" {
		newPlaylistName = playlistName
	}

	if newPlaylistName, err = acceptingAccount.CreatePlaylist(newPlaylistName, fmt.Sprintf("Copied playlist: %s from service: %s", playlistName, ogAccount.GetName())); err != nil {
		log.WithError(err).WithFields(log.Fields{
			"playlistName":     playlistName,
			"newName":          newPlaylistName,
			"OriginalAccount":  ogAccount.GetName(),
			"acceptingAccount": acceptingAccount.GetName(),
		})
		return SyncedSongs{}, err
	}

	return acceptingAccount.AddSongs(newPlaylistName, songs)
}
