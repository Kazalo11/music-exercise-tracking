package routes

import (
	"encoding/json"
	"fmt"
	"music-exercise-tracking/internal/types"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LastFMRoutes(superRoute *gin.RouterGroup) {
	lastFMRouter := superRoute.Group("lastfm")
	{
		lastFMRouter.POST("/tracks", getTracks)
	}
}

func getTracks(c *gin.Context) {
	var recentTrackRequest types.RecentTracksRequest
	err := json.NewDecoder(c.Request.Body).Decode(&recentTrackRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to decode request due to: \n %+v", err)})
	}
	API_KEY := os.Getenv("LASTFM_API_KEY")

	baseUrl, err := url.Parse("https://ws.audioscrobbler.com/2.0")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the url"})
		return
	}
	params := url.Values{}
	params.Set("method", "user.getrecenttracks")
	params.Set("user", recentTrackRequest.UserName)
	params.Set("api_key", API_KEY)
	params.Set("from", strconv.FormatInt(recentTrackRequest.Start.Unix(), 10))
	params.Set("to", strconv.FormatInt(recentTrackRequest.End.Unix(), 10))
	params.Set("format", "json")

	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get track info"})
		return
	}
	defer resp.Body.Close()

	var response types.RecentTracksResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Printf("%+v", response)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to decode response due to: \n %+v", err)})
	}

	c.JSON(http.StatusOK, gin.H{"tracks": response.RecentTracks.Tracks})

}
