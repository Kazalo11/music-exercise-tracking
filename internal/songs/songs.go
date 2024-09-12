package somepackage

import (
	"context"
	"log"
	clientManager "music-exercise-tracking/internal/client"
	"music-exercise-tracking/internal/mapping"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func GetRecentlyPlayed(c *gin.Context) {
	var playerResult []spotify.RecentlyPlayedItem
	var err error
	client := clientManager.GetClient()
	if client == nil {
		log.Fatal("Client is not initialized")
		return
	}

	playerResult, err = client.PlayerRecentlyPlayed(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string][]mapping.Song{"songs": mapping.MapSpotifySongs(playerResult)})

}
