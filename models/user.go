package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Password  string
	FirstName string
	LastName  string
	Email     string
	FkTeamID  uuid.UUID `gorm:"column:fk_team_id"`
}
type Users []User
