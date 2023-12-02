package services

import (
	"encoding/json"
	"log"
	"os"

	m "github.com/splitify/models"
)

// / Danceability, Energy, Instrumentalness, High Tempo, Low Tempo, High Valence, Low Valence
// for list of features
// iterate through and add to track map // calculate rolling average
// id -> attribute -> value
// / id -> map -> Track
//
//	SomeAttribute
//
// Generate list method
// Avg map -> Attribute -> Avg
// Iterate through the trackMap
// For each attribute that is not the track compare it to the average
// if above the average add to playlist with name

func GeneratePlaylists(req []m.Track, audioFeatures []m.AudioFeature) []m.Playlist {
	trackMap := make(map[string]map[string]interface{})
	avgMap := make(map[string]float64)
	playlistMap := make(map[string][]m.Track)

	for _, track := range req {
		trackMap[track.Track.ID] = make(map[string]interface{})
		trackMap[track.Track.ID]["Track"] = track
	}

	for i, feature := range audioFeatures {
		trackMap[feature.ID]["Danceability"] = feature.Danceability
		trackMap[feature.ID]["Energy"] = feature.Energy
		trackMap[feature.ID]["Instrumentalness"] = feature.Instrumentalness
		trackMap[feature.ID]["Tempo"] = feature.Tempo
		trackMap[feature.ID]["Valence"] = feature.Valence

		avgMap["Danceability"] = calcAvg(i+1, avgMap["Danceability"], feature.Danceability)
		avgMap["Energy"] = calcAvg(i+1, avgMap["Energy"], feature.Energy)
		avgMap["Instrumentalness"] = calcAvg(i+1, avgMap["Instrumentalness"], feature.Instrumentalness)
		avgMap["Tempo"] = calcAvg(i+1, avgMap["Tempo"], feature.Tempo)
		avgMap["Valence"] = calcAvg(i+1, avgMap["Valence"], feature.Valence)
	}

	for id, track := range trackMap {
		// we want to iterate through the values in track and compare against avg
		for key, val := range track {
			_, ok := avgMap[key]
			if !ok {
				continue
			}

			if val.(float64) >= avgMap[key] {
				// create list of tracks
				playlistMap[key] = append(playlistMap[key], trackMap[id]["Track"].(m.Track))
			}
		}
	}

	newPlaylists := []m.Playlist{}
	var id = 0
	for name, list := range playlistMap {
		newPlaylists = append(newPlaylists, createNewPlaylist(id, name, list))
		id++
	}

	return newPlaylists
}

func calcAvg(n int, prev float64, curr float64) float64 {
	return (float64(n-1)*prev + curr) / float64(n)
}

func getContent[T interface{}](fileName string) *T {
	// Let's first read the `config.json` file
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload T
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &payload
}

func map2[T interface{}, V interface{}](data []V, f func(int, V) T) []T {

	mapped := make([]T, len(data))

	for i, e := range data {
		mapped[i] = f(i, e)
	}

	return mapped
}

func createNewPlaylist(id int, name string, tracks []m.Track) m.Playlist {
	return m.Playlist{
		ID:          id,
		Images:      []m.Image{{URL: "https://source.unsplash.com/random/60×60/?music", Height: 60, Width: 60}},
		Name:        name,
		Description: "A playlist generated by AI",
		Owner: struct {
			DisplayName string "json:\"display_name\""
		}{"Current User"}, // "Current User"
		Tracks: struct {
			Total int       "json:\"total\""
			Items []m.Track "json:\"items\""
		}{len(tracks), map2(tracks, func(i int, t m.Track) m.Track {
			t.ID = i + 1
			return t
		})}, //len(tracks), tracks
	}
}