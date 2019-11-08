package backend

import (
	log "github.com/sirupsen/logrus"
)

type Song struct {
	artist string
	album  string
	name   string
}

type PlayCount struct {
	playlistName string
	songCoung    int
}

type Playlist struct {
	name  string
	songs []Song
}

type User interface {
	GetEmail() string
	GetAvailableServiceAccounts() []string
	GetServiceAccount() ServiceAccount
	// TODO: Add service account
	GetAllPlaylistCounts() map[string]PlayCount
}

type SyncedSongs struct {
	AcceptedSongs []Song
	FailedSongs   []Song
}

type ServiceAccount interface {
	GetName() string
	GetPlaylist(string) ([]Song, error)
	CreatePlaylist(string) (string, error)
	AddSongs(string, []Song) (SyncedSongs, error)
}

type UserImpl struct{}

func (user *UserImpl) TransferPlaylist(newPlaylistName, playlistName string, ogAccount, acceptingAccount ServiceAccount) (SyncedSongs, error) {
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

	if newPlaylistName, err = acceptingAccount.CreatePlaylist(newPlaylistName); err != nil {
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
