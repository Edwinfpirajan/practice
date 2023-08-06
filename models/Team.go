package models

import "github.com/google/uuid"

type Team struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	FkCityId uuid.UUID `gorm:"column:fk_city_id"`
}

type Teams []Team
