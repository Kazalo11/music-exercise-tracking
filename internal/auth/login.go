package auth

import (
	"fmt"
	"log"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadRecentlyPlayed),
	)
	ch     = make(chan *spotify.Client)
	state  = "abc123"
	Client *spotify.Client
)

func AuthURL() string {
	return auth.AuthURL(state)
}

func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	Client = spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- Client
}
