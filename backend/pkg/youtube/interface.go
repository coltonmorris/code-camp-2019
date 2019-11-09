package main

type Song struct {
	Artist string
	Album  string
	Name   string

	// Use this when you can't split a song between artist/album/name
	// split on delimeters?
	// remove everything in parens and brackets?
	Unknown string
}

// All the possible relevant info for a youtube video
type YoutubeVideoDetails struct {
	VideoId     string
	Description string
	Title       string
	Note        string
}

type PlayCount struct {
	playlistName string
	songCoung    int
}

type Playlist struct {
	name        string
	songs       []Song
	description string
	songCount   int
}

type PlaylistId string
type Playlists map[PlaylistId]Playlist

type SyncedSongs struct {
	AcceptedSongs []Song
	FailedSongs   []Song
}

type ServiceAccount interface {
	GetName() string
	GetPlaylist(string) ([]*Song, error)
	CreatePlaylist(string, string) (string, error)
	AddSongs(string, []Song) (SyncedSongs, error)
}
