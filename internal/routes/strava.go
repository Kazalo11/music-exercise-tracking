package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ()

type TokenReponse struct {
	TokenType string `json:"token_type"`
	RefreshTokenResponse
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

type RefreshTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
}

func StravaRoutes(superRoute *gin.RouterGroup) {
	stravaRouter := superRoute.Group("strava")
	{
		stravaRouter.GET("/auth", getStravaAuthURL)
		stravaRouter.GET("/exchange_token", getStravaToken)
		stravaRouter.POST("/refresh", refreshStravaAuthToken)
	}
}

func refreshStravaAuthToken(c *gin.Context) {
	var refresh_token RefreshTokenResponse
	err := json.NewDecoder(c.Request.Body).Decode(&refresh_token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	var tokens TokenReponse

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})

}

func getStravaAuthURL(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CLIENT_ID := os.Getenv("CLIENT_ID")

	authURL := fmt.Sprintf("http://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost:8080/v1/strava/exchange_token&approval_prompt=force&scope=read_all", CLIENT_ID)

	c.JSON(http.StatusOK, gin.H{"url": authURL})
}

func getStravaToken(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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

	var tokens TokenReponse

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode json"})
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}
