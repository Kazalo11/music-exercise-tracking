package mapping

import (
	"time"

	"github.com/zmb3/spotify/v2"
)

type Song struct {
	Name     string
	Artists  []string
	PlayedAt time.Time
}

func mapSpotifySongs(spSongs []spotify.RecentlyPlayedItem) []Song {
	var songs []Song
	for _, song := range spSongs {
		songs = append(songs, Song{
			Name:     song.Track.Name,
			Artists:  mapArtists(song.Track.Artists),
			PlayedAt: song.PlayedAt,
		})
	}
	return songs

}

func mapArtists(spArtists []spotify.SimpleArtist) []string {
	var artists []string
	for _, artist := range spArtists {
		artists = append(artists, artist.Name)
	}
	return artists

}
