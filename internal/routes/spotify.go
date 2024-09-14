package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	clientManager "music-exercise-tracking/internal/client"
	songs "music-exercise-tracking/internal/songs"
)

const redirectURI = "http://localhost:8080/v1/spotify/callback"

var (
	SPOTIFY_ID = os.Getenv("SPOTIFY_ID")
	auth       = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadRecentlyPlayed),
		spotifyauth.WithClientID(SPOTIFY_ID),
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

	c.JSON(http.StatusOK, gin.H{"message": "Login completed"})

	fmt.Println("You are logged in as:", user.ID)
}

func getSpotifyAuthURL(c *gin.Context) {
	url := auth.AuthURL(state)

	c.JSON(http.StatusOK, gin.H{"url": url})

}
