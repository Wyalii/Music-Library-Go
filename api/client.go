package api

import (
	"GoMusicLibrary/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchSongs(query string) ([]models.Song, error) {
	url := fmt.Sprintf("https://itunes.apple.com/search?term=%s&media=music&limit=20", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
