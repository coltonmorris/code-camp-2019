package main

import (
	"os"
)

type YoutubeService struct {
	apiKey       string
	clientId     string
	clientSecret string
}

func (a *YoutubeService) GetPlaylist() Playlist {

}

func (a *YoutubeService) CreatePlaylist() {
}

func NewYoutubeService() (*YoutubeService, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("the env var YOUTUBE_API_KEY must exist")
	}

	clientId := os.Getenv("YOUTUBE_CLIENT_ID")
	if clientId == "" {
		return nil, fmt.Errorf("the env var YOUTUBE_CLIENT_ID must exist")
	}

	clientSecret := os.Getenv("YOUTUBE_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, fmt.Errorf("the env var YOUTUBE_CLIENT_SECRET must exist")
	}

	return &YoutubeService{
		apiKey:       apiKey,
		clientId:     clientId,
		clientSecret: clientSecret,
	}, nil
}
