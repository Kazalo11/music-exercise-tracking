package types

import "time"

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

type date struct {
	Timestamp string `json:"uts"`
	DateTime  string `json:"#text"`
}

type track struct {
	Name       string     `json:"name"`
	Artist     artist     `json:"artist"`
	Image      []image    `json:"image"`
	URL        string     `json:"url"`
	NowPlaying nowPlaying `json:"@attr"`
	Album      album      `json:"album"`
	Date       date       `json:"date"`
}

type tracks struct {
	Tracks []track `json:"track"`
}

type RecentTracksResponse struct {
	RecentTracks tracks `json:"recenttracks"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05Z"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type RecentTracksRequest struct {
	UserName string     `json:"user_name"`
	Start    CustomTime `json:"start"`
	End      CustomTime `json:"end"`
}
