package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func TransferPlaylist(newPlaylistName, playlistName string, ogAccount, acceptingAccount ServiceAccount) (SyncedSongs, error) {
	songs, err := ogAccount.GetPlaylist(playlistName)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"playlistName":     playlistName,
			"newName":          newPlaylistName,
			"OriginalAccount":  ogAccount.GetName(),
			"acceptingAccount": acceptingAccount.GetName(),
		}).Warn("Failed get og playlist")
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
		}).Warn("Failed create playlist")
		return SyncedSongs{}, err
	}

	return acceptingAccount.AddSongs(newPlaylistName, songs)
}
