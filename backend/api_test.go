package backend

import (
	"fmt"

	wordGen "github.com/zhexuany/wordGenerator"
	"testing"
)

func TestTransferPlaylist(t *testing.T) {
	a := newServiceTestImpl("Spotify")
	b := newServiceTestImpl("Youtube")

	pName := "C/B Torture Tunes"
	a.GenerateRandomPlaylist(pName, 17)

	usr := UserImpl{}
	fmt.Println(usr.TransferPlaylist("hip rap", pName, a, b))
}

type serviceTestImpl struct {
	name      string
	playlists map[string][]Song
}

func newServiceTestImpl(name string) *serviceTestImpl {
	return &serviceTestImpl{
		name:      name,
		playlists: map[string][]Song{},
	}
}

func (this *serviceTestImpl) GetName() string {
	return this.name
}

func (this *serviceTestImpl) GetPlaylist(pName string) ([]Song, error) {
	songs, ok := this.playlists[pName]
	if ok {
		return songs, nil
	}
	return nil, fmt.Errorf("Playlist %s DNE", pName)
}

func (this *serviceTestImpl) CreatePlaylist(pName string) (string, error) {
	createName := pName
	i := 0
	for {
		if i != 0 {
			createName = fmt.Sprintf("pName-%d", i)
		}
		if _, ok := this.playlists[createName]; !ok {
			this.playlists[createName] = []Song{}
			return createName, nil
		}
		i++
	}
}

func (this *serviceTestImpl) AddSongs(pName string, songs []Song) (SyncedSongs, error) {
	resp := SyncedSongs{
		AcceptedSongs: []Song{},
		FailedSongs:   []Song{},
	}

	_, ok := this.playlists[pName]
	if !ok {
		return SyncedSongs{}, fmt.Errorf("Playlist %s DNE", pName)
	}

	for _, song := range songs {
		this.playlists[pName] = append(this.playlists[pName], song)
	}

	resp.AcceptedSongs = songs

	return resp, nil
}

func (this *serviceTestImpl) GenerateRandomPlaylist(pName string, songCount int) {
	this.playlists[pName] = []Song{}
	for i := 0; i < songCount; i++ {
		strs := wordGen.GetWords(3, 10)
		this.playlists[pName] = append(this.playlists[pName], Song{
			artist: strs[0],
			album:  strs[1],
			name:   strs[2],
		})
	}
}
