package somepackage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	clientManager "music-exercise-tracking/internal/client"
	"music-exercise-tracking/internal/mapping"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05.000"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type SongsRequest struct {
	After  CustomTime `json:"after"`
	Before CustomTime `json:"before"`
}

func GetRecentlyPlayed(c *gin.Context) {
	var req SongsRequest
	var err error

	err = json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
		return
	}

	var playerResult1 []spotify.RecentlyPlayedItem
	var playerResult2 []spotify.RecentlyPlayedItem
	playerResult := make([]spotify.RecentlyPlayedItem, 0)

	set := make(map[spotify.ID]bool)

	client := clientManager.GetClient()
	if client == nil {
		log.Fatal("Client is not initialized")
		return
	}

	rpo := &spotify.RecentlyPlayedOptions{
		Limit:         50,
		BeforeEpochMs: req.Before.UnixMilli(),
	}

	playerResult1, err = client.PlayerRecentlyPlayedOpt(context.Background(), rpo)
	if err != nil {
		log.Fatal(err)
	}

	rpo2 := &spotify.RecentlyPlayedOptions{
		Limit:         50,
		BeforeEpochMs: req.After.UnixMilli(),
	}

	playerResult2, err = client.PlayerRecentlyPlayedOpt(context.Background(), rpo2)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range playerResult2 {
		set[item.Track.ID] = true
	}

	for _, item := range playerResult1 {
		fmt.Println(set[item.Track.ID])
		if !set[item.Track.ID] {
			playerResult = append(playerResult, item)
		}
	}

	c.JSON(http.StatusOK, map[string][]mapping.Song{"songs": mapping.MapSpotifySongs(playerResult)})

}
