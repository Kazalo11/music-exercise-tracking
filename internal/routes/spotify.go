package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	clientManager "music-exercise-tracking/internal/client"
	songs "music-exercise-tracking/internal/songs"
)

const redirectURI = "http://localhost:8080/v1/spotify/callback"

var (
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadRecentlyPlayed),
	)
	state = "abc123"
)

func SpotifyRoutes(superRoute *gin.RouterGroup) {
	spotifyRouter := superRoute.Group("/spotify")
	{
		spotifyRouter.GET("/callback", completeSpotifyAuth)
		spotifyRouter.GET("/auth", getSpotifyAuthURL)
		spotifyRouter.POST("/songs", songs.GetRecentlyPlayed)

	}
}
func completeSpotifyAuth(c *gin.Context) {
	tok, err := auth.Token(c.Request.Context(), state, c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Couldn't get token"})
		log.Fatal(err)
		return
	}
	if st := c.Request.FormValue("state"); st != state {
		c.JSON(http.StatusNotFound, gin.H{"error": "State mismatch"})
		log.Fatalf("State mismatch: %s != %s\n", st, state)
		return
	}
	client := spotify.New(auth.Client(c.Request.Context(), tok))
	clientManager.SetClient(client)
	user, err := client.CurrentUser(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, fmt.Sprintf("Login Completed! You are logged in as: %s", user.ID))

	fmt.Println("You are logged in as:", user.ID)
}

func getSpotifyAuthURL(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url := auth.AuthURL(state)

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{"url": url})

}
