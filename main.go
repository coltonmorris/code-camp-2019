package main

import (
	"os"
)

func init() {
	os.Setenv("SPOTIFY_ID", "462eace056d94eaaa00f678b93d9bd0d")
	os.Setenv("SPOTIFY_SECRET", "4b6e16b7fa2349938b9f58f1c685b269")
}

func main() {
	api := &API{
		Users: make(map[string]*LameUser, 0),
	}

	RunHttpServer(api)
}
