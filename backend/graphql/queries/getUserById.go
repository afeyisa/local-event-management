package queries

import "github.com/google/uuid"

type UsersById struct {
	Data_users []struct {
		User_id  uuid.UUID `json:"user_id"`
	} `graphql:"data_users(where: {user_id: {_eq: $user_id}})"`
}