package middleware

import (
	"music-exercise-tracking/internal/auth"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Client == nil {
			url := auth.AuthURL()
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}
