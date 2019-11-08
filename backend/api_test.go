package backend

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	wordGen "github.com/zhexuany/wordGenerator"
	"testing"
)

func TestTransferPlaylist(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&serviceTestImpl{})
	a := newServiceTestImpl("Spotify")
	b := newServiceTestImpl("Youtube")

	pName := "C/B Torture Tunes"
	a.GenerateRandomPlaylist(pName, 17)

	usr := UserImpl{}
	fmt.Println(usr.TransferPlaylist("hip rap", pName, a, b))
	db.Create(a)
	db.Create(b)
}

type serviceTestImpl struct {
	gorm.Model
	Name      string
	Playlists map[string][]Song `gorm:"-"`
}

func newServiceTestImpl(name string) *serviceTestImpl {
	return &serviceTestImpl{
		Name:      name,
		Playlists: map[string][]Song{},
	}
}

func (this *serviceTestImpl) GetName() string {
	return this.Name
}

func (this *serviceTestImpl) GetPlaylist(pName string) ([]Song, error) {
	songs, ok := this.Playlists[pName]
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
		if _, ok := this.Playlists[createName]; !ok {
			this.Playlists[createName] = []Song{}
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

	_, ok := this.Playlists[pName]
	if !ok {
		return SyncedSongs{}, fmt.Errorf("Playlist %s DNE", pName)
	}

	for _, song := range songs {
		this.Playlists[pName] = append(this.Playlists[pName], song)
	}

	resp.AcceptedSongs = songs

	return resp, nil
}

func (this *serviceTestImpl) GenerateRandomPlaylist(pName string, songCount int) {
	this.Playlists[pName] = []Song{}
	for i := 0; i < songCount; i++ {
		strs := wordGen.GetWords(3, 10)
		this.Playlists[pName] = append(this.Playlists[pName], Song{
			artist: strs[0],
			album:  strs[1],
			name:   strs[2],
		})
	}
}
