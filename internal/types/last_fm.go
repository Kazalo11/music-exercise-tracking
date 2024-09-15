package types

type artist struct {
	Name string `json:"#text"`
}
type nowPlaying struct {
	NowPlaying string `json:"nowplaying"`
}
type album struct {
	id   string `json:"mbid"`
	Name string `json:"#text"`
}
type image struct {
	size string `json:"size"`
	Link string `json:"#text"`
}

type track struct {
	Name       string     `json:"name"`
	Artist     artist     `json:"artist"`
	Image      []image    `json:"image"`
	URL        string     `json:"url"`
	NowPlaying nowPlaying `json:"@attr"`
}

type RecentTracksResponse struct {
	RecentTracks []track `json:"recenttracks"`
}
