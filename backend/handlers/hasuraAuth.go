package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"local-event-management-backend/auth"

	"github.com/google/uuid"
)


type hasuraResponse struct{
	XHasuraUserId string `json:"X-Hasura-User-Id"`
	XHasuraRole string `json:"X-Hasura-Role"`
	XHasuraAllowedRole string `json:"X-Hasura-Allowed-Role"`
}


func HasuraAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test auth")
	// it check both cookies and header to check where the token is carried
	c, err := r.Cookie("jwttoken")
	if err != nil {
		authHeader := r.Header.Get("Authorization")
		id, err := checkHeader(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(hasuraResponse{
				XHasuraRole:        "anonymous",
				XHasuraAllowedRole: "anonymous",
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hasuraResponse{
			XHasuraUserId:      fmt.Sprintf("%v", id),
			XHasuraRole:        "user",
			XHasuraAllowedRole: "anonymous",
		})
		return

	}

	id, err := checkCookies(c.String())
	if err != nil {
		authHeader := r.Header.Get("Authorization")
		id, err := checkHeader(authHeader)
		if err != nil {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(hasuraResponse{
				XHasuraRole:        "anonymous",
				XHasuraAllowedRole: "anonymous",
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hasuraResponse{
			XHasuraUserId:      fmt.Sprintf("%v", id),
			XHasuraRole:        "user",
			XHasuraAllowedRole: "anonymous",
		})
		return

	}
	res := hasuraResponse{
		XHasuraUserId:      fmt.Sprintf("%v", id),
		XHasuraRole:        "user",
		XHasuraAllowedRole: "anonymous",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func checkCookies(c string) (uuid.UUID, error){
	if  c == "" {
				
		return uuid.UUID{} , fmt.Errorf("no cookie")
	}
			
	token := strings.TrimPrefix(c, "jwttoken=")
	return  auth.ValidateJWT(token, []byte(os.Getenv("JWT_KEY")))
}

func checkHeader(h string) (uuid.UUID,error){
	// authHeader := h.Get("Authorization")
	if h == "" || !strings.HasPrefix(h, "Bearer "){       
		return uuid.UUID{}, fmt.Errorf("no header")
	}
	// Extract the token by stripping the "Bearer " prefix
	token := strings.TrimPrefix(h, "Bearer ")
    return auth.ValidateJWT(token, []byte(os.Getenv("JWT_KEY")))

}

