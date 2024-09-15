package routes

import (
	"encoding/json"
	"fmt"
	authManager "music-exercise-tracking/internal/client"
	"music-exercise-tracking/internal/types"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func StravaRoutes(superRoute *gin.RouterGroup) {
	stravaRouter := superRoute.Group("strava")
	{
		stravaRouter.GET("/auth", getStravaAuthURL)
		stravaRouter.GET("/exchange_token", getStravaToken)
		stravaRouter.POST("/refresh", refreshStravaAuthToken)
		stravaRouter.GET("/athlete", getAthlete)
		stravaRouter.GET("/activities", getActivities)
	}
}

func getActivities(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities?per_page=20", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	access_token := authManager.GetAccessToken()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get activities"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Received non-200 response: %d", resp.StatusCode)})
		return
	}

	var activities []types.Activity

	err = json.NewDecoder(resp.Body).Decode(&activities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	for i := range activities {
		activities[i].Finish = activities[i].CalculateFinishTime()
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})

}

func getAthlete(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	access_token := authManager.GetAccessToken()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get athlete"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Received non-200 response: %d", resp.StatusCode)})
		return
	}

	var athlete types.Athlete

	err = json.NewDecoder(resp.Body).Decode(&athlete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	c.JSON(http.StatusOK, gin.H{"athlete": athlete})

}

func refreshStravaAuthToken(c *gin.Context) {
	var refresh_token types.RefreshTokenResponse
	err := json.NewDecoder(c.Request.Body).Decode(&refresh_token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	CLIENT_ID := os.Getenv("CLIENT_ID")

	formData := url.Values{}
	formData.Set("client_id", CLIENT_ID)
	formData.Set("client_secret", CLIENT_SECRET)
	formData.Set("refresh_token", refresh_token.RefreshToken)
	formData.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", "https://www.strava.com/api/v3/oauth/token", strings.NewReader(formData.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}
	defer resp.Body.Close()

	var tokens types.TokenReponse

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	authManager.SetAccessToken(tokens.AccessToken)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})

}

func getStravaAuthURL(c *gin.Context) {

	CLIENT_ID := os.Getenv("CLIENT_ID")

	authURL := fmt.Sprintf("http://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost:8080/v1/strava/exchange_token&approval_prompt=force&scope=activity:read_all", CLIENT_ID)
	c.JSON(http.StatusOK, gin.H{"url": authURL})
}

func getStravaToken(c *gin.Context) {

	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	CLIENT_ID := os.Getenv("CLIENT_ID")
	code := c.Query("code")

	formData := url.Values{}
	formData.Set("client_id", CLIENT_ID)
	formData.Set("client_secret", CLIENT_SECRET)
	formData.Set("code", code)
	formData.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", "https://www.strava.com/api/v3/oauth/token", strings.NewReader(formData.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}
	defer resp.Body.Close()

	var tokens types.TokenReponse
	err = json.NewDecoder(resp.Body).Decode(&tokens)
	authManager.SetAccessToken(tokens.AccessToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token obtained successfully"})
}
