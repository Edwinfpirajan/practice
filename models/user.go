package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	TeamGroup Team
}

type Users []User
