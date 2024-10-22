package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type hasuraResponse struct{
	XHasuraUserId string `json:"X-Hasura-User-Id"`
	XHasuraRole string `json:"X-Hasura-Role"`
	XHasuraAllowedRole string `json:"X-Hasura-Allowed-Role"`
}

//*****************password************************************************************
// hash password for newly joining user
func HashPassword(password string) (string, error){
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

// compare hashed password with plain text password.
// password is plain text from frontend.
// hash is hashed password from hasura.
func CheckPasswordHash(password, hash string) error{
    err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
    return err
}
//********************************************************************************************

//**************authentication with JWT*******************************************************
func MakeJWT(userID uuid.UUID, tokenSecret []byte, expiresIn time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    "local-event-management",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err // Consider adding context here
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, tokenSecret []byte) (uuid.UUID, error) {
	// Parse the token and validate its claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return tokenSecret, nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Extract and return the user ID from the subject
		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			return uuid.Nil, fmt.Errorf("invalid subject in token: %v", claims.Subject)
		}
		return userID, nil
	}
	return uuid.Nil, fmt.Errorf("invalid token")
}

func checkCookies(c string) (uuid.UUID, error){
	if  c == "" {
				
		return uuid.UUID{} , fmt.Errorf("no cookie")
	}
			
	token := strings.TrimPrefix(c, "jwttoken=")
	return  ValidateJWT(token, []byte(os.Getenv("JWT_KEY")))
}
func checkHeader(h string) (uuid.UUID,error){
	// authHeader := h.Get("Authorization")
	if h == "" || !strings.HasPrefix(h, "Bearer "){       
		return uuid.UUID{}, fmt.Errorf("no header")
	}
	// Extract the token by stripping the "Bearer " prefix
	token := strings.TrimPrefix(h, "Bearer ")
    return ValidateJWT(token, []byte(os.Getenv("JWT_KEY")))

}


func HasuraAuthHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("test auth")
// it check both cookies and header to check where the token is carried
	c, err :=r.Cookie("jwttoken")
	if err != nil{
		authHeader := r.Header.Get("Authorization")
		id,err := checkHeader(authHeader)
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
			XHasuraUserId:    fmt.Sprintf("%v", id),
			XHasuraRole:      "user",
			XHasuraAllowedRole: "anonymous",
		})
		return

	}

	id,err := checkCookies(c.String())
	if err != nil{
		authHeader := r.Header.Get("Authorization")
		id,err := checkHeader(authHeader)
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
			XHasuraUserId:    fmt.Sprintf("%v", id),
			XHasuraRole:      "user",
			XHasuraAllowedRole: "anonymous",
		})
		return

	}
	res := hasuraResponse{
        XHasuraUserId:    fmt.Sprintf("%v", id),
        XHasuraRole:      "user",
		XHasuraAllowedRole: "anonymous",
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)

}

// **************************************************************************************
