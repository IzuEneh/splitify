package models

type Artist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Track struct {
	ID      int    `json:"id"`
	AddedAt string `json:"added_at"`
	AddedBy Artist `json:"added_by"`
	IsLocal bool   `json:"is_local"`
	Track   struct {
		Album struct {
			AlbumType        string   `json:"album_type"`
			TotalTracks      int      `json:"total_tracks"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href                 string  `json:"href"`
			ID                   string  `json:"id"`
			Images               []Image `json:"images"`
			Name                 string  `json:"name"`
			ReleaseDate          string  `json:"release_date"`
			ReleaseDatePrecision string  `json:"release_date_precision"`
			Restrictions         struct {
				Reason string `json:"reason"`
			} `json:"restrictions"`
			Type    string   `json:"type"`
			URI     string   `json:"uri"`
			Artists []Artist `json:"artists"`
		} `json:"album"`
		Artists          []Artist `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurarionMS       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIDs      struct {
			Isrc string `json:"isrc"`
			Ean  string `json:"ean"`
			Upc  string `json:"upc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href         string   `json:"href"`
		ID           string   `json:"id"`
		IsPlayable   bool     `json:"is_playable"`
		LinkedFrom   struct{} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
		IsLocal     bool   `json:"is_local"`
	} `json:"track"`
}

type Playlist struct {
	ID          int     `json:"id"`
	Images      []Image `json:"images"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Owner       struct {
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Tracks struct {
		Total int     `json:"total"`
		Items []Track `json:"items"`
	} `json:"tracks"`
}

type AudioFeature struct {
	Acousticness     float64 `json:"acousticness"`
	AnalysisURL      string  `json:"analysis_url"`
	Danceability     float64 `json:"danceability"`
	DurationMs       int     `json:"duration_ms"`
	Energy           float64 `json:"energy"`
	ID               string  `json:"id"`
	Instrumentalness float64 `json:"instrumentalness"`
	Key              int     `json:"key"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float64 `json:"speechiness"`
	Tempo            float64 `json:"tempo"`
	TimeSignature    int     `json:"time_signature"`
	TrackHref        string  `json:"track_href"`
	Type             string  `json:"type"`
	URI              string  `json:"uri"`
	Valence          float64 `json:"valence"`
}
