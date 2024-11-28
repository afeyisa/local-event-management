package middleware

import (
	"net/http"
	"os"
)

func HasuraAuth(next http.Handler) http.Handler{

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("X-ACTION-SECRET")
			expectedSecret := os.Getenv("ACTION_SECRET")
			if expectedSecret != header{
				http.Error(w, "Unauthorized action", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w,r)
		},
	)
}