package mutations

import "github.com/google/uuid"

type InsertUser struct {
	Insert_data_users_one struct{
		User_id uuid.UUID 
		Role string
		}`graphql:"insert_data_users_one(object: {email: $email,role:$role ,password: $password})"`
}