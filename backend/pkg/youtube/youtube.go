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

type Playlist struct{}

type YoutubeService struct {
	authUrl string
	service *youtube.Service
	config  *oauth2.Config // every user will needs this config to generate their oauth2 token
	ctx     context.Context
	apiKey  string // typically used for non oauth2 requests
}

func (ys *YoutubeService) GetPlaylist() *Playlist {
	// must be authenticated first
	if ys.service == nil {
		log.Fatal("was not authenticated before getPlaylist")
	}

	return nil
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

func playlistsList(service *youtube.Service, part string, maxResults int64, pageToken string, playlistId string) (*youtube.PlaylistListResponse, error) {
	// part    = flag.String("part", "snippet", "Comma-separated list of playlist resource parts that API response will include.")
	// maxResults    = flag.Int64("maxResults", 5, "The maximum number of playlist resources to include in the API response.")
	// pageToken    = flag.String("pageToken", "", "Token that identifies a specific page in the result set that should be returned.")
	// playlistId       = flag.String("playlistId", "", "Retrieve information about this playlist.")

	call := service.Playlists.List(part)
	call = call.Mine(true)
	call = call.MaxResults(maxResults)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	if playlistId != "" {
		call = call.Id(playlistId)
	}
	response, err := call.Do()

	return response, err
}

func main() {
	os.Setenv("YOUTUBE_API_KEY", "")
	os.Setenv("YOUTUBE_CLIENT_SECRET_JSON", "")
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

		res, err := playlistsList(ys.service, "snippet,contentDetails", int64(5), "", "")
		if err != nil {
			log.Fatal(err)
		}

		for _, playlist := range res.Items {
			playlistId := playlist.Id
			playlistTitle := playlist.Snippet.Title

			fmt.Println(playlistId, ": ", playlistTitle)
		}
	}
}
