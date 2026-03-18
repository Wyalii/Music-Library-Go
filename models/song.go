package models

type SearchResponse struct {
	ResultCount int    `json:"resultCount"`
	Results     []Song `json:"results"`
}
type Song struct {
	TrackName       string `json:"trackName"`
	ArtistName      string `json:"artistName"`
	CollectionName  string `json:"collectionName"`
	PreviewUrl      string `json:"previewUrl"`
	ArtworkUrl      string `json:"artworkUrl100"`
	Genre           string `json:"primaryGenreName"`
	ReleaseDate     string `json:"releaseDate"`
	TrackTimeMillis int    `json:"trackTimeMillis"`
}
