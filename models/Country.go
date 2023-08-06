package models

import "github.com/google/uuid"

type Country struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Countries []Country
