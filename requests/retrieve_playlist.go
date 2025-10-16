package requests

import (
	"encoding/json"
	"log"
	"todo-api/helpers"
)

type SpotifyPlaylistResponse struct {
	Playlists Playlists `json:"playlists"`
}

type Playlists struct {
	Items []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	Images       []Image      `json:"images"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	Url string `json:"url"`
}

var RetrievePlaylist = func(genre string) SpotifyPlaylistResponse {
	url := "https://api.spotify.com/v1/search?q=remaster%2520genre%2520%22" + genre + "%22&type=playlist&market=BR&limit=1"

	token, err := helpers.GetAccessToken()
	if err != nil {
		log.Fatal("Error getting access token: %v", err)
	}

	resp, err := helpers.PerformHttpRequest(url, token, "GET", nil)
	if err != nil {
		log.Fatal("Error retrieving playlist from Spotify: %v", err)
	}

	var spotifyPlaylistResponse SpotifyPlaylistResponse
	err = json.Unmarshal(resp, &spotifyPlaylistResponse)
	if err != nil {
		log.Fatal("Error unmarshaling playlist to Playlists struct: %v", err)
	}
	return spotifyPlaylistResponse
}
