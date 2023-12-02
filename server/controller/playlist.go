package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	m "github.com/splitify/models"
	"github.com/splitify/services"
	"github.com/splitify/utils"
)

type GenerateRequest struct {
	Tracks []m.Track `json:"tracks"`
}

type AudioFeatureResponse struct {
	AudioFeatures []m.AudioFeature `json:"audio_features"`
}

func fetchAudioFeatures(token string, url string) ([]m.AudioFeature, error) {
	//  from track ids fetch audio features
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("Audio Feature Get request failed with: %s", resp.Status)
	}

	defer resp.Body.Close()
	var features struct {
		AudioFeatures []m.AudioFeature `json:"audio_features"`
	}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &features)
	if err != nil {
		return nil, err
	}

	return features.AudioFeatures, nil
}

func GeneratePlaylists(c *gin.Context) {
	// get tracks from request
	var request GenerateRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	accessToken := c.Request.Header["Authorization"][0]
	if len(accessToken) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	idArr := utils.Map(request.Tracks, func(i int, track m.Track) string {
		return track.Track.ID
	})

	url := fmt.Sprintf("https://api.spotify.com/v1/audio-features?ids=%s", strings.Join(idArr, ","))
	audio_features, err := fetchAudioFeatures(accessToken, url)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	playlists := services.GeneratePlaylists(request.Tracks, audio_features)
	c.IndentedJSON(http.StatusOK, gin.H{"playlists": playlists})
}

func savePlaylists(c *gin.Context) {
	println("Saving playlist")
	client := &http.Client{}
	var request struct {
		UserID      string   `json:"user_id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		IsPublic    bool     `json:"is_public"`
		Uris        []string `json:"uris"`
	}

	if err := c.BindJSON(&request); err != nil {
		return
	}

	accessToken := c.Request.Header["Authorization"][0]
	if len(accessToken) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jsonBytes, err := json.Marshal(map[string]interface{}{
		"name":        request.Name,
		"description": request.Description,
		"public":      request.IsPublic,
	})
	if err != nil {
		fmt.Println("Error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	path := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", request.UserID)
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error creating save request")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	req.Header.Add("Authorization", accessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error saving playlist")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(c.Request.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		println(string(body))
		fmt.Println("Unable to create playlist")
		c.AbortWithStatus(resp.StatusCode)
		return
	}

	var savedPlaylistResponse interface{}
	err = json.NewDecoder(resp.Body).Decode(&savedPlaylistResponse)
	if err != nil {
		fmt.Println("Unable to decode saved playlist")
		c.AbortWithStatus(resp.StatusCode)
		return
	}

	fmt.Printf("%s", savedPlaylistResponse)
}
