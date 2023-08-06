package models

import "github.com/google/uuid"

type City struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Country uuid.UUID `gorm:"column:fk_country_id"`
}

type Cities []City
