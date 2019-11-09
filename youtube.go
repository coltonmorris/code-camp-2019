package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct {
	playlists Playlists
	authUrl   string
	service   *youtube.Service
	config    *oauth2.Config // every user will needs this config to generate their oauth2 token
	ctx       context.Context
	apiKey    string // typically used for non oauth2 requests
}

func (ys *YoutubeService) GetPlaylists() []PlayCount {
	resp := make([]PlayCount, 0)
	for _, value := range ys.playlists {
		resp = append(resp, PlayCount{
			PlaylistName: value.name,
			SongCount:    uint(value.songCount),
		})
	}
	return resp
}

func NewYoutubeService(ctx context.Context, user string) (*YoutubeService, error) {
	// TODO move env vars to main.go
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("the env var YOUTUBE_API_KEY must exist")
	}

	secretJson := os.Getenv("YOUTUBE_CLIENT_SECRET_JSON")
	if secretJson == "" {
		return nil, fmt.Errorf("the env var YOUTUBE_CLIENT_SECRET_JSON must exist")
	}

	var jsonMap map[string]interface{}
	secretBytes := []byte(secretJson)

	if err := json.Unmarshal(secretBytes, &jsonMap); err != nil {
		return nil, fmt.Errorf("the env var YOUTUBE_CLIENT_SECRET_JSON was invalid json: %v", err)
	}

	scopes := []string{
		youtube.YoutubeScope,
		youtube.YoutubeUploadScope,
		youtube.YoutubeForceSslScope,
	}

	config, err := google.ConfigFromJSON(secretBytes, scopes...)
	if err != nil {
		return nil, err
	}

	config.RedirectURL = "http://localhost:8080/callback/youtube"

	return &YoutubeService{
		authUrl: config.AuthCodeURL(user),
		service: nil,
		config:  config,
		ctx:     ctx,
		apiKey:  apiKey,
	}, nil
}

// LoadAllPlaylists queries Youtube API for a list of all the authenticated users playlists
func (ys *YoutubeService) LoadAllPlaylists() error {
	// must be authenticated first
	if ys.service == nil {
		return fmt.Errorf("was not authenticated before LoadAllPlaylists")
	}

	// TODO eventually need to loop and use pagination... dont need to right now because most people don't have more than 50 playlists
	playlists, err := ys.loadPlaylists("snippet,contentDetails", int64(50), "", "")
	if err != nil {
		return err
	}

	ys.playlists = playlists

	return nil
}

func (ys *YoutubeService) GetPlaylist(playlistId string) ([]Song, error) {
	// must be authenticated first
	if ys.service == nil {
		log.Fatal("was not authenticated before getPlaylist")
	}

	// TODO use this for pagination
	pageToken := ""

	ytVideos, pt, err := ys.loadPlaylistItems("snippet,contentDetails", int64(50), pageToken, playlistId, "")
	fmt.Println(ytVideos)
	if err != nil {
		return nil, err
	}
	for pt != "" {
		vids := make([]Song, 0)
		vids, pt, err = ys.loadPlaylistItems("snippet,contentDetails", int64(50), pt, playlistId, "")
		if err != nil {
			return nil, err
		}
		ytVideos = append(ytVideos, vids...)
	}
	fmt.Println(ytVideos)

	return ytVideos, nil
}

func (ys *YoutubeService) CreatePlaylist(pName, desc string) (string, error) {
	// must be authenticated first
	if ys.service == nil {
		log.Fatal("was not authenticated before createPlaylist")
	}

	_, err := ys.service.Playlists.Insert("snippet,status", &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title: pName,
		},
	}).Do()
	return pName, err
}

// Authenticate takes in the code from the redirect url and turns it into a longer living token. The code will be in the http.Request.FormValue("code"). Be sure to validate that FormValue("state") exists for security reasons...
func (ys *YoutubeService) Authenticate(code string) error {
	token, err := ys.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}

	client := ys.config.Client(oauth2.NoContext, token)
	ys.service, err = youtube.New(client)
	if err != nil {
		return err
	}

	if err := ys.LoadAllPlaylists(); err != nil {
		fmt.Print("ERROR LOADING")
		return err
	}

	return nil
}

// loadPlaylists is a single request to Youtube APIs playlist endpoint. Some times a user will have more than the maxResults allowed, so this will have to be called multiple times using the pagination parameter "pageToken"
func (ys *YoutubeService) loadPlaylistItems(part string, maxResults int64, pageToken string, playlistId string, videoId string) ([]Song, string, error) {
	call := ys.service.PlaylistItems.List(part)
	call = call.MaxResults(maxResults)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	if playlistId != "" {
		call = call.PlaylistId(playlistId)
	}
	if videoId != "" {
		call = call.VideoId(videoId)
	}

	response, err := call.Do()
	fmt.Printf("RES: %v", response)
	if err != nil {
		return nil, "", err
	}

	ret := []Song{}

	for _, ele := range response.Items {
		fmt.Printf("ELE: %v", ele)
		ret = append(ret, Song{
			Name: ele.Snippet.Title,
		})
	}

	return ret, response.NextPageToken, nil
}

// loadPlaylists is a single request to Youtube APIs playlist endpoint. Some times a user will have more than the maxResults allowed, so this will have to be called multiple times using the pagination parameter "pageToken"
func (ys *YoutubeService) loadPlaylists(part string, maxResults int64, pageToken string, playlistId string) (Playlists, error) {
	call := ys.service.Playlists.List(part)
	call = call.Mine(true)
	call = call.MaxResults(maxResults)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	if playlistId != "" {
		call = call.Id(playlistId)
	}
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var ps Playlists = make(map[PlaylistId]Playlist)

	for _, ele := range response.Items {
		ps[PlaylistId(ele.Id)] = Playlist{
			name:      ele.Snippet.Title,
			songCount: int(ele.ContentDetails.ItemCount),
		}
	}

	return ps, nil
}

func (ys *YoutubeService) GetPlaylistIdFromName(pName string) string {
	for key, value := range ys.playlists {
		if value.name == pName {
			return string(key)
		}
	}
	return ""
}

func (ys *YoutubeService) AddSongs(pName string, songs []Song) (SyncedSongs, error) {
	ids := make([]string, 0)
	finalres := SyncedSongs{
		AcceptedSongs: make([]Song, 0),
		FailedSongs:   make([]Song, 0),
	}
	for _, song := range songs {
		res, err := ys.service.Search.List("snippet").Q(fmt.Sprintf("%s - %s", song.Artist, song.Name)).Type("video").MaxResults(1).Do()
		if err != nil {
			fmt.Print("Call failed", err)
			return SyncedSongs{}, err
		}

		added := false
		for i, item := range res.Items {
			added = true
			if i == 0 {
				ids = append(ids, item.Id.VideoId)
			}
		}
		if added {
			finalres.AcceptedSongs = append(finalres.AcceptedSongs, song)
		} else {
			finalres.FailedSongs = append(finalres.FailedSongs, song)
		}

	}

	pId := ys.GetPlaylistIdFromName(pName)
	for _, id := range ids {
		_, err := ys.service.PlaylistItems.Insert("snippet", &youtube.PlaylistItem{
			Id: id,
			Snippet: &youtube.PlaylistItemSnippet{
				PlaylistId: pId,
			},
		}).Do()
		if err != nil {
			fmt.Println("ERR inserting song: ", err)
		}
	}

	return finalres, nil
}

func (ys *YoutubeService) GetName() string {
	return "Youtube"
}
