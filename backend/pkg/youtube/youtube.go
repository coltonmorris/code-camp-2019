package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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

func (ys *YoutubeService) GetPlaylist(playlistId string) ([]*Song, error) {
	// must be authenticated first
	if ys.service == nil {
		log.Fatal("was not authenticated before getPlaylist")
	}

	// TODO use this for pagination
	pageToken := ""

	ytVideos, err := ys.loadPlaylistItems("snippet,contentDetails", int64(50), pageToken, playlistId, "")
	if err != nil {
		return nil, err
	}

	// TODO use some fancy string parsing to turn a ytVideo into a song
	fmt.Println("ytVideos: ", ytVideos)

	return nil, nil
}

func (ys *YoutubeService) CreatePlaylist() {
	// must be authenticated first
	if ys.service == nil {
		log.Fatal("was not authenticated before createPlaylist")
	}
}

// Authenticate takes in the code from the redirect url and turns it into a longer living token. The code will be in the http.Request.FormValue("code"). Be sure to validate that FormValue("state") exists for security reasons...
func (ys *YoutubeService) Authenticate(code string) {
	token, err := ys.config.Exchange(ys.ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := ys.config.Client(ys.ctx, token)
	ys.service, err = youtube.New(client)
	if err != nil {
		log.Fatal(err)
	}
}

func NewYoutubeService(ctx context.Context) (*YoutubeService, error) {
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

	config.RedirectURL = "http://localhost:8080/youtube_callback"

	return &YoutubeService{
		authUrl: config.AuthCodeURL("state"),
		service: nil,
		config:  config,
		ctx:     ctx,
		apiKey:  apiKey,
	}, nil
}

// loadPlaylists is a single request to Youtube APIs playlist endpoint. Some times a user will have more than the maxResults allowed, so this will have to be called multiple times using the pagination parameter "pageToken"
func (ys *YoutubeService) loadPlaylistItems(part string, maxResults int64, pageToken string, playlistId string, videoId string) ([]*YoutubeVideoDetails, error) {
	call := ys.service.PlaylistItems.List(part)
	call = call.MaxResults(maxResults)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	if playlistId != "" {
		call = call.Id(playlistId)
	}
	if videoId != "" {
		call = call.VideoId(videoId)
	}

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	ret := []*YoutubeVideoDetails{}

	for _, ele := range response.Items {
		ret = append(ret, &YoutubeVideoDetails{
			Description: ele.Snippet.Description,
			Title:       ele.Snippet.Title,
			Note:        ele.ContentDetails.Note,
			VideoId:     ele.ContentDetails.VideoId,
		})
	}

	return ret, nil
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
			name:        ele.Snippet.Title,
			description: ele.Snippet.Description,
			// songs:
			songCount: int(ele.ContentDetails.ItemCount),
		}
	}

	return ps, nil
}

func main() {
	os.Setenv("YOUTUBE_API_KEY", "AIzaSyB6zO_vHL4r3bjzjoBGEDpirBOa_ozoRkM")
	os.Setenv("YOUTUBE_CLIENT_SECRET_JSON", "{\"web\":{\"client_id\":\"151708764487-rp0lbvppvfudv9p24miqv6lm6jf2o3kt.apps.googleusercontent.com\",\"project_id\":\"coce-camp-2019\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_secret\":\"qR-UNuoMgVcdxFkR4njlB9PZ\",\"redirect_uris\":[\"http://synclist.tech/youtube_callback\",\"http://localhost:8080/youtube_callback\",\"http://www.synclist.tech/youtube_callback\"],\"javascript_origins\":[\"http://synclist.tech\",\"http://www.synclist.tech\",\"http://localhost:8080\"]}}")

	ys, err := NewYoutubeService(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("auth url: ", ys.authUrl)
	// input auth code
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("auth url -> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		fmt.Println("---------------")

		ys.Authenticate(text)

		if err := ys.LoadAllPlaylists(); err != nil {
			log.Fatal(err)
		}

		for id, playlist := range ys.playlists {
			fmt.Println(id, ": ", playlist)
		}
	}
}
