package auth

import "net/http"

func SetCookie(token string, time int) *http.Cookie{
	cookie := http.Cookie{
		Name:     "jwttoken",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteDefaultMode,
		MaxAge:   time,
	}

	return &cookie
}