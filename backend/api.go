package main

type API struct {
	Storage
}

type SyncedSongs struct {
	ServiceAccount string
	FailedSongs    []Song
	CompletedSongs []Song
}

func (api *API) SyncPlaylist(hostService MusicService, syncService MusicService, playlist string) (SyncedSongs, error) {
	originalPlaylist, err := hostService.GetPlaylist(playlist)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"hostService": hostService.GetName(),
			"syncService": syncService.GetName(),
			"playlist":    playlist,
		}).Warn("Failed to get playlist from host account")
		return SyncedSongs{}, err
	}

	newPlaylist, err := syncService.CreatePlaylist(playlist)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"hostService": hostService.GetName(),
			"syncService": syncService.GetName(),
			"playlist":    playlist,
		}).Warn("Failed to make playlist on sync service")
		return SyncedSongs{}, err
	}

	syncedSongs := &SyncedSongs{
		ServiceAccount: syncService.GetName(),
		FailedSongs:    []*Song{},
		CompletedSongs: []*Song{},
	}

	for _, song := range originalPlaylist.Songs {
		ok := syncService.FindAddSong(newPlaylist, song)
		if ok {
			syncedSongs.CompletedSongs = append(syncedSongs.CompletedSongs, song)
		} else {
			log.WithFields(log.Fields{
				"song":     hostService.GetName(),
				"playlist": newPlaylist,
			}).Warn("Failed to add song to playlist")
			syncedSongs.FailedSongs = append(syncedSongs.FailedSongs, song)
		}
	}

	return syncedSongs, nil
}
