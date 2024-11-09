package queries

import "github.com/google/uuid"

type Users struct {
	Data_users []struct {
		User_id  uuid.UUID `json:"user_id"`
		Password string    `json:"password"`
		Role string `json:"role"`
	} `graphql:"data_users(where: {email: {_eq: $email}})"`
}