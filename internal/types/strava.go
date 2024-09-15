package types

import "time"

type TokenReponse struct {
	TokenType string `json:"token_type"`
	RefreshTokenResponse
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

type RefreshTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
}

type Athlete struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

type Activity struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start_date"`
	TimeTaken int       `json:"elapsed_time"`
	Finish    time.Time `json:"finish_date,omitempty"`
}

func (a Activity) CalculateFinishTime() time.Time {
	return a.Start.Add(time.Duration(a.TimeTaken) * time.Second)
}
