package main

type Song struct {
	artist string
	album  string
	name   string
}

type PlayCount struct {
	playlistName string
	songCoung    int
}

type User interface {
	GetEmail() string
	GetAvailableServiceAccounts() []string
	GetServiceAccount() ServiceAccount
	// TODO: Add service account
	GetAllPlaylistCounts() map[string]PlayCount
}

type ServiceAccount interface {
	GetName() string
	GetPlaylist(string) ([]Song, error)
	CreatePlaylist(string) (string, error)
	AddSongs(string, []Song) (SyncedSongs, error)
}

type API struct{}

func (api *API) TransferPlaylist(newPlaylistName, playlistName string, ogAccount, acceptingAccount ServiceAccount) (SyncedSongs, error) {
	songs, err := ogAccount.GetPlaylist()
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

	newPlaylistName, err := acceptingAccount.CreatePlaylist(newPlaylistName)
	if err != nil {
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
