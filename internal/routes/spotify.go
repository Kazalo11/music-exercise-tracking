package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zmb3/spotify/v2"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	clientManager "music-exercise-tracking/internal/client"
	songs "music-exercise-tracking/internal/songs"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadRecentlyPlayed),
	)
	state = "abc123"
)

func SpotifyRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/callback", completeAuth)
	r.Get("/auth", getSpotifyAuthURL)
	r.Get("/songs", songs.GetRecentlyPlayed)
	return r

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
	clientManager.SetClient(client)

	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}

func getSpotifyAuthURL(w http.ResponseWriter, r *http.Request) {
	url := auth.AuthURL(state)

	response := map[string]string{"url": url}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
