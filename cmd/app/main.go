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
	"fmt"
	"log"
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

	http.HandleFunc("/callback", completeAuth)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	go func() {
		url := auth.AuthURL(state)
		fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

		// wait for auth to complete
		client = <-ch

		// use the client to make calls that require authorization
		user, err := client.CurrentUser(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)
	}()

	http.ListenAndServe(":8080", nil)

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
