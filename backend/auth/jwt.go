package auth

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)
		
//**************authentication with JWT*******************************************************
func MakeJWT(userID uuid.UUID, tokenSecret []byte, role string, expiresIn time.Duration) (string, error) {
	fmt.Println("making token")
	
	claims := &jwt.MapClaims{
		
		"iss": "local-event-management",
		"sub": userID.String(),
		"iat": jwt.NewNumericDate(time.Now().UTC()),
		"exp": jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		"hasura":  map[string]interface{}{
			"claims": map[string]interface{}{
			"x-hasura-allowed-roles": []string{"anonymous","user"},
			"x-hasura-default-role": role,
			"x-hasura-user-id":  userID.String() ,
			},

		},
	
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err 
	}

	return tokenString, nil
}


func ValidateJWT(tokenString string, tokenSecret []byte) (uuid.UUID, error) {
	// Parse the token and validate its claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
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
