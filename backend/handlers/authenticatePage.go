package handlers

import (
	"encoding/json"
	"local-event-management-backend/auth"
	"net/http"
	"os"
	"strings"
)

func ProtectPageRoute(w http.ResponseWriter, r *http.Request) {

		// Get the JWT token from the cookies
		c, err :=r.Cookie("jwttoken")
		if  c.String() == "" || err != nil {
			w.WriteHeader(http.StatusOK)       
			json.NewEncoder(w).Encode(false) 
			return
		}
				
		token := strings.TrimPrefix(c.String(), "jwttoken=")

		_, err = auth.ValidateJWT(token, []byte(os.Getenv("JWT_KEY")))
		if err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(false)       
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)        
	
}