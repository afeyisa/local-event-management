package actions

import (
	"encoding/json"
	"net/http"
)

type LogOutResponse struct{
	Success bool `json:"success"`
}


func LogOutHandler(w http.ResponseWriter, r *http.Request) {

	    cookie := http.Cookie{
        Name:     "jwttoken",
        Value:    "hh",
        Path:     "/",
        HttpOnly: true, 
        Secure:   true, 
        SameSite: http.SameSiteDefaultMode,
        MaxAge:   -86400, // Cookie expiration (1 day)
    }
    http.SetCookie(w, &cookie)
	res := LogOutResponse{Success: true}
	data, _ := json.Marshal(res)
    w.Write(data)
}