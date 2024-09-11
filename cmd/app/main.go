// This example demonstrates how to authenticate with Spotify.
// In order to run this example yourself, you'll need to:
//
//  1. Register an application at: https://developer.spotify.com/my-applications/
//     - Use "http://localhost:8080/callback" as the redirect URI
//  2. Set the SPOTIFY_ID environment variable to the client ID you got in step 1.
//  3. Set the SPOTIFY_SECRET environment variable to the client secret from step 1.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"music-exercise-tracking/middleware"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadRecentlyPlayed),
	)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	var client *spotify.Client
	router := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
	)
	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	router.HandleFunc("GET /callback", completeAuth)
	router.HandleFunc("GET /auth", getAuthURL)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	go func() {
		client = <-ch

		// use the client to make calls that require authorization
		user, err := client.CurrentUser(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)
	}()

	fmt.Println("Server is running at http://localhost:8080")
	server.ListenAndServe()

}

func getAuthURL(w http.ResponseWriter, r *http.Request) {
	url := auth.AuthURL(state)

	response := map[string]string{"url": url}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	client := spotify.New(auth.Client(r.Context(), tok))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}
