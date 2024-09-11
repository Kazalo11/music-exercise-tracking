package somepackage

import (
	"context"
	"encoding/json"
	"log"
	clientManager "music-exercise-tracking/internal/client"
	"music-exercise-tracking/internal/mapping"
	"net/http"

	"github.com/zmb3/spotify/v2"
)

func GetRecentlyPlayed(w http.ResponseWriter, r *http.Request) {
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
	response := map[string][]mapping.Song{"songs": mapping.MapSpotifySongs(playerResult)}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
