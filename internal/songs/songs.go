package somepackage

import (
	"context"
	"encoding/json"
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
	date, err := time.Parse(`"2006-01-02T15:04:05Z"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type SongsRequest struct {
	Start CustomTime `json:"start"`
	End   CustomTime `json:"end"`
}

func GetRecentlyPlayed(c *gin.Context) {
	var req SongsRequest
	var err error

	err = json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
		return
	}

	var songsStart []spotify.RecentlyPlayedItem
	var songsEnd []spotify.RecentlyPlayedItem
	playerResult := make([]spotify.RecentlyPlayedItem, 0)

	set := make(map[spotify.ID]bool)

	client := clientManager.GetClient()
	if client == nil {
		log.Fatal("Client is not initialized")
		return
	}

	songsStart, err = client.PlayerRecentlyPlayedOpt(context.Background(), &spotify.RecentlyPlayedOptions{
		Limit:        50,
		AfterEpochMs: req.Start.UnixMilli(),
	})
	if err != nil {
		log.Fatal(err)
	}

	songsEnd, err = client.PlayerRecentlyPlayedOpt(context.Background(), &spotify.RecentlyPlayedOptions{
		Limit:        50,
		AfterEpochMs: req.End.UnixMilli(),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Songs Start: %v \n", songsStart)
	log.Println("-------------")
	log.Printf("Songs End: %v", songsEnd)

	for _, item := range songsEnd {
		set[item.Track.ID] = true

	}

	for _, item := range songsStart {
		if !set[item.Track.ID] {
			playerResult = append(playerResult, item)
		}
	}
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, map[string][]mapping.Song{"songs": mapping.MapSpotifySongs(playerResult)})

}
