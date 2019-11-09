package main

import (
	"fmt"
	"github.com/zmb3/spotify"
	"strings"
)

type sPlaylist struct {
	Id         string
	TrackCount uint
}

type SpotifyUser struct {
	ID        string
	token     string
	client    *spotify.Client
	Playlists map[string]sPlaylist
}

func NewSpotifyUser(t string, c *spotify.Client) *SpotifyUser {
	sUser := &SpotifyUser{
		token:     t,
		client:    c,
		Playlists: make(map[string]sPlaylist),
	}
	err := sUser.Initialize()
	if err != nil {
		fmt.Printf("failed to initialize spotify user ERR: %v", err)
	}
	return sUser
}

func (this *SpotifyUser) Initialize() error {
	playlistPage, _ := this.client.CurrentUsersPlaylists()
	for _, playlist := range playlistPage.Playlists {
		this.Playlists[playlist.Name] = sPlaylist{
			Id:         string(playlist.ID),
			TrackCount: playlist.Tracks.Total,
		}
	}

	user, err := this.client.CurrentUser()
	if err != nil {
		return err
	}

	this.ID = user.ID

	return nil
}

func (this *SpotifyUser) GetName() string {
	return "Spotify"
}

func (this *SpotifyUser) GetPlaylist(pName string) ([]Song, error) {
	storedP, ok := this.Playlists[pName]
	if !ok {
		return nil, fmt.Errorf("Playlists %s DNE", pName)
	}
	trackPage, err := this.client.GetPlaylistTracks(spotify.ID(storedP.Id))
	if err != nil {
		return nil, err
	}

	songs := make([]Song, 0)
	for _, track := range trackPage.Tracks {
		fTrack := track.Track
		names := make([]string, 0)
		for _, artist := range fTrack.SimpleTrack.Artists {
			names = append(names, artist.Name)
		}
		songs = append(songs, Song{
			Artist: strings.Join(names, ","),
			Album:  fTrack.Album.Name,
			Name:   fTrack.SimpleTrack.Name,
		})
	}

	for {
		if err = this.client.NextPage(trackPage); err != nil {
			fmt.Println("page error", err)
			break
		}
		fmt.Println("New PAGE")
		for _, track := range trackPage.Tracks {
			fTrack := track.Track
			names := make([]string, 0)
			for _, artist := range fTrack.SimpleTrack.Artists {
				names = append(names, artist.Name)
			}
			songs = append(songs, Song{
				Artist: strings.Join(names, ","),
				Album:  fTrack.Album.Name,
				Name:   fTrack.SimpleTrack.Name,
			})
		}
	}

	return songs, nil
}

func (this *SpotifyUser) CreatePlaylist(pName string, description string) (string, error) {
	pl, err := this.client.CreatePlaylistForUser(this.ID, pName, description, false)
	if err != nil {
		return "", err
	}
	this.Playlists[pName] = sPlaylist{
		Id: string(pl.ID),
	}
	return pName, nil
}

func (this *SpotifyUser) AddSongs(pName string, songs []Song) (SyncedSongs, error) {
	resp := SyncedSongs{
		AcceptedSongs: make([]Song, 0),
		FailedSongs:   make([]Song, 0),
	}
	ids := make([]spotify.ID, 0)
	for _, song := range songs {
		splitArtist := strings.Split(song.Artist, ",")
		a, b, c := splitArtist[0], song.Album, song.Name
		sQuery := fmt.Sprintf("artist:%s album:%s track:%s", a, b, c)
		//fmt.Println(sQuery)
		sResult, err := this.client.Search(sQuery, spotify.SearchTypeTrack)
		if err != nil {
			return SyncedSongs{}, err
		}
		added := false
		for i, track := range sResult.Tracks.Tracks {
			added = true
			if i == 0 {
				ids = append(ids, track.ID)
			} else {
			}
		}
		if added {
			resp.AcceptedSongs = append(resp.AcceptedSongs, song)
		} else {
			fmt.Printf("\n Failed to get song using query: %s", sQuery)
			fmt.Printf("\n Full artist: %s", song.Artist)
			resp.FailedSongs = append(resp.FailedSongs, song)
		}
	}

	i := 99
	cnt := len(ids)
	for {
		if cnt < 99 {
			_, err := this.client.AddTracksToPlaylist(spotify.ID(this.Playlists[pName].Id), ids...)
			if err != nil {
				fmt.Printf("EROR ON first PAGE: ERR: %v", err)
				return SyncedSongs{}, err
			}
			return resp, nil
		}

		if cnt > i {
			_, err := this.client.AddTracksToPlaylist(spotify.ID(this.Playlists[pName].Id), ids[i-99:i]...)
			if err != nil {
				fmt.Printf("EROR ON middle page PAGE: ERR: %v", err)
				return SyncedSongs{}, err
			}

		} else {
			_, err := this.client.AddTracksToPlaylist(spotify.ID(this.Playlists[pName].Id), ids[i-99:]...)
			if err != nil {
				fmt.Printf("EROR ON LAST PAGE: ERR: %v", err)
				return SyncedSongs{}, err
			}
			return resp, nil
		}
		i = i + 99
	}

}
